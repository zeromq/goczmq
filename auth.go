package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"

zactor_t *Auth_new () { zactor_t *auth = zactor_new(zauth, NULL); return auth; }
*/
import "C"

import (
	"unsafe"
)

// Auth wraps a CZMQ zauth zactor
type Auth struct {
	zactor_t *C.struct__zactor_t
}

// NewAuth creates a new Auth actor.
func NewAuth() *Auth {
	z := &Auth{}
	z.zactor_t = C.Auth_new()
	return z
}

// Verbose sets the auth actor to log information to stdout.
func (z *Auth) Verbose() error {
	rc := C.zstr_send(unsafe.Pointer(z.zactor_t), C.CString("VERBOSE"))
	if rc == -1 {
		return ErrActorCmd
	}
	C.zsock_wait(unsafe.Pointer(z.zactor_t))

	return nil
}

// Deny adds an address to a socket's deny list
func (z *Auth) Deny(address string) error {
	rc := C.zstr_sendm(unsafe.Pointer(z.zactor_t), C.CString("DENY"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(z.zactor_t), C.CString(address))
	if rc == -1 {
		return ErrActorCmd
	}

	C.zsock_wait(unsafe.Pointer(z.zactor_t))

	return nil
}

// Allow removes a previous Deny
func (z *Auth) Allow(address string) error {
	rc := C.zstr_sendm(unsafe.Pointer(z.zactor_t), C.CString("ALLOW"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(z.zactor_t), C.CString(address))
	if rc == -1 {
		return ErrActorCmd
	}

	C.zsock_wait(unsafe.Pointer(z.zactor_t))

	return nil
}

// Destroy destroys the auth actor.
func (z *Auth) Destroy() {
	C.zactor_destroy(&z.zactor_t)
}
