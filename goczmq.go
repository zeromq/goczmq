// Package goczmq is a golang binding for CZMQ 3. CZMQ is a high level binding
// for ZeroMQ. Along with ZeroMQ socket support, CZMQ provides "actor" based
// services for authentication, service discovery, and creating proxies.
// GoCZMQ provides direct bindings to CZMQ along with higher level go
// abstractions such as channels and io.ReadWriter interface support.
package goczmq

/*
#cgo pkg-config: libczmq libzmq libsodium
#include "czmq.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"errors"
)

const (
	REQ    = int(C.ZMQ_REQ)
	REP    = int(C.ZMQ_REP)
	DEALER = int(C.ZMQ_DEALER)
	ROUTER = int(C.ZMQ_ROUTER)
	PUB    = int(C.ZMQ_PUB)
	SUB    = int(C.ZMQ_SUB)
	XPUB   = int(C.ZMQ_XPUB)
	XSUB   = int(C.ZMQ_XSUB)
	PUSH   = int(C.ZMQ_PUSH)
	PULL   = int(C.ZMQ_PULL)
	PAIR   = int(C.ZMQ_PAIR)
	STREAM = int(C.ZMQ_STREAM)

	POLLIN  = int(C.ZMQ_POLLIN)
	POLLOUT = int(C.ZMQ_POLLOUT)

	ZMSG_TAG = 0x003cafe
	MORE     = int(C.ZFRAME_MORE)
	REUSE    = int(C.ZFRAME_REUSE)
	DONTWAIT = int(C.ZFRAME_DONTWAIT)

	CURVE_ALLOW_ANY = "*"
)

var (
	ErrActorCmd   = errors.New("error sending actor command")
	ErrSockAttach = errors.New("error attaching zsock")
)

func getStringType(k int) string {
	switch k {
	case REQ:
		return "REQ"
	case REP:
		return "REP"
	case DEALER:
		return "DEALER"
	case ROUTER:
		return "ROUTER"
	case PUB:
		return "PUB"
	case SUB:
		return "SUB"
	case XPUB:
		return "XPUB"
	case XSUB:
		return "XSUB"
	case PUSH:
		return "PUSH"
	case PULL:
		return "PULL"
	case PAIR:
		return "PAIR"
	case STREAM:
		return "STREAM"
	default:
		return ""
	}
}
