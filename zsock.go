package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"
#include <stdlib.h>
#include <string.h>

int Zsock_connect(zsock_t *self, const char *format) {return zsock_connect(self, format, NULL);}
int Zsock_bind(zsock_t *self, const char *format) {return zsock_bind(self, format, NULL);}
*/
import "C"

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"unsafe"
)

// Zsock wraps the zsock_t class in CZMQ.
type Zsock struct {
	zsock_t *C.struct__zsock_t
	file    string
	line    int
	zType   Type
}

// NewZsock creates a new socket.  The caller source and
// line number are passed so CZMQ can report socket leaks
// intelligently.
func NewZsock(t Type) *Zsock {
	var z *Zsock
	_, file, line, ok := runtime.Caller(1)

	if ok {
		z = &Zsock{file: file, line: line, zType: t}
	} else {
		z = &Zsock{file: "", line: 0, zType: t}
	}

	z.zsock_t = C.zsock_new_(C.int(z.zType), C.CString(z.file), C.size_t(z.line))
	return z
}

// NewPUB creates a PUB socket.  The endpoint is empty, or starts with
// '@' (connect) or '>' (bind).  Multiple endpoints are allowed, separated
// by commas.  If the endpoint does not start with '@' or '>', it binds.
func NewPUB(endpoints string) (*Zsock, error) {
	z := NewZsock(PUB)
	z.zsock_t = C.zsock_new_(C.int(z.zType), C.CString(z.file), C.size_t(z.line))
	rc := C.zsock_attach(z.zsock_t, C.CString(endpoints), C._Bool(true))
	if rc == -1 {
		return nil, ErrZsockAttach
	}
	return z, nil
}

// NewSUB creates a SUB socket.  The enpoint is empty, or starts with
// '@' (connect) or '>' (bind).  Multiple endpoints are allowed, separated
// by commas.  If the endpoint does not start with '@' or '>', it connects.
// The second argument is a comma delimited list of topics to subscribe to.
func NewSUB(endpoints string, subscribe string) (*Zsock, error) {
	z := NewZsock(SUB)
	subscriptions := strings.Split(subscribe, ",")
	for _, s := range subscriptions {
		z.SetSubscribe(s)
	}

	rc := C.zsock_attach(z.zsock_t, C.CString(endpoints), C._Bool(false))
	if rc == -1 {
		return nil, ErrZsockAttach
	}
	return z, nil
}

// NewREP creates a REP socket.  The endpoint is empty, or starts with
// '@' (connect) or '>' (bind).  Multiple endpoints are allowed, separated
// by commas.  If the endpoint does not start with '@' or '>', it binds.
func NewREP(endpoints string) (*Zsock, error) {
	z := NewZsock(REP)
	rc := C.zsock_attach(z.zsock_t, C.CString(endpoints), C._Bool(true))
	if rc == -1 {
		return nil, ErrZsockAttach
	}
	return z, nil
}

// NewREQ creates a REQ socket.  The endpoint is empty, or starts with
// '@' (connect) or '>' (bind).  Multiple endpoints are allowed, separated
// by commas.  If the endpoint does not start with '@' or '>', it connects.
func NewREQ(endpoints string) (*Zsock, error) {
	z := NewZsock(REQ)
	rc := C.zsock_attach(z.zsock_t, C.CString(endpoints), C._Bool(false))
	if rc == -1 {
		return nil, ErrZsockAttach
	}
	return z, nil
}

// NewPULL creates a PULL socket.  The endpoint is empty, or starts with
// '@' (connect) or '>' (bind).  Multiple endpoints are allowed, separated
// by commas.  If the endpoint does not start with '@' or '>', it binds.
func NewPULL(endpoints string) (*Zsock, error) {
	z := NewZsock(PULL)
	rc := C.zsock_attach(z.zsock_t, C.CString(endpoints), C._Bool(true))
	if rc == -1 {
		return nil, ErrZsockAttach
	}
	return z, nil
}

// NewPUSH creates a PUSH socket.  The endpoint is empty, or starts with
// '@' (connect) or '>' (bind).  Multiple endpoints are allowed, separated
// by commas.  If the endpoint does not start with '@' or '>', it connects.
func NewPUSH(endpoints string) (*Zsock, error) {
	z := NewZsock(PUSH)
	rc := C.zsock_attach(z.zsock_t, C.CString(endpoints), C._Bool(false))
	if rc == -1 {
		return nil, ErrZsockAttach
	}
	return z, nil
}

// NewROUTER creates a ROUTER socket.  The endpoint is empty, or starts with
// '@' (connect) or '>' (bind).  Multiple endpoints are allowed, separated
// by commas.  If the endpoint does not start with '@' or '>', it binds.
func NewROUTER(endpoints string) (*Zsock, error) {
	z := NewZsock(ROUTER)
	rc := C.zsock_attach(z.zsock_t, C.CString(endpoints), C._Bool(true))
	if rc == -1 {
		return nil, ErrZsockAttach
	}
	return z, nil
}

// NewDEALER creates a DEALER socket.  The endpoint is empty, or starts with
// '@' (connect) or '>' (bind).  Multiple endpoints are allowed, separated
// by commas.  If the endpoint does not start with '@' or '>', it connects.
func NewDEALER(endpoints string) (*Zsock, error) {
	z := NewZsock(DEALER)
	rc := C.zsock_attach(z.zsock_t, C.CString(endpoints), C._Bool(false))
	if rc == -1 {
		return nil, ErrZsockAttach
	}
	return z, nil
}

// NewXPUB creates an XPUB socket.  The endpoint is empty, or starts with
// '@' (connect) or '>' (bind).  Multiple endpoints are allowed, separated
// by commas.  If the endpoint does not start with '@' or '>', it binds.
func NewXPUB(endpoints string) (*Zsock, error) {
	z := NewZsock(XPUB)
	rc := C.zsock_attach(z.zsock_t, C.CString(endpoints), C._Bool(true))
	if rc == -1 {
		return nil, ErrZsockAttach
	}
	return z, nil
}

// NewXSUB creates an XSUB socket.  The endpoint is empty, or starts with
// '@' (connect) or '>' (bind).  Multiple endpoints are allowed, separated
// by commas.  If the endpoint does not start with '@' or '>', it connects.
func NewXSUB(endpoints string) (*Zsock, error) {
	z := NewZsock(XSUB)
	rc := C.zsock_attach(z.zsock_t, C.CString(endpoints), C._Bool(false))
	if rc == -1 {
		return nil, ErrZsockAttach
	}
	return z, nil
}

// NewPAIR creates a PAIR socket.  The endpoint is empty, or starts with
// '@' (connect) or '>' (bind).  If the endpoint does not start with '@' or
// '>', it connects.
func NewPAIR(endpoints string) (*Zsock, error) {
	z := NewZsock(PAIR)
	rc := C.zsock_attach(z.zsock_t, C.CString(endpoints), C._Bool(false))
	if rc == -1 {
		return nil, ErrZsockAttach
	}
	return z, nil
}

// NewSTREAM creates a STREAM socket.  The endpoint is empty, or starts with
// '@' (connect) or '>' (bind).  Multiple endpoints are allowed, separated
// by commas.  If the endpoint does not start with '@' or '>', it connects.
func NewSTREAM(endpoints string) (*Zsock, error) {
	z := NewZsock(STREAM)
	rc := C.zsock_attach(z.zsock_t, C.CString(endpoints), C._Bool(false))
	if rc == -1 {
		return nil, ErrZsockAttach
	}
	return z, nil
}

// Connect connects a socket to an endpoint
// returns an error if the connect failed.
func (z *Zsock) Connect(endpoint string) error {
	rc := C.Zsock_connect(z.zsock_t, C.CString(endpoint))
	if rc == C.int(-1) {
		return errors.New("failed")
	}
	return nil
}

// Bind binds a socket to an endpoint.  On success returns
// the port number used for tcp transports, or 0 for other
// transports.  On failure returns a -1 for port, and an error.
func (z *Zsock) Bind(endpoint string) (int, error) {
	port := C.Zsock_bind(z.zsock_t, C.CString(endpoint))
	if port == C.int(-1) {
		return -1, errors.New("failed")
	}
	return int(port), nil
}

// SendMessage is a variadic function that currently accepts ints,
// strings, and bytes, and sends them as an atomic multi frame
// message over zeromq as a series of byte arrays.  In the case
// of numeric data, the resulting byte array is a textual representation
// of the number (e.g., 100 turns to "100").  This may be changed to
// network byte ordered representation in the near future - I have
// not decided yet!
func (z *Zsock) SendMessage(parts ...interface{}) error {
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
			err := z.SendString(fmt.Sprintf("%d", val.(int)), f)
			if err != nil {
				return err
			}
		case string:
			err := z.SendString(val.(string), f)
			if err != nil {
				return err
			}
		case []byte:
			var err error
			if len(val.([]byte)) == 0 {
				err = z.SendBytes([]byte{0}, f)
			} else {
				err = z.SendBytes(val.([]byte), f)
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

// RecvMessage receives a full message from the socket
// and returns it as an array of byte arrays.
func (z *Zsock) RecvMessage() ([][]byte, error) {
	var msg [][]byte
	for {
		frame, flag, err := z.RecvBytes()
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
func (z *Zsock) SendBytes(data []byte, flags Flag) error {
	frame := C.zframe_new(unsafe.Pointer(&data[0]), C.size_t(len(data)))
	rc := C.zframe_send(&frame, unsafe.Pointer(z.zsock_t), C.int(flags))
	if rc == C.int(-1) {
		return errors.New("failed")
	}
	return nil
}

// SendString sends a string via the socket.  For the flags
// value, use 0 for a single message, or SNDMORE if it is
// a multi-part message
func (z *Zsock) SendString(data string, flags Flag) error {
	err := z.SendBytes([]byte(data), flags)
	return err
}

// RecvBytes reads a frame from the socket and returns it
// as a byte array,  Returns an error if the call fails.
func (z *Zsock) RecvBytes() ([]byte, Flag, error) {
	frame := C.zframe_recv(unsafe.Pointer(z.zsock_t))
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
func (z *Zsock) RecvString() (string, error) {
	b, _, err := z.RecvBytes()
	if err != nil {
		return "", err
	}
	return string(b), err
}

// Destroy destroys the underlying zsock_t.
func (z *Zsock) Destroy() {
	C.zsock_destroy_(&z.zsock_t, C.CString(z.file), C.size_t(z.line))
}
