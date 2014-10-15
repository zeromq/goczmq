package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"

zpoller_t *Poller_new(void *reader) { zpoller_t *poller = zpoller_new(reader, NULL); return poller; }
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// Poller is a simple poller for Socks
type Poller struct {
	zpoller_t *C.struct__zpoller_t
	socks     []*Sock
}

// NewPoller creates a new Poller instance.  It accepts one or more readers to poll.
func NewPoller(readers ...*Sock) (*Poller, error) {
	if len(readers) == 0 {
		return nil, fmt.Errorf("requires at least one reader")
	}

	z := &Poller{
		zpoller_t: C.Poller_new(unsafe.Pointer(readers[0].zsock_t)),
		socks:     make([]*Sock, 0),
	}

	z.socks = append(z.socks, readers[0])
	if len(readers) == 1 {
		return z, nil
	}

	for i, reader := range readers[1:] {
		rc := C.zpoller_add(z.zpoller_t, unsafe.Pointer(reader.zsock_t))
		if int(rc) == -1 {
			return z, fmt.Errorf("error creating proxy")
		}
		z.socks = append(z.socks, readers[i])
	}
	return z, nil
}

// Add adds a reader to be polled.
func (p *Poller) Add(reader *Sock) error {
	rc := C.zpoller_add(p.zpoller_t, unsafe.Pointer(reader.zsock_t))
	if int(rc) == -1 {
		return fmt.Errorf("error adding reader")
	}
	p.socks = append(p.socks, reader)
	return nil
}

// Remove removes a zsock from the poller
func (p *Poller) Remove(reader *Sock) {
	num_items := len(p.socks)
	for i := 0; i < num_items; i++ {
		if p.socks[i] == reader {
			if i == num_items-1 {
				p.socks = p.socks[:i]
			} else {
				p.socks = append(p.socks[:i], p.socks[i+1:]...)
			}
		}
	}
}

// Wait waits for the timeout period in milliseconds for a POLLIN
// event, and returns the first socket that returns one
func (p *Poller) Wait(timeout int) *Sock {
	s := C.zpoller_wait(p.zpoller_t, C.int(timeout))
	s = unsafe.Pointer(s)
	if s == nil {
		return nil
	}
	for _, zsock := range p.socks {
		if unsafe.Pointer(zsock.zsock_t) == s {
			return zsock
		}
	}
	return nil
}

// Destroy destroys the Poller
func (p *Poller) Destroy() {
	C.zpoller_destroy(&p.zpoller_t)
}
