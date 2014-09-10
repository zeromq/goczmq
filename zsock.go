// A Go interface to CZMQ
package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"
#include <stdlib.h>
#include <string.h>

int Zsock_connect(zsock_t *self, const char *format) {return zsock_connect(self, format, NULL);}
int Zsock_bind(zsock_t *self, const char *format) {return zsock_bind(self, format, NULL);}
*/
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
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
	if frame == nil {
		return []byte{0}, errors.New("failed")
	}
	dataSize := C.zframe_size(frame)
	dataPtr := C.zframe_data(frame)
	b := C.GoBytes(unsafe.Pointer(dataPtr), C.int(dataSize))
	C.zframe_destroy(&frame)
	return b, nil
}

func (z *Zsock) Destroy() {
	C.zsock_destroy_(&z.zsock_t, C.CString(z.file), C.size_t(z.line))
}
