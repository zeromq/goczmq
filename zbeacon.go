package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"

zactor_t *Zbeacon_new () { zactor_t *beacon = zactor_new(zbeacon, NULL); return beacon; }
*/
import "C"

import (
	"strconv"
	"unsafe"
)

type Zbeacon struct {
	zactor_t *C.struct__zactor_t
}

// NewZbeacon creates a new Zbeacon instance.
func NewZbeacon() *Zbeacon {
	z := &Zbeacon{}
	z.zactor_t = C.Zbeacon_new()
	return z
}

// Verbose sets the beacon to log information to stdout.
func (z *Zbeacon) Verbose() error {
	rc := C.zstr_send(unsafe.Pointer(z.zactor_t), C.CString("VERBOSE"))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Configure accepts a port number and configures the beacon, returning an address
func (z *Zbeacon) Configure(port int) (string, error) {
	rc := C.zstr_sendm(unsafe.Pointer(z.zactor_t), C.CString("CONFIGURE"))
	if rc == -1 {
		return "", ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(z.zactor_t), C.CString(strconv.Itoa(port)))
	if rc == -1 {
		return "", ErrActorCmd
	}

	Chostname := C.zstr_recv(unsafe.Pointer(z.zactor_t))
	hostname := C.GoString(Chostname)
	return hostname, nil
}

// Publish publishes an announcement at an interval
func (z *Zbeacon) Publish(announcement string, interval int) error {
	rc := C.zstr_sendm(unsafe.Pointer(z.zactor_t), C.CString("PUBLISH"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_sendm(unsafe.Pointer(z.zactor_t), C.CString(announcement))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(z.zactor_t), C.CString(strconv.Itoa(interval)))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Destroy destroys the beacon.
func (z *Zbeacon) Destroy() {
	C.zactor_destroy(&z.zactor_t)
}
