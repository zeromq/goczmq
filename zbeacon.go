package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"

zactor_t *Zbeacon_new () { zactor_t *beacon = zactor_new(zbeacon, NULL); return beacon; }
*/
import "C"

import ()

type Zbeacon struct {
	zactor_t *C.struct__zactor_t
}

// NewZbeacon creates a new Zbeacon instance.
func NewZbeacon() *Zbeacon {
	z := &Zbeacon{}
	z.zactor_t = C.Zbeacon_new()
	return z
}

// Destroy destroys the beacon.
func (z *Zbeacon) Destroy() {
	C.zactor_destroy(&z.zactor_t)
}
