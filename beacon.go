package goczmq

/*
#include "czmq.h"

zactor_t *Beacon_new () {
	zactor_t *beacon = zactor_new(zbeacon, NULL); return beacon;
}
*/
import "C"

import (
	"strconv"
	"unsafe"
)

// Beacon wraps the CZMQ beacon actor. It implements a
// peer-to-peer discovery service for local networks.  Beacons
// can broadcast and receive UDPv4 service broadcasts.
type Beacon struct {
	zactorT *C.struct__zactor_t
}

// NewBeacon creates a new Beacon instance.
func NewBeacon() *Beacon {
	z := &Beacon{}
	z.zactorT = C.Beacon_new()
	return z
}

// Verbose sets the beacon to log information to stdout.
func (b *Beacon) Verbose() error {
	rc := C.zstr_send(unsafe.Pointer(b.zactorT), C.CString("VERBOSE"))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Configure accepts a port number and configures
// the beacon, returning an address
func (b *Beacon) Configure(port int) (string, error) {
	rc := C.zstr_sendm(unsafe.Pointer(b.zactorT), C.CString("CONFIGURE"))
	if rc == -1 {
		return "", ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(b.zactorT), C.CString(strconv.Itoa(port)))
	if rc == -1 {
		return "", ErrActorCmd
	}

	Chostname := C.zstr_recv(unsafe.Pointer(b.zactorT))
	hostname := C.GoString(Chostname)
	return hostname, nil
}

// Publish publishes an announcement at an interval
func (b *Beacon) Publish(announcement string, interval int) error {
	rc := C.zstr_sendm(unsafe.Pointer(b.zactorT), C.CString("PUBLISH"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_sendm(unsafe.Pointer(b.zactorT), C.CString(announcement))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(b.zactorT),
		C.CString(strconv.Itoa(interval)))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Subscribe subscribes to beacons matching the filter
func (b *Beacon) Subscribe(filter string) error {
	rc := C.zstr_sendm(unsafe.Pointer(b.zactorT), C.CString("SUBSCRIBE"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(b.zactorT), C.CString(filter))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Recv waits for the specific timeout in milliseconds to receive a beacon
func (b *Beacon) Recv(timeout int) [][]byte {
	C.zsock_set_rcvtimeo(unsafe.Pointer(b.zactorT), C.int(timeout))
	frame := C.zframe_recv(unsafe.Pointer(b.zactorT))
	dataSize := C.zframe_size(frame)
	dataPtr := C.zframe_data(frame)
	ipBytes := C.GoBytes(unsafe.Pointer(dataPtr), C.int(dataSize))
	C.zframe_destroy(&frame)

	C.zsock_set_rcvtimeo(unsafe.Pointer(b.zactorT), C.int(timeout))
	frame = C.zframe_recv(unsafe.Pointer(b.zactorT))
	dataSize = C.zframe_size(frame)
	dataPtr = C.zframe_data(frame)
	msgBytes := C.GoBytes(unsafe.Pointer(dataPtr), C.int(dataSize))
	C.zframe_destroy(&frame)

	return [][]byte{ipBytes, msgBytes}
}

// Destroy destroys the beacon.
func (b *Beacon) Destroy() {
	C.zactor_destroy(&b.zactorT)
}
