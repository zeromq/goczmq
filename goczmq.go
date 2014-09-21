// A Go Interface to CZMQ
package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"errors"
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

	POLLIN  = int(C.ZMQ_POLLIN)
	POLLOUT = int(C.ZMQ_POLLOUT)

	ZMSG_TAG = 0x003cafe
	MORE     = Flag(C.ZFRAME_MORE)
	REUSE    = Flag(C.ZFRAME_REUSE)
	DONTWAIT = Flag(C.ZFRAME_DONTWAIT)
)

var (
	ErrActorCmd    = errors.New("error sending actor command")
	ErrZsockAttach = errors.New("error attaching zsock")
)

func getStringType(k Type) string {
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
