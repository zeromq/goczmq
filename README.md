A go interface to [CZMQ](http://czmq.zeromq.org)

This requires CZMQ head, and is targetted to be compatible with the next stable release of CZMQ.

Development is currently using CZMQ head compiled against ZeroMQ 4.0.4 Stable.

## Install

### Required

* ZeroMQ 4.0.4 or higher ( http://zeromq.org/intro:get-the-software )
* CZMQ Head ( https://github.com/zeromq/czmq )

### Get Go Library

  go get github.com/zeromq/goczmq

## Status

This library is alpha.  Not all features are complete.  API changes will happen.

Currently implemented:

* ZSock
* ZProxy
* ZBeacon

## Goals

Todo to finish inital phase::

* ZAuth
* ZGossip
* ZLoop
* ZMonitor
* ZPoller

Secondary: Provide additional abstractions for "Go-isms" such as providing Zsocks as channel
accessable "services" within a go process.

## See Also

Peter Kleiweg's excellent zmq4 library for libzmq: http://github.com/pebbe/zmq4

## Smart Constructor Example
```go
package main

import (
	"flag"
	czmq "github.com/zeromq/goczmq"
	"log"
	"time"
)

func main() {
	var messageSize = flag.Int("message_size", 0, "size of message")
	var messageCount = flag.Int("message_count", 0, "number of messages")
	flag.Parse()

	pullSock, err := czmq.NewPULL("inproc://test")
	if err != nil {
		panic(err)
	}

	defer pullSock.Destroy()

	go func() {
		pushSock, err := czmq.NewPUSH("inproc://test")
		if err != nil {
			panic(err)
		}

		defer pushSock.Destroy()
		
		for i := 0; i < *messageCount; i++ {
			payload := make([]byte, *messageSize)
			err = pushSock.SendMessage(payload)
			if err != nil {
				panic(err)
			}
		}
	}()

	startTime := time.Now()
	for i := 0; i < *messageCount; i++ {
		msg, err := pullSock.RecvMessage()
		if err != nil {
			panic(err)
		}
		if len(msg) != 1 {
			panic("msg too small")
		}
	}
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	throughput := float64(*messageCount) / elapsed.Seconds()
	megabits := float64(throughput*float64(*messageSize)*8.0) / 1e6

	log.Printf("message size: %d", *messageSize)
	log.Printf("message count: %d", *messageCount)
	log.Printf("test time (seconds): %f", elapsed.Seconds())
	log.Printf("mean throughput: %f [msg/s]", throughput)
	log.Printf("mean throughput: %f [Mb/s]", megabits)
}
```

## Zbeacon Example
```go
package main

import (
	"fmt"
	czmq "github.com/taotetek/goczmq"
)

func main() {
	speaker := czmq.NewZbeacon()
	addr, err := speaker.Configure(9999)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Beacon configured on: %s\n", addr)

	listener := czmq.NewZbeacon()
	addr, err = listener.Configure(9999)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Beacon configured on: %s\n", addr)

	listener.Subscribe("HI")

	speaker.Publish("HI", 100)
	reply := listener.Recv(500)
	fmt.Printf("Received beacon: %v\n", reply)

	listener.Destroy()
	speaker.Destroy()
}
```

## Godoc

# goczmq
--
    import "github.com/taotetek/goczmq"

A Go Interface to CZMQ

## Usage

```go
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
```

```go
var (
	ErrActorCmd    = errors.New("error sending actor command")
	ErrZsockAttach = errors.New("error attaching zsock")
)
```

#### type Flag

```go
type Flag int
```


#### type Type

```go
type Type int
```


#### type Zbeacon

```go
type Zbeacon struct {
}
```


#### func  NewZbeacon

```go
func NewZbeacon() *Zbeacon
```
NewZbeacon creates a new Zbeacon instance.

#### func (*Zbeacon) Configure

```go
func (z *Zbeacon) Configure(port int) (string, error)
```
Configure accepts a port number and configures the beacon, returning an address

#### func (*Zbeacon) Destroy

```go
func (z *Zbeacon) Destroy()
```
Destroy destroys the beacon.

#### func (*Zbeacon) Publish

```go
func (z *Zbeacon) Publish(announcement string, interval int) error
```
Publish publishes an announcement at an interval

#### func (*Zbeacon) Recv

```go
func (z *Zbeacon) Recv(timeout int) string
```
Recv waits for the specific timeout in milliseconds to receive a beacon

#### func (*Zbeacon) Subscribe

```go
func (z *Zbeacon) Subscribe(filter string) error
```
Subscribe subscribes to beacons matching the filter

#### func (*Zbeacon) Verbose

```go
func (z *Zbeacon) Verbose() error
```
Verbose sets the beacon to log information to stdout.

#### type Zproxy

```go
type Zproxy struct {
}
```

Zproxy actors switch messages between a frontend and backend socket. The Zproxy
struct holds a reference to a CZMQ zactor_t.

#### func  NewZproxy

```go
func NewZproxy() *Zproxy
```
NewZproxy creates a new Zproxy instance.

#### func (*Zproxy) Destroy

```go
func (z *Zproxy) Destroy()
```
Destroy destroys the proxy.

#### func (*Zproxy) Pause

```go
func (z *Zproxy) Pause() error
```
Pause sends a message to the zproxy actor telling it to pause.

#### func (*Zproxy) Resume

```go
func (z *Zproxy) Resume() error
```
Resume sends a message to the zproxy actor telling it to resume.

#### func (*Zproxy) SetBackend

```go
func (z *Zproxy) SetBackend(sockType Type, endpoint string) error
```
SetBackend accepts a socket type and endpoint, and sends a message to the zactor
thread telling it to set up a socket bound to the endpoint.

#### func (*Zproxy) SetCapture

```go
func (z *Zproxy) SetCapture(endpoint string) error
```
SetCapture accepts a socket endpoint and sets up a PUSH socket bound to that
endpoint, that sends a copy of all messages passing through the proxy.

#### func (*Zproxy) SetFrontend

```go
func (z *Zproxy) SetFrontend(sockType Type, endpoint string) error
```
SetFrontend accepts a socket type and endpoint, and sends a message to the
zactor thread telling it to set up a socket bound to the endpoint.

#### func (*Zproxy) Verbose

```go
func (z *Zproxy) Verbose() error
```
Verbose sets the proxy to log information to stdout.

#### type Zsock

```go
type Zsock struct {
}
```

Zsock wraps the zsock_t class in CZMQ.

#### func  NewDEALER

```go
func NewDEALER(endpoints string) (*Zsock, error)
```
NewDEALER creates a DEALER socket. The endpoint is empty, or starts with '@'
(connect) or '>' (bind). Multiple endpoints are allowed, separated by commas. If
the endpoint does not start with '@' or '>', it connects.

#### func  NewPAIR

```go
func NewPAIR(endpoints string) (*Zsock, error)
```
NewPAIR creates a PAIR socket. The endpoint is empty, or starts with '@'
(connect) or '>' (bind). If the endpoint does not start with '@' or '>', it
connects.

#### func  NewPUB

```go
func NewPUB(endpoints string) (*Zsock, error)
```
NewPUB creates a PUB socket. The endpoint is empty, or starts with '@' (connect)
or '>' (bind). Multiple endpoints are allowed, separated by commas. If the
endpoint does not start with '@' or '>', it binds.

#### func  NewPULL

```go
func NewPULL(endpoints string) (*Zsock, error)
```
NewPULL creates a PULL socket. The endpoint is empty, or starts with '@'
(connect) or '>' (bind). Multiple endpoints are allowed, separated by commas. If
the endpoint does not start with '@' or '>', it binds.

#### func  NewPUSH

```go
func NewPUSH(endpoints string) (*Zsock, error)
```
NewPUSH creates a PUSH socket. The endpoint is empty, or starts with '@'
(connect) or '>' (bind). Multiple endpoints are allowed, separated by commas. If
the endpoint does not start with '@' or '>', it connects.

#### func  NewREP

```go
func NewREP(endpoints string) (*Zsock, error)
```
NewREP creates a REP socket. The endpoint is empty, or starts with '@' (connect)
or '>' (bind). Multiple endpoints are allowed, separated by commas. If the
endpoint does not start with '@' or '>', it binds.

#### func  NewREQ

```go
func NewREQ(endpoints string) (*Zsock, error)
```
NewREQ creates a REQ socket. The endpoint is empty, or starts with '@' (connect)
or '>' (bind). Multiple endpoints are allowed, separated by commas. If the
endpoint does not start with '@' or '>', it connects.

#### func  NewROUTER

```go
func NewROUTER(endpoints string) (*Zsock, error)
```
NewROUTER creates a ROUTER socket. The endpoint is empty, or starts with '@'
(connect) or '>' (bind). Multiple endpoints are allowed, separated by commas. If
the endpoint does not start with '@' or '>', it binds.

#### func  NewSTREAM

```go
func NewSTREAM(endpoints string) (*Zsock, error)
```
NewSTREAM creates a STREAM socket. The endpoint is empty, or starts with '@'
(connect) or '>' (bind). Multiple endpoints are allowed, separated by commas. If
the endpoint does not start with '@' or '>', it connects.

#### func  NewSUB

```go
func NewSUB(endpoints string, subscribe string) (*Zsock, error)
```
NewSUB creates a SUB socket. The enpoint is empty, or starts with '@' (connect)
or '>' (bind). Multiple endpoints are allowed, separated by commas. If the
endpoint does not start with '@' or '>', it connects. The second argument is a
comma delimited list of topics to subscribe to.

#### func  NewXPUB

```go
func NewXPUB(endpoints string) (*Zsock, error)
```
NewXPUB creates an XPUB socket. The endpoint is empty, or starts with '@'
(connect) or '>' (bind). Multiple endpoints are allowed, separated by commas. If
the endpoint does not start with '@' or '>', it binds.

#### func  NewXSUB

```go
func NewXSUB(endpoints string) (*Zsock, error)
```
NewXSUB creates an XSUB socket. The endpoint is empty, or starts with '@'
(connect) or '>' (bind). Multiple endpoints are allowed, separated by commas. If
the endpoint does not start with '@' or '>', it connects.

#### func  NewZsock

```go
func NewZsock(t Type) *Zsock
```
NewZsock creates a new socket. The caller source and line number are passed so
CZMQ can report socket leaks intelligently.

#### func (*Zsock) Affinity

```go
func (z *Zsock) Affinity() int
```
Affinity returns the current value of the socket's affinity option

#### func (*Zsock) Backlog

```go
func (z *Zsock) Backlog() int
```
Backlog returns the current value of the socket's backlog option

#### func (*Zsock) Bind

```go
func (z *Zsock) Bind(endpoint string) (int, error)
```
Bind binds a socket to an endpoint. On success returns the port number used for
tcp transports, or 0 for other transports. On failure returns a -1 for port, and
an error.

#### func (*Zsock) Connect

```go
func (z *Zsock) Connect(endpoint string) error
```
Connect connects a socket to an endpoint returns an error if the connect failed.

#### func (*Zsock) CurvePublickey

```go
func (z *Zsock) CurvePublickey() string
```
CurvePublickey returns the current value of the socket's curve_publickey option

#### func (*Zsock) CurveSecretkey

```go
func (z *Zsock) CurveSecretkey() string
```
CurveSecretkey returns the current value of the socket's curve_secretkey option

#### func (*Zsock) CurveServer

```go
func (z *Zsock) CurveServer() int
```
CurveServer returns the current value of the socket's curve_server option

#### func (*Zsock) CurveServerkey

```go
func (z *Zsock) CurveServerkey() string
```
CurveServerkey returns the current value of the socket's curve_serverkey option

#### func (*Zsock) Destroy

```go
func (z *Zsock) Destroy()
```
Destroy destroys the underlying zsock_t.

#### func (*Zsock) Events

```go
func (z *Zsock) Events() int
```
Events returns the current value of the socket's events option

#### func (*Zsock) Fd

```go
func (z *Zsock) Fd() int
```
Fd returns the current value of the socket's fd option

#### func (*Zsock) GssapiPlaintext

```go
func (z *Zsock) GssapiPlaintext() int
```
GssapiPlaintext returns the current value of the socket's gssapi_plaintext
option

#### func (*Zsock) GssapiPrincipal

```go
func (z *Zsock) GssapiPrincipal() string
```
GssapiPrincipal returns the current value of the socket's gssapi_principal
option

#### func (*Zsock) GssapiServer

```go
func (z *Zsock) GssapiServer() int
```
GssapiServer returns the current value of the socket's gssapi_server option

#### func (*Zsock) GssapiServicePrincipal

```go
func (z *Zsock) GssapiServicePrincipal() string
```
GssapiServicePrincipal returns the current value of the socket's
gssapi_service_principal option

#### func (*Zsock) Identity

```go
func (z *Zsock) Identity() string
```
Identity returns the current value of the socket's identity option

#### func (*Zsock) Immediate

```go
func (z *Zsock) Immediate() int
```
Immediate returns the current value of the socket's immediate option

#### func (*Zsock) Ipv4only

```go
func (z *Zsock) Ipv4only() int
```
Ipv4only returns the current value of the socket's ipv4only option

#### func (*Zsock) Ipv6

```go
func (z *Zsock) Ipv6() int
```
Ipv6 returns the current value of the socket's ipv6 option

#### func (*Zsock) LastEndpoint

```go
func (z *Zsock) LastEndpoint() string
```
LastEndpoint returns the current value of the socket's last_endpoint option

#### func (*Zsock) Linger

```go
func (z *Zsock) Linger() int
```
Linger returns the current value of the socket's linger option

#### func (*Zsock) Maxmsgsize

```go
func (z *Zsock) Maxmsgsize() int
```
Maxmsgsize returns the current value of the socket's maxmsgsize option

#### func (*Zsock) Mechanism

```go
func (z *Zsock) Mechanism() int
```
Mechanism returns the current value of the socket's mechanism option

#### func (*Zsock) MulticastHops

```go
func (z *Zsock) MulticastHops() int
```
MulticastHops returns the current value of the socket's multicast_hops option

#### func (*Zsock) PlainPassword

```go
func (z *Zsock) PlainPassword() string
```
PlainPassword returns the current value of the socket's plain_password option

#### func (*Zsock) PlainServer

```go
func (z *Zsock) PlainServer() int
```
PlainServer returns the current value of the socket's plain_server option

#### func (*Zsock) PlainUsername

```go
func (z *Zsock) PlainUsername() string
```
PlainUsername returns the current value of the socket's plain_username option

#### func (*Zsock) Rate

```go
func (z *Zsock) Rate() int
```
Rate returns the current value of the socket's rate option

#### func (*Zsock) Rcvbuf

```go
func (z *Zsock) Rcvbuf() int
```
Rcvbuf returns the current value of the socket's rcvbuf option

#### func (*Zsock) Rcvhwm

```go
func (z *Zsock) Rcvhwm() int
```
Rcvhwm returns the current value of the socket's rcvhwm option

#### func (*Zsock) Rcvmore

```go
func (z *Zsock) Rcvmore() int
```
Rcvmore returns the current value of the socket's rcvmore option

#### func (*Zsock) Rcvtimeo

```go
func (z *Zsock) Rcvtimeo() int
```
Rcvtimeo returns the current value of the socket's rcvtimeo option

#### func (*Zsock) ReconnectIvl

```go
func (z *Zsock) ReconnectIvl() int
```
ReconnectIvl returns the current value of the socket's reconnect_ivl option

#### func (*Zsock) ReconnectIvlMax

```go
func (z *Zsock) ReconnectIvlMax() int
```
ReconnectIvlMax returns the current value of the socket's reconnect_ivl_max
option

#### func (*Zsock) RecoveryIvl

```go
func (z *Zsock) RecoveryIvl() int
```
RecoveryIvl returns the current value of the socket's recovery_ivl option

#### func (*Zsock) RecvBytes

```go
func (z *Zsock) RecvBytes() ([]byte, Flag, error)
```
RecvBytes reads a frame from the socket and returns it as a byte array, Returns
an error if the call fails.

#### func (*Zsock) RecvMessage

```go
func (z *Zsock) RecvMessage() ([][]byte, error)
```
RecvMessage receives a full message from the socket and returns it as an array
of byte arrays.

#### func (*Zsock) RecvString

```go
func (z *Zsock) RecvString() (string, error)
```
RecvString reads a frame from the socket and returns it as a string, Returns an
error if the call fails.

#### func (*Zsock) SendBytes

```go
func (z *Zsock) SendBytes(data []byte, flags Flag) error
```
SendBytes sends a byte array via the socket. For the flags value, use 0 for a
single message, or SNDMORE if it is a multi-part message

#### func (*Zsock) SendMessage

```go
func (z *Zsock) SendMessage(parts ...interface{}) error
```
SendMessage is a variadic function that currently accepts ints, strings, and
bytes, and sends them as an atomic multi frame message over zeromq as a series
of byte arrays. In the case of numeric data, the resulting byte array is a
textual representation of the number (e.g., 100 turns to "100"). This may be
changed to network byte ordered representation in the near future - I have not
decided yet!

#### func (*Zsock) SendString

```go
func (z *Zsock) SendString(data string, flags Flag) error
```
SendString sends a string via the socket. For the flags value, use 0 for a
single message, or SNDMORE if it is a multi-part message

#### func (*Zsock) SetAffinity

```go
func (z *Zsock) SetAffinity(val int)
```
SetAffinity sets the affinity option for the socket

#### func (*Zsock) SetBacklog

```go
func (z *Zsock) SetBacklog(val int)
```
SetBacklog sets the backlog option for the socket

#### func (*Zsock) SetConflate

```go
func (z *Zsock) SetConflate(val int)
```
SetConflate sets the conflate option for the socket

#### func (*Zsock) SetCurvePublickey

```go
func (z *Zsock) SetCurvePublickey(val string)
```
SetCurvePublickey sets the curve_publickey option for the socket

#### func (*Zsock) SetCurveSecretkey

```go
func (z *Zsock) SetCurveSecretkey(val string)
```
SetCurveSecretkey sets the curve_secretkey option for the socket

#### func (*Zsock) SetCurveServer

```go
func (z *Zsock) SetCurveServer(val int)
```
SetCurveServer sets the curve_server option for the socket

#### func (*Zsock) SetCurveServerkey

```go
func (z *Zsock) SetCurveServerkey(val string)
```
SetCurveServerkey sets the curve_serverkey option for the socket

#### func (*Zsock) SetDelayAttachOnConnect

```go
func (z *Zsock) SetDelayAttachOnConnect(val int)
```
SetDelayAttachOnConnect sets the delay_attach_on_connect option for the socket

#### func (*Zsock) SetGssapiPlaintext

```go
func (z *Zsock) SetGssapiPlaintext(val int)
```
SetGssapiPlaintext sets the gssapi_plaintext option for the socket

#### func (*Zsock) SetGssapiPrincipal

```go
func (z *Zsock) SetGssapiPrincipal(val string)
```
SetGssapiPrincipal sets the gssapi_principal option for the socket

#### func (*Zsock) SetGssapiServer

```go
func (z *Zsock) SetGssapiServer(val int)
```
SetGssapiServer sets the gssapi_server option for the socket

#### func (*Zsock) SetGssapiServicePrincipal

```go
func (z *Zsock) SetGssapiServicePrincipal(val string)
```
SetGssapiServicePrincipal sets the gssapi_service_principal option for the
socket

#### func (*Zsock) SetIdentity

```go
func (z *Zsock) SetIdentity(val string)
```
SetIdentity sets the identity option for the socket

#### func (*Zsock) SetImmediate

```go
func (z *Zsock) SetImmediate(val int)
```
SetImmediate sets the immediate option for the socket

#### func (*Zsock) SetIpv4only

```go
func (z *Zsock) SetIpv4only(val int)
```
SetIpv4only sets the ipv4only option for the socket

#### func (*Zsock) SetIpv6

```go
func (z *Zsock) SetIpv6(val int)
```
SetIpv6 sets the ipv6 option for the socket

#### func (*Zsock) SetLinger

```go
func (z *Zsock) SetLinger(val int)
```
SetLinger sets the linger option for the socket

#### func (*Zsock) SetMaxmsgsize

```go
func (z *Zsock) SetMaxmsgsize(val int)
```
SetMaxmsgsize sets the maxmsgsize option for the socket

#### func (*Zsock) SetMulticastHops

```go
func (z *Zsock) SetMulticastHops(val int)
```
SetMulticastHops sets the multicast_hops option for the socket

#### func (*Zsock) SetPlainPassword

```go
func (z *Zsock) SetPlainPassword(val string)
```
SetPlainPassword sets the plain_password option for the socket

#### func (*Zsock) SetPlainServer

```go
func (z *Zsock) SetPlainServer(val int)
```
SetPlainServer sets the plain_server option for the socket

#### func (*Zsock) SetPlainUsername

```go
func (z *Zsock) SetPlainUsername(val string)
```
SetPlainUsername sets the plain_username option for the socket

#### func (*Zsock) SetProbeRouter

```go
func (z *Zsock) SetProbeRouter(val int)
```
SetProbeRouter sets the probe_router option for the socket

#### func (*Zsock) SetRate

```go
func (z *Zsock) SetRate(val int)
```
SetRate sets the rate option for the socket

#### func (*Zsock) SetRcvbuf

```go
func (z *Zsock) SetRcvbuf(val int)
```
SetRcvbuf sets the rcvbuf option for the socket

#### func (*Zsock) SetRcvhwm

```go
func (z *Zsock) SetRcvhwm(val int)
```
SetRcvhwm sets the rcvhwm option for the socket

#### func (*Zsock) SetRcvtimeo

```go
func (z *Zsock) SetRcvtimeo(val int)
```
SetRcvtimeo sets the rcvtimeo option for the socket

#### func (*Zsock) SetReconnectIvl

```go
func (z *Zsock) SetReconnectIvl(val int)
```
SetReconnectIvl sets the reconnect_ivl option for the socket

#### func (*Zsock) SetReconnectIvlMax

```go
func (z *Zsock) SetReconnectIvlMax(val int)
```
SetReconnectIvlMax sets the reconnect_ivl_max option for the socket

#### func (*Zsock) SetRecoveryIvl

```go
func (z *Zsock) SetRecoveryIvl(val int)
```
SetRecoveryIvl sets the recovery_ivl option for the socket

#### func (*Zsock) SetReqCorrelate

```go
func (z *Zsock) SetReqCorrelate(val int)
```
SetReqCorrelate sets the req_correlate option for the socket

#### func (*Zsock) SetReqRelaxed

```go
func (z *Zsock) SetReqRelaxed(val int)
```
SetReqRelaxed sets the req_relaxed option for the socket

#### func (*Zsock) SetRouterHandover

```go
func (z *Zsock) SetRouterHandover(val int)
```
SetRouterHandover sets the router_handover option for the socket

#### func (*Zsock) SetRouterMandatory

```go
func (z *Zsock) SetRouterMandatory(val int)
```
SetRouterMandatory sets the router_mandatory option for the socket

#### func (*Zsock) SetRouterRaw

```go
func (z *Zsock) SetRouterRaw(val int)
```
SetRouterRaw sets the router_raw option for the socket

#### func (*Zsock) SetSndbuf

```go
func (z *Zsock) SetSndbuf(val int)
```
SetSndbuf sets the sndbuf option for the socket

#### func (*Zsock) SetSndhwm

```go
func (z *Zsock) SetSndhwm(val int)
```
SetSndhwm sets the sndhwm option for the socket

#### func (*Zsock) SetSndtimeo

```go
func (z *Zsock) SetSndtimeo(val int)
```
SetSndtimeo sets the sndtimeo option for the socket

#### func (*Zsock) SetSubscribe

```go
func (z *Zsock) SetSubscribe(val string)
```
SetSubscribe sets the subscribe option for the socket

#### func (*Zsock) SetTcpAcceptFilter

```go
func (z *Zsock) SetTcpAcceptFilter(val string)
```
SetTcpAcceptFilter sets the tcp_accept_filter option for the socket

#### func (*Zsock) SetTcpKeepalive

```go
func (z *Zsock) SetTcpKeepalive(val int)
```
SetTcpKeepalive sets the tcp_keepalive option for the socket

#### func (*Zsock) SetTcpKeepaliveCnt

```go
func (z *Zsock) SetTcpKeepaliveCnt(val int)
```
SetTcpKeepaliveCnt sets the tcp_keepalive_cnt option for the socket

#### func (*Zsock) SetTcpKeepaliveIdle

```go
func (z *Zsock) SetTcpKeepaliveIdle(val int)
```
SetTcpKeepaliveIdle sets the tcp_keepalive_idle option for the socket

#### func (*Zsock) SetTcpKeepaliveIntvl

```go
func (z *Zsock) SetTcpKeepaliveIntvl(val int)
```
SetTcpKeepaliveIntvl sets the tcp_keepalive_intvl option for the socket

#### func (*Zsock) SetTos

```go
func (z *Zsock) SetTos(val int)
```
SetTos sets the tos option for the socket

#### func (*Zsock) SetUnsubscribe

```go
func (z *Zsock) SetUnsubscribe(val string)
```
SetUnsubscribe sets the unsubscribe option for the socket

#### func (*Zsock) SetXpubVerbose

```go
func (z *Zsock) SetXpubVerbose(val int)
```
SetXpubVerbose sets the xpub_verbose option for the socket

#### func (*Zsock) SetZapDomain

```go
func (z *Zsock) SetZapDomain(val string)
```
SetZapDomain sets the zap_domain option for the socket

#### func (*Zsock) Sndbuf

```go
func (z *Zsock) Sndbuf() int
```
Sndbuf returns the current value of the socket's sndbuf option

#### func (*Zsock) Sndhwm

```go
func (z *Zsock) Sndhwm() int
```
Sndhwm returns the current value of the socket's sndhwm option

#### func (*Zsock) Sndtimeo

```go
func (z *Zsock) Sndtimeo() int
```
Sndtimeo returns the current value of the socket's sndtimeo option

#### func (*Zsock) TcpAcceptFilter

```go
func (z *Zsock) TcpAcceptFilter() string
```
TcpAcceptFilter returns the current value of the socket's tcp_accept_filter
option

#### func (*Zsock) TcpKeepalive

```go
func (z *Zsock) TcpKeepalive() int
```
TcpKeepalive returns the current value of the socket's tcp_keepalive option

#### func (*Zsock) TcpKeepaliveCnt

```go
func (z *Zsock) TcpKeepaliveCnt() int
```
TcpKeepaliveCnt returns the current value of the socket's tcp_keepalive_cnt
option

#### func (*Zsock) TcpKeepaliveIdle

```go
func (z *Zsock) TcpKeepaliveIdle() int
```
TcpKeepaliveIdle returns the current value of the socket's tcp_keepalive_idle
option

#### func (*Zsock) TcpKeepaliveIntvl

```go
func (z *Zsock) TcpKeepaliveIntvl() int
```
TcpKeepaliveIntvl returns the current value of the socket's tcp_keepalive_intvl
option

#### func (*Zsock) Tos

```go
func (z *Zsock) Tos() int
```
Tos returns the current value of the socket's tos option

#### func (*Zsock) Type

```go
func (z *Zsock) Type() int
```
Type returns the current value of the socket's type option

#### func (*Zsock) ZapDomain

```go
func (z *Zsock) ZapDomain() string
```
ZapDomain returns the current value of the socket's zap_domain option
