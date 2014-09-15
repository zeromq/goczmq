package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"

zpoller_t *Zpoller_new(void *reader) { zpoller_t *poller = zpoller_new(reader); return poller; }
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type Zpoller struct {
	zpoller_t *C.struct__zpoller_t
}

// NewZpoller creates a new Zpoller instance.  It accepts one or more readers to poll.
func NewZpoller(readers ...*Zsock) (*Zpoller, error) {
	if len(readers) == 0 {
		return nil, fmt.Errorf("requires at least one reader")
	}

	z := &Zpoller{}
	z.zpoller_t = C.Zpoller_new(unsafe.Pointer(readers[0]))
	if len(readers) == 1 {
		return z, nil
	}

	for _, reader := range readers[1:] {
		rc := C.zpoller_add(z.zpoller_t, unsafe.Pointer(reader))
		if int(rc) == -1 {
			return z, fmt.Errorf("error creating proxy")
		}
	}
	return z, nil
}

// Destroy destroys the Zpoller
func (z *Zpoller) Destroy() {
	C.zpoller_destroy(&z.zpoller_t)
}
