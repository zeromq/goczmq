package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"

zactor_t *Auth_new () {
	zactor_t *auth = zactor_new(zauth, NULL); return auth;
}
*/
import "C"

import (
	"unsafe"
)

// Auth wraps a CZMQ zauth zactor
type Auth struct {
	zactorT *C.struct__zactor_t
}

// NewAuth creates a new Auth actor.
func NewAuth() *Auth {
	z := &Auth{}
	z.zactorT = C.Auth_new()
	return z
}

// Verbose sets the auth actor to log information to stdout.
func (a *Auth) Verbose() error {
	rc := C.zstr_send(unsafe.Pointer(a.zactorT), C.CString("VERBOSE"))
	if rc == -1 {
		return ErrActorCmd
	}
	C.zsock_wait(unsafe.Pointer(a.zactorT))

	return nil
}

// Deny adds an address to a socket's deny list
func (a *Auth) Deny(address string) error {
	rc := C.zstr_sendm(unsafe.Pointer(a.zactorT), C.CString("DENY"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(a.zactorT), C.CString(address))
	if rc == -1 {
		return ErrActorCmd
	}

	C.zsock_wait(unsafe.Pointer(a.zactorT))

	return nil
}

// Allow removes a previous Deny
func (a *Auth) Allow(address string) error {
	rc := C.zstr_sendm(unsafe.Pointer(a.zactorT), C.CString("ALLOW"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(a.zactorT), C.CString(address))
	if rc == -1 {
		return ErrActorCmd
	}

	C.zsock_wait(unsafe.Pointer(a.zactorT))

	return nil
}

// Curve sets auth method to curve
func (a *Auth) Curve(allowed string) error {
	rc := C.zstr_sendm(unsafe.Pointer(a.zactorT), C.CString("CURVE"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(a.zactorT), C.CString(allowed))
	if rc == -1 {
		return ErrActorCmd
	}

	C.zsock_wait(unsafe.Pointer(a.zactorT))

	return nil
}

// Plain sets auth method to plain
func (a *Auth) Plain(directory string) error {
	rc := C.zstr_sendm(unsafe.Pointer(a.zactorT), C.CString("PLAIN"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(a.zactorT), C.CString(directory))
	if rc == -1 {
		return ErrActorCmd
	}

	C.zsock_wait(unsafe.Pointer(a.zactorT))

	return nil
}

// Destroy destroys the auth actor.
func (a *Auth) Destroy() {
	C.zactor_destroy(&a.zactorT)
}
