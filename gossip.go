package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"

zactor_t *Gossip_new (char *name) {
	zactor_t *gossip = zactor_new(zgossip, name);
	return gossip;
}
*/
import "C"

import "unsafe"

// Gossip actors use a gossip protocol for decentralized configuration management.
// Gossip nodes form a loosely connected network that publishes and redistributed
// name/value tuples.  A network of Gossip actors will eventually achieve
// a consistent state
type Gossip struct {
	zactorT *C.struct__zactor_t
}

// NewGossip creates a new Gossip actor
func NewGossip(name string) *Gossip {
	g := &Gossip{}
	g.zactorT = C.Gossip_new(C.CString(name))
	return g
}

// Bind binds the gossip service to a specified endpoint
func (g *Gossip) Bind(endpoint string) error {
	rc := C.zstr_sendm(unsafe.Pointer(g.zactorT), C.CString("BIND"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(g.zactorT), C.CString(endpoint))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Connect connects the gossip service to a specified endpoint
func (g *Gossip) Connect(endpoint string) error {
	rc := C.zstr_sendm(unsafe.Pointer(g.zactorT), C.CString("CONNECT"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(g.zactorT), C.CString(endpoint))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Publish announces a key / value pair
func (g *Gossip) Publish(key, value string) error {
	rc := C.zstr_sendm(unsafe.Pointer(g.zactorT), C.CString("PUBLISH"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_sendm(unsafe.Pointer(g.zactorT), C.CString(key))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(g.zactorT), C.CString(value))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Verbose sets the gossip actor to log information to stdout.
func (g *Gossip) Verbose() error {
	rc := C.zstr_send(unsafe.Pointer(g.zactorT), C.CString("VERBOSE"))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Destroy destroys the gossip actor.
func (g *Gossip) Destroy() {
	C.zactor_destroy(&g.zactorT)
}
