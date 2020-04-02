// +build draft

package goczmq

/*
#include "czmq.h"
int Sock_sendserverframe(
	zsock_t *sock,
	const void *data,
	size_t size,
	int flags,
	uint32_t routing_id
) {
	zframe_t *frame = zframe_new (data, size);
	assert(frame != NULL);
	zframe_set_routing_id(frame, routing_id);
	int rc = zframe_send (&frame, sock, flags);
	return rc;
}
*/
import "C"

import (
	"unsafe"
)

const (
	// Scatter is a ZMQ_SCATTER socket type
	Scatter = int(C.ZMQ_SCATTER)

	// Gather is a ZMQ_GATHER socket type
	Gather = int(C.ZMQ_GATHER)

	// Client is a ZMQ_CLIENT socket type
	Client = int(C.ZMQ_CLIENT)

	// Gather is a ZMQ_SERVER socket type
	Server = int(C.ZMQ_SERVER)
)

// NewGather creates a Gather socket and calls Attach.
// The socket will Bind by default.
func NewGather(endpoints string) (*Sock, error) {
	s := NewSock(Gather)
	return s, s.Attach(endpoints, true)
}

// NewScatter creates a Scatter socket and calls Attach.
// The socket will Connect by default.
func NewScatter(endpoints string) (*Sock, error) {
	s := NewSock(Scatter)
	return s, s.Attach(endpoints, false)
}

// NewServer creates a Server socket and calls Attach.
// The socket will Bind by default.
func NewServer(endpoints string) (*Sock, error) {
	s := NewSock(Server)
	return s, s.Attach(endpoints, true)
}

// NewClient creates a Client socket and calls Attach.
// The socket will Connect by default.
func NewClient(endpoints string) (*Sock, error) {
	s := NewSock(Client)
	return s, s.Attach(endpoints, false)
}

// RecvServerFrame reads a frame from the socket and returns it
// as a byte array, along with a more flag, routing ID and error
// (if there is an error)
func (s *Sock) RecvServerFrame() ([]byte, uint32, error) {
	if s.zsockT == nil {
		return nil, 0, ErrRecvFrameAfterDestroy
	}

	frame := C.zframe_recv(unsafe.Pointer(s.zsockT))
	if frame == nil {
		return []byte{0}, 0, ErrRecvFrame
	}
	dataSize := C.zframe_size(frame)
	dataPtr := C.zframe_data(frame)
	b := C.GoBytes(unsafe.Pointer(dataPtr), C.int(dataSize))
	var routing_id C.uint32_t = C.zframe_routing_id(frame)
	C.zframe_destroy(&frame)
	return b, uint32(routing_id), nil
}

// SendFrame sends a byte array via the socket.  For the flags
// value, use FlagNone (0) for a single message, or FlagMore if it is
// a multi-part message
func (s *Sock) SendServerFrame(data []byte, routing_id uint32) error {
	var rc C.int
	if len(data) == 0 {
		rc = C.Sock_sendserverframe(
			s.zsockT,
			nil,
			C.size_t(0),
			C.int(FlagNone),
			C.uint32_t(routing_id),
		)
	} else {
		rc = C.Sock_sendserverframe(
			s.zsockT,
			unsafe.Pointer(&data[0]),
			C.size_t(len(data)),
			C.int(FlagNone),
			C.uint32_t(routing_id),
		)
	}
	if rc == C.int(-1) {
		return ErrSendFrame
	}
	return nil
}
