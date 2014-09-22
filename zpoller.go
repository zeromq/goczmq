package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"

zpoller_t *Zpoller_new(void *reader) { zpoller_t *poller = zpoller_new(reader, NULL); return poller; }
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// Zpoller is a simple poller for Zsocks
type Zpoller struct {
	zpoller_t *C.struct__zpoller_t
	zsocks    []*Zsock
}

// NewZpoller creates a new Zpoller instance.  It accepts one or more readers to poll.
func NewZpoller(readers ...*Zsock) (*Zpoller, error) {
	if len(readers) == 0 {
		return nil, fmt.Errorf("requires at least one reader")
	}

	z := &Zpoller{
		zpoller_t: C.Zpoller_new(unsafe.Pointer(readers[0].zsock_t)),
		zsocks:    make([]*Zsock, 0),
	}

	z.zsocks = append(z.zsocks, readers[0])
	if len(readers) == 1 {
		return z, nil
	}

	for i, reader := range readers[1:] {
		rc := C.zpoller_add(z.zpoller_t, unsafe.Pointer(reader.zsock_t))
		if int(rc) == -1 {
			return z, fmt.Errorf("error creating proxy")
		}
		z.zsocks = append(z.zsocks, readers[i])
	}
	return z, nil
}

// Add adds a reader to be polled.
func (z *Zpoller) Add(reader *Zsock) error {
	rc := C.zpoller_add(z.zpoller_t, unsafe.Pointer(reader.zsock_t))
	if int(rc) == -1 {
		return fmt.Errorf("error adding reader")
	}
	z.zsocks = append(z.zsocks, reader)
	return nil
}

// Remove removes a zsock from the poller
func (z *Zpoller) Remove(reader *Zsock) {
	num_items := len(z.zsocks)
	for i := 0; i < num_items; i++ {
		if z.zsocks[i] == reader {
			if i == num_items-1 {
				z.zsocks = z.zsocks[:i]
			} else {
				z.zsocks = append(z.zsocks[:i], z.zsocks[i+1:]...)
			}
		}
	}
}

// Wait waits for the timeout period in milliseconds for a POLLIN
// event, and returns the first socket that returns one
func (z *Zpoller) Wait(timeout int) *Zsock {
	s := C.zpoller_wait(z.zpoller_t, C.int(timeout))
	s = unsafe.Pointer(s)
	if s == nil {
		return nil
	}
	for _, zsock := range z.zsocks {
		if unsafe.Pointer(zsock.zsock_t) == s {
			return zsock
		}
	}
	return nil
}

// Destroy destroys the Zpoller
func (z *Zpoller) Destroy() {
	C.zpoller_destroy(&z.zpoller_t)
}
