// A Go interface to CZMQ
package czmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"
#include <stdlib.h>
#include <string.h>

int Zsock_connect(zsock_t *self, const char *format) {return zsock_connect(self, format, NULL);}

int Zsock_bind(zsock_t *self, const char *format) {return zsock_bind(self, format, NULL);}

void *my_memcpy(void *dest, const void *src, size_t n) {
	return memcpy(dest, src, n);
}
*/
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

type Type int
type Flag int

const (
	REQ    = Type(C.ZMQ_REQ)
	REP    = Type(C.ZMQ_REP)
	DEALER = Type(C.ZMQ_DEALER)
	ROUTER = Type(C.ZMQ_ROUTER)
	PUB    = Type(C.ZMQ_PUB)
	SUB    = Type(C.ZMQ_SUB)
	XPUB   = Type(C.ZMQ_XPUB)
	XSUB   = Type(C.ZMQ_XSUB)
	PUSH   = Type(C.ZMQ_PUSH)
	PULL   = Type(C.ZMQ_PULL)
	PAIR   = Type(C.ZMQ_PAIR)
	STREAM = Type(C.ZMQ_STREAM)

	ZMSG_TAG = 0x003cafe
	MORE     = Flag(C.ZFRAME_MORE)
	REUSE    = Flag(C.ZFRAME_REUSE)
	DONTWAIT = Flag(C.ZFRAME_DONTWAIT)
)

type Zsock struct {
	zsock_t *C.struct__zsock_t
	file    string
	line    int
	zType   Type
}

func NewZsock(t Type) *Zsock {
	var z *Zsock

	_, file, line, ok := runtime.Caller(1)

	if ok {
		z = &Zsock{file: file, line: line, zType: t}
	} else {
		z = &Zsock{file: "", line: 0, zType: t}
	}

	z.zsock_t = C.zsock_new_(C.int(t), C.CString(z.file), C.size_t(z.line))
	return z
}

func (z *Zsock) Connect(endpoint string) error {
	rc := C.Zsock_connect(z.zsock_t, C.CString(endpoint))
	if rc == C.int(-1) {
		return errors.New("failed")
	} else {
		return nil
	}
}

func (z *Zsock) Bind(endpoint string) error {
	rc := C.Zsock_bind(z.zsock_t, C.CString(endpoint))
	if rc == C.int(-1) {
		return errors.New("failed")
	} else {
		return nil
	}
}

type Zmsg struct {
	tag    uint32
	frames []string
}

func (z *Zsock) SendBytes(data []byte, flags Flag) error {
	frame := C.zframe_new(unsafe.Pointer(&data[0]), C.size_t(len(data)))
	rc := C.zframe_send(&frame, unsafe.Pointer(z.zsock_t), C.int(flags))
	if rc == C.int(-1) {
		return errors.New("failed")
	} else {
		return nil
	}
}

func (z *Zsock) RecvBytes() ([]byte, error) {
	frame := C.zframe_recv(unsafe.Pointer(z.zsock_t))
	size := C.zframe_size(frame)
	data := C.zframe_data(frame)
	return C.GoBytes(unsafe.Pointer(data), C.int(size)), nil
}

func (z *Zsock) Destroy() {
	C.zsock_destroy_(&z.zsock_t, C.CString(z.file), C.size_t(z.line))
}
