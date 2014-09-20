package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"

zactor_t *Zgossip_new (char *name) { zactor_t *gossip = zactor_new(zgossip, name); return gossip; }
*/
import "C"

import (
	"unsafe"
)

// Zgossip actors use a gossip protocol for decentralized configuration management.
// Zgossip nodes form a loosely connected network that publishes and redistributed
// name/value tuples.  A network of Zgossip actors will eventually achieve
// a consistent state
type Zgossip struct {
	zactor_t *C.struct__zactor_t
}

// NewZgossip creates a new Zgossip actor
func NewZgossip(name string) *Zgossip {
	z := &Zgossip{}
	z.zactor_t = C.Zgossip_new(C.CString(name))
	return z
}

// Bind binds the gossip service to a specified endpoint
func (z *Zgossip) Bind(endpoint string) error {
	rc := C.zstr_sendm(unsafe.Pointer(z.zactor_t), C.CString("BIND"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(z.zactor_t), C.CString(endpoint))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Verbose sets the gossip actor to log information to stdout.
func (z *Zgossip) Verbose() error {
	rc := C.zstr_send(unsafe.Pointer(z.zactor_t), C.CString("VERBOSE"))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Destroy destroys the gossip actor.
func (z *Zgossip) Destroy() {
	C.zactor_destroy(&z.zactor_t)
}
