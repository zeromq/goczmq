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
	Req    = int(C.ZMQ_REQ)
	Rep    = int(C.ZMQ_REP)
	Dealer = int(C.ZMQ_DEALER)
	Router = int(C.ZMQ_ROUTER)
	Pub    = int(C.ZMQ_PUB)
	Sub    = int(C.ZMQ_SUB)
	XPub   = int(C.ZMQ_XPUB)
	XSub   = int(C.ZMQ_XSUB)
	Push   = int(C.ZMQ_PUSH)
	Pull   = int(C.ZMQ_PULL)
	Pair   = int(C.ZMQ_PAIR)
	Stream = int(C.ZMQ_STREAM)

	Pollin  = int(C.ZMQ_POLLIN)
	Pollout = int(C.ZMQ_POLLOUT)

	ZmsgTag  = 0x003cafe
	More     = int(C.ZFRAME_MORE)
	Reuse    = int(C.ZFRAME_REUSE)
	DontWait = int(C.ZFRAME_DONTWAIT)

	CurveAllowAny = "*"
)

var (
	ErrActorCmd   = errors.New("error sending actor command")
	ErrSockAttach = errors.New("error attaching zsock")
)

func getStringType(k int) string {
	switch k {
	case Req:
		return "REQ"
	case Rep:
		return "REP"
	case Dealer:
		return "DEALER"
	case Router:
		return "ROUTER"
	case Pub:
		return "PUB"
	case Sub:
		return "SUB"
	case XPub:
		return "XPUB"
	case XSub:
		return "XSUB"
	case Push:
		return "PUSH"
	case Pull:
		return "PULL"
	case Pair:
		return "PAIR"
	case Stream:
		return "STREAM"
	default:
		return ""
	}
}
