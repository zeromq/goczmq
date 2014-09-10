// A Go interface to CZMQ
package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"

zactor_t *Zproxy_new () { zactor_t *proxy = zactor_new(zproxy, NULL); return proxy; }
int send_proxy_command(void *dest, const char *command, const char *socktype, const char *endpoint) { return zstr_sendx(dest, command, socktype, endpoint, NULL); }
*/
import "C"

import "unsafe"

type Zproxy struct {
	zactor_t *C.struct__zactor_t
}

func NewZproxy() *Zproxy {
	z := &Zproxy{}
	z.zactor_t = C.Zproxy_new()
	return z
}

func (z *Zproxy) SetFrontend(sockType Type, endpoint string) error {
	typeString := getStringType(sockType)
	C.send_proxy_command(unsafe.Pointer(z.zactor_t), C.CString("FRONTEND"), C.CString(typeString), C.CString(endpoint))
	C.zsock_wait(unsafe.Pointer(z.zactor_t))
	return nil
}

func (z *Zproxy) SetBackend(sockType Type, endpoint string) error {
	typeString := getStringType(sockType)
	C.send_proxy_command(unsafe.Pointer(z.zactor_t), C.CString("BACKEND"), C.CString(typeString), C.CString(endpoint))
	C.zsock_wait(unsafe.Pointer(z.zactor_t))
	return nil
}

func (z *Zproxy) Destroy() {
	C.zactor_destroy(&z.zactor_t)
}
