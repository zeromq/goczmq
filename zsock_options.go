package goczmq

/*
#include "czmq.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"unsafe"
)

func (z *Zsock) SetSndhwm(val int) {
	C.zsock_set_sndhwm(unsafe.Pointer(z.zsock_t), C.int(val))
}

func (z *Zsock) Sndhwm() int {
	val := C.zsock_sndhwm(unsafe.Pointer(z.zsock_t))
	return int(val)
}

func (z *Zsock) SetRcvhwm(val int) {
	C.zsock_set_rcvhwm(unsafe.Pointer(z.zsock_t), C.int(val))
}

func (z *Zsock) Rcvhwm() int {
	val := C.zsock_rcvhwm(unsafe.Pointer(z.zsock_t))
	return int(val)
}
