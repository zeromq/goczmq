package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"
#include <stdlib.h>
#include <string.h>

int Sock_connect(zsock_t *self, const char *format) {return zsock_connect(self, format, NULL);}
int Sock_disconnect(zsock_t *self, const char *format) {return zsock_disconnect(self, format, NULL);}
int Sock_bind(zsock_t *self, const char *format) {return zsock_bind(self, format, NULL);}
int Sock_unbind(zsock_t *self, const char *format) {return zsock_unbind(self, format, NULL);}
*/
import "C"

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"unsafe"
)

// Sock wraps the zsock_t class in CZMQ.
type Sock struct {
	zsockT *C.struct__zsock_t
	file   string
	line   int
	zType  Type
}

func init() {
	if err := os.Setenv("ZSYS_SIGHANDLER", "false"); err != nil {
		panic(err)
	}
}

// NewSock creates a new socket.  The caller source and
// line number are passed so CZMQ can report socket leaks
// intelligently.
func NewSock(t Type) *Sock {
	var s *Sock
	_, file, line, ok := runtime.Caller(1)

	if ok {
		s = &Sock{
			file:  file,
			line:  line,
			zType: t,
		}
	} else {
		s = &Sock{
			file:  "",
			line:  0,
			zType: t,
		}
	}

	s.zsockT = C.zsock_new_(C.int(s.zType), C.CString(s.file), C.size_t(s.line))
	return s
}

// Connect connects a socket to an endpoint
// returns an error if the connect failed.
func (s *Sock) Connect(endpoint string) error {
	rc := C.Sock_connect(s.zsockT, C.CString(endpoint))
	if rc != C.int(0) {
		return errors.New("failed")
	}
	return nil
}

// Disconnect disconnects a socket from an endpoint.  If returns
// an error if the endpoint was not found
func (s *Sock) Disconnect(endpoint string) error {
	rc := C.Sock_disconnect(s.zsockT, C.CString(endpoint))
	if int(rc) == -1 {
		return fmt.Errorf("endopint was not bound")
	}
	return nil
}

// Bind binds a socket to an endpoint.  On success returns
// the port number used for tcp transports, or 0 for other
// transports.  On failure returns a -1 for port, and an error.
func (s *Sock) Bind(endpoint string) (int, error) {
	port := C.Sock_bind(s.zsockT, C.CString(endpoint))
	if port == C.int(-1) {
		return -1, errors.New("failed")
	}
	return int(port), nil
}

// Unbind unbinds a socket from an endpoint.  If returns
// an error if the endpoint was not found
func (s *Sock) Unbind(endpoint string) error {
	rc := C.Sock_unbind(s.zsockT, C.CString(endpoint))
	if int(rc) == -1 {
		return fmt.Errorf("endopint was not bound")
	}
	return nil
}

// Attach attaches a socket to zero or more endpoints. If endpoints is not null,
// parses as list of ZeroMQ endpoints, separated by commas, and prefixed by
// '@' (to bind the socket) or '>' (to attach the socket). If the endpoint
// does not start with '@' or '>', the serverish argument determines whether
// it is used to bind (serverish = true) or connect (serverish = false)
func (s *Sock) Attach(endpoints string, serverish bool) error {
	rc := C.zsock_attach(s.zsockT, C.CString(endpoints), C._Bool(serverish))
	if rc == -1 {
		return ErrSockAttach
	}
	return nil
}

// NewPUB creates a PUB socket and calls Attach.
// The socket will Bind by default.
func NewPUB(endpoints string) (*Sock, error) {
	s := NewSock(PUB)
	return s, s.Attach(endpoints, true)
}

// NewSUB creates a SUB socket and calls Attach.
// 'subscribe' is a comma delimited list of topics to subscribe to.
// The socket will Connect by default.
func NewSUB(endpoints string, subscribe string) (*Sock, error) {
	s := NewSock(SUB)
	subscriptions := strings.Split(subscribe, ",")

	for _, topic := range subscriptions {
		s.SetSubscribe(topic)
	}

	return s, s.Attach(endpoints, false)
}

// NewREP creates a REP socket and calls Attach.
// The socket will Bind by default.
func NewREP(endpoints string) (*Sock, error) {
	s := NewSock(REP)
	return s, s.Attach(endpoints, true)
}

// NewREQ creates a REQ socket and calls Attach.
// The socket will Connect by default.
func NewREQ(endpoints string) (*Sock, error) {
	s := NewSock(REQ)
	return s, s.Attach(endpoints, false)
}

// NewPULL creates a PULL socket and calls Attach.
// The socket will Bind by default.
func NewPULL(endpoints string) (*Sock, error) {
	s := NewSock(PULL)
	return s, s.Attach(endpoints, true)
}

// NewPUSH creates a PUSH socket and calls Attach.
// The socket will Connect by default.
func NewPUSH(endpoints string) (*Sock, error) {
	s := NewSock(PUSH)
	return s, s.Attach(endpoints, false)
}

// NewROUTER creates a ROUTER socket and calls Attach.
// The socket will Bind by default.
func NewROUTER(endpoints string) (*Sock, error) {
	s := NewSock(ROUTER)
	return s, s.Attach(endpoints, true)
}

// NewDEALER creates a DEALER socket and calls Attach.
// The socket will Connect by default.
func NewDEALER(endpoints string) (*Sock, error) {
	s := NewSock(DEALER)
	return s, s.Attach(endpoints, false)
}

// NewXPUB creates an XPUB socket and calls Attach.
// The socket will Bind by default.
func NewXPUB(endpoints string) (*Sock, error) {
	s := NewSock(XPUB)
	return s, s.Attach(endpoints, true)
}

// NewXSUB creates an XSUB socket and calls Attach.
// The socket will Connect by default.
func NewXSUB(endpoints string) (*Sock, error) {
	s := NewSock(XSUB)
	return s, s.Attach(endpoints, false)
}

// NewPAIR creates a PAIR socket and calls Attach.
// The socket will Connect by default.
func NewPAIR(endpoints string) (*Sock, error) {
	s := NewSock(PAIR)
	return s, s.Attach(endpoints, false)
}

// NewSTREAM creates a STREAM socket and calls Attach.
// The socket will Connect by default.
func NewSTREAM(endpoints string) (*Sock, error) {
	s := NewSock(STREAM)
	return s, s.Attach(endpoints, false)
}

// Pollin returns true if there is a POLLIN
// event on the socket
func (s *Sock) Pollin() bool {
	return s.Events() == POLLIN
	// return C.zsock_events(unsafe.Pointer(s.zsockT)) == C.ZMQ_POLLIN
}

// Pollout returns true if there is a POLLOUT
// event on the socket
func (s *Sock) Pollout() bool {
	return s.Events() == POLLOUT
}

// SendMessage is a variadic function that currently accepts ints,
// strings, and bytes, and sends them as an atomic multi frame
// message over zeromq as a series of byte arrays.  In the case
// of numeric data, the resulting byte array is a textual representation
// of the number (e.g., 100 turns to "100").  This may be changed to
// network byte ordered representation in the near future - I have
// not decided yet!
func (s *Sock) SendMessage(parts ...interface{}) error {
	numParts := len(parts)
	var f Flag

	var allParts []interface{}
	for _, part := range parts {
		switch t := part.(type) {
		case []string:
			for _, p := range t {
				allParts = append(allParts, p)
			}
		case [][]byte:
			for _, p := range t {
				allParts = append(allParts, p)
			}
		default:
			allParts = append(allParts, t)
		}
	}

	numParts = len(allParts)
	for i, val := range allParts {
		if i == numParts-1 {
			f = 0
		} else {
			f = MORE
		}

		switch val.(type) {
		case int:
			err := s.SendString(fmt.Sprintf("%d", val.(int)), f)
			if err != nil {
				return err
			}
		case string:
			err := s.SendString(val.(string), f)
			if err != nil {
				return err
			}
		case []byte:
			var err error
			if len(val.([]byte)) == 0 {
				err = s.SendBytes([]byte{0}, f)
			} else {
				err = s.SendBytes(val.([]byte), f)
			}
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type at index %d", i)
		}
	}
	return nil
}

// RecvMessageNoWait receives a full message from the socket
// and returns it as an array of byte arrays if one is waiting.
// returns an empty message and an error if one is not immediately
// available
func (s *Sock) RecvMessageNoWait() ([][]byte, error) {
	var msg [][]byte
	if !s.Pollin() {
		return msg, fmt.Errorf("no message")
	}

	for {
		frame, flag, err := s.RecvBytes()
		if err != nil {
			return msg, err
		}
		msg = append(msg, frame)
		if flag != MORE {
			break
		}
	}
	return msg, nil
}

// RecvMessage receives a full message from the socket
// and returns it as an array of byte arrays.
func (s *Sock) RecvMessage() ([][]byte, error) {
	var msg [][]byte

	for {
		frame, flag, err := s.RecvBytes()
		if err != nil {
			return msg, err
		}
		msg = append(msg, frame)
		if flag != MORE {
			break
		}
	}
	return msg, nil
}

// SendBytes sends a byte array via the socket.  For the flags
// value, use 0 for a single message, or SNDMORE if it is
// a multi-part message
func (s *Sock) SendBytes(data []byte, flags Flag) error {
	frame := C.zframe_new(unsafe.Pointer(&data[0]), C.size_t(len(data)))
	rc := C.zframe_send(&frame, unsafe.Pointer(s.zsockT), C.int(flags))
	if rc == C.int(-1) {
		return errors.New("failed")
	}
	return nil
}

// SendString sends a string via the socket.  For the flags
// value, use 0 for a single message, or SNDMORE if it is
// a multi-part message
func (s *Sock) SendString(data string, flags Flag) error {
	err := s.SendBytes([]byte(data), flags)
	return err
}

// RecvBytes reads a frame from the socket and returns it
// as a byte array,  Returns an error if the call fails.
func (s *Sock) RecvBytes() ([]byte, Flag, error) {
	frame := C.zframe_recv(unsafe.Pointer(s.zsockT))
	if frame == nil {
		return []byte{0}, 0, errors.New("failed")
	}
	dataSize := C.zframe_size(frame)
	dataPtr := C.zframe_data(frame)
	b := C.GoBytes(unsafe.Pointer(dataPtr), C.int(dataSize))
	more := C.zframe_more(frame)
	C.zframe_destroy(&frame)
	return b, Flag(more), nil
}

// RecvString reads a frame from the socket and returns it
// as a string,  Returns an error if the call fails.
func (s *Sock) RecvString() (string, error) {
	b, _, err := s.RecvBytes()
	if err != nil {
		return "", err
	}
	return string(b), err
}

// GetType returns the socket's type
func (s *Sock) GetType() Type {
	return s.zType
}

// Destroy destroys the underlying zsockT.
func (s *Sock) Destroy() {
	C.zsock_destroy_(&s.zsockT, C.CString(s.file), C.size_t(s.line))
}
