// A Go interface to CZMQ
package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"

zactor_t *Zproxy_new () { zactor_t *proxy = zactor_new(zproxy, NULL); return proxy; }
*/
import "C"

type Zproxy struct {
	zactor_t *C.struct__zactor_t
}

func NewZproxy() *Zproxy {
	z := &Zproxy{}
	z.zactor_t = C.Zproxy_new()
	return z
}

func (z *Zproxy) Destroy() {
	C.zactor_destroy(&z.zactor_t)
}
