package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"

zactor_t *Zproxy_new () { zactor_t *proxy = zactor_new(zproxy, NULL); return proxy; }
*/
import "C"

import (
	"unsafe"
)

// Zproxy actors switch messages between a frontend and backend socket.  The
// Zproxy struct holds a reference to a CZMQ zactor_t.
type Proxy struct {
	zactor_t *C.struct__zactor_t
}

// NewZproxy creates a new Zproxy instance.
func NewProxy() *Proxy {
	p := &Proxy{}
	p.zactor_t = C.Zproxy_new()
	return p
}

// SetFrontend accepts a socket type and endpoint, and sends a message
// to the zactor thread telling it to set up a socket bound to the endpoint.
func (p *Proxy) SetFrontend(sockType Type, endpoint string) error {
	typeString := getStringType(sockType)

	rc := C.zstr_sendm(unsafe.Pointer(p.zactor_t), C.CString("FRONTEND"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_sendm(unsafe.Pointer(p.zactor_t), C.CString(typeString))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(p.zactor_t), C.CString(endpoint))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zsock_wait(unsafe.Pointer(p.zactor_t))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// SetBackend accepts a socket type and endpoint, and sends a message
// to the zactor thread telling it to set up a socket bound to the endpoint.
func (p *Proxy) SetBackend(sockType Type, endpoint string) error {
	typeString := getStringType(sockType)

	rc := C.zstr_sendm(unsafe.Pointer(p.zactor_t), C.CString("BACKEND"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_sendm(unsafe.Pointer(p.zactor_t), C.CString(typeString))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(p.zactor_t), C.CString(endpoint))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zsock_wait(unsafe.Pointer(p.zactor_t))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// SetCapture accepts a socket endpoint and sets up a PUSH socket bound
// to that endpoint, that sends a copy of all messages passing through
// the proxy.
func (p *Proxy) SetCapture(endpoint string) error {
	rc := C.zstr_sendm(unsafe.Pointer(p.zactor_t), C.CString("CAPTURE"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(p.zactor_t), C.CString(endpoint))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Pause sends a message to the zproxy actor telling it to pause.
func (p *Proxy) Pause() error {
	rc := C.zstr_send(unsafe.Pointer(p.zactor_t), C.CString("PAUSE"))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Resume sends a message to the zproxy actor telling it to resume.
func (p *Proxy) Resume() error {
	rc := C.zstr_send(unsafe.Pointer(p.zactor_t), C.CString("RESUME"))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Verbose sets the proxy to log information to stdout.
func (p *Proxy) Verbose() error {
	rc := C.zstr_send(unsafe.Pointer(p.zactor_t), C.CString("VERBOSE"))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Destroy destroys the proxy.
func (p *Proxy) Destroy() {
	C.zactor_destroy(&p.zactor_t)
}
