package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"

int Set_meta(zcert_t *self, const char *key, const char *value) {zcert_set_meta(self, key, value);}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// Cert holds a czmq zcert_t
type Cert struct {
	zcert_t *C.struct__zcert_t
}

// NewCert creates a new empty Cert instance
func NewCert() *Cert {
	return &Cert{
		zcert_t: C.zcert_new(),
	}
}

// NewCertFrom creates a new Cert from a public and private key
func NewCertFrom(public []byte, secret []byte) (*Cert, error) {
	if len(public) != 32 {
		return nil, fmt.Errorf("invalid public key")
	}

	if len(secret) != 32 {
		return nil, fmt.Errorf("invalid private key")
	}

	return &Cert{
		zcert_t: C.zcert_new_from(
			(*C.byte)(unsafe.Pointer(&public[0])),
			(*C.byte)(unsafe.Pointer(&secret[0]))),
	}, nil
}

// SetMeta sets meta data for a Cert
func (c *Cert) SetMeta(key string, value string) {
	C.Set_meta(c.zcert_t, C.CString(key), C.CString(value))
}

// Meta returns a meta data item from a Cert given a key
func (c *Cert) Meta(key string) string {
	val := C.zcert_meta(c.zcert_t, C.CString(key))
	return C.GoString(val)
}

// PublicText returns the public key as a string
func (c *Cert) PublicText() string {
	val := C.zcert_public_txt(c.zcert_t)
	return C.GoString(val)
}

// Apply sets the public and private keys for a socket
func (c *Cert) Apply(s *Sock) {
	handle := C.zsock_resolve(unsafe.Pointer(s.zsock_t))
	C.zsocket_set_curve_secretkey_bin(handle, C.zcert_secret_key(c.zcert_t))
	C.zsocket_set_curve_publickey_bin(handle, C.zcert_public_key(c.zcert_t))
}

// Destroy destroys Cert instance
func (c *Cert) Destroy() {
	C.zcert_destroy(&c.zcert_t)
}
