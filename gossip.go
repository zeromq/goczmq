package goczmq

/*
#include "czmq.h"

zactor_t *Gossip_new (char *name) {
	zactor_t *gossip = zactor_new(zgossip, name);
	return gossip;
}
*/
import "C"

import "unsafe"

// Gossip wraps the CZMQ gossip actor.  This actor speaks a
// gossip protocol for decentralized configuration management.
// Nodes form a loosely connected network and publish name / value
// pair tuples.  Each node redistributes tuples it receives.
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
	cmd := C.CString("BIND")
	defer C.free(unsafe.Pointer(cmd))

	rc := C.zstr_sendm(unsafe.Pointer(g.zactorT), cmd)
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
	cmd := C.CString("CONNECT")
	defer C.free(unsafe.Pointer(cmd))

	rc := C.zstr_sendm(unsafe.Pointer(g.zactorT), cmd)
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
	cmd := C.CString("PUBLISH")
	defer C.free(unsafe.Pointer(cmd))

	rc := C.zstr_sendm(unsafe.Pointer(g.zactorT), cmd)
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
	cmd := C.CString("VERBOSE")
	defer C.free(unsafe.Pointer(cmd))

	rc := C.zstr_send(unsafe.Pointer(g.zactorT), cmd)
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Destroy destroys the gossip actor.
func (g *Gossip) Destroy() {
	C.zactor_destroy(&g.zactorT)
}
