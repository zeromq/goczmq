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
* ZPoller

## Goals

Todo to finish inital phase::

* ZAuth
* ZGossip
* ZLoop
* ZMonitor

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

	POLLIN  = int(C.ZMQ_POLLIN)
	POLLOUT = int(C.ZMQ_POLLOUT)

	ZMSG_TAG = 0x003cafe
	MORE     = Flag(C.ZFRAME_MORE)
	REUSE    = Flag(C.ZFRAME_REUSE)
	DONTWAIT = Flag(C.ZFRAME_DONTWAIT)
)
```

```go
var (
	ErrActorCmd   = errors.New("error sending actor command")
	ErrSockAttach = errors.New("error attaching zsock")
)
```

#### type Beacon

```go
type Beacon struct {
}
```


#### func  NewBeacon

```go
func NewBeacon() *Beacon
```
NewBeacon creates a new Beacon instance.

#### func (*Beacon) Configure

```go
func (z *Beacon) Configure(port int) (string, error)
```
Configure accepts a port number and configures the beacon, returning an address

#### func (*Beacon) Destroy

```go
func (z *Beacon) Destroy()
```
Destroy destroys the beacon.

#### func (*Beacon) Publish

```go
func (z *Beacon) Publish(announcement string, interval int) error
```
Publish publishes an announcement at an interval

#### func (*Beacon) Recv

```go
func (z *Beacon) Recv(timeout int) string
```
Recv waits for the specific timeout in milliseconds to receive a beacon

#### func (*Beacon) Subscribe

```go
func (z *Beacon) Subscribe(filter string) error
```
Subscribe subscribes to beacons matching the filter

#### func (*Beacon) Verbose

```go
func (z *Beacon) Verbose() error
```
Verbose sets the beacon to log information to stdout.

#### type Flag

```go
type Flag int
```


#### type Gossip

```go
type Gossip struct {
}
```

Gossip actors use a gossip protocol for decentralized configuration management.
Gossip nodes form a loosely connected network that publishes and redistributed
name/value tuples. A network of Gossip actors will eventually achieve a
consistent state

#### func  NewGossip

```go
func NewGossip(name string) *Gossip
```
NewGossip creates a new Gossip actor

#### func (*Gossip) Bind

```go
func (z *Gossip) Bind(endpoint string) error
```
Bind binds the gossip service to a specified endpoint

#### func (*Gossip) Destroy

```go
func (z *Gossip) Destroy()
```
Destroy destroys the gossip actor.

#### func (*Gossip) Verbose

```go
func (z *Gossip) Verbose() error
```
Verbose sets the gossip actor to log information to stdout.

#### type Poller

```go
type Poller struct {
}
```

Poller is a simple poller for Socks

#### func  NewPoller

```go
func NewPoller(readers ...*Sock) (*Poller, error)
```
NewPoller creates a new Poller instance. It accepts one or more readers to poll.

#### func (*Poller) Add

```go
func (p *Poller) Add(reader *Sock) error
```
Add adds a reader to be polled.

#### func (*Poller) Destroy

```go
func (p *Poller) Destroy()
```
Destroy destroys the Poller

#### func (*Poller) Remove

```go
func (p *Poller) Remove(reader *Sock)
```
Remove removes a zsock from the poller

#### func (*Poller) Wait

```go
func (p *Poller) Wait(timeout int) *Sock
```
Wait waits for the timeout period in milliseconds for a POLLIN event, and
returns the first socket that returns one

#### type Proxy

```go
type Proxy struct {
}
```

Proxy actors switch messages between a frontend and backend socket. The Proxy
struct holds a reference to a CZMQ zactor_t.

#### func  NewProxy

```go
func NewProxy() *Proxy
```
NewProxy creates a new Proxy instance.

#### func (*Proxy) Destroy

```go
func (p *Proxy) Destroy()
```
Destroy destroys the proxy.

#### func (*Proxy) Pause

```go
func (p *Proxy) Pause() error
```
Pause sends a message to the zproxy actor telling it to pause.

#### func (*Proxy) Resume

```go
func (p *Proxy) Resume() error
```
Resume sends a message to the zproxy actor telling it to resume.

#### func (*Proxy) SetBackend

```go
func (p *Proxy) SetBackend(sockType Type, endpoint string) error
```
SetBackend accepts a socket type and endpoint, and sends a message to the zactor
thread telling it to set up a socket bound to the endpoint.

#### func (*Proxy) SetCapture

```go
func (p *Proxy) SetCapture(endpoint string) error
```
SetCapture accepts a socket endpoint and sets up a PUSH socket bound to that
endpoint, that sends a copy of all messages passing through the proxy.

#### func (*Proxy) SetFrontend

```go
func (p *Proxy) SetFrontend(sockType Type, endpoint string) error
```
SetFrontend accepts a socket type and endpoint, and sends a message to the
zactor thread telling it to set up a socket bound to the endpoint.

#### func (*Proxy) Verbose

```go
func (p *Proxy) Verbose() error
```
Verbose sets the proxy to log information to stdout.

#### type Sock

```go
type Sock struct {
}
```

Sock wraps the zsock_t class in CZMQ.

#### func  NewDEALER

```go
func NewDEALER(endpoints string) (*Sock, error)
```
NewDEALER creates a DEALER socket. The endpoint is empty, or starts with '@'
(connect) or '>' (bind). Multiple endpoints are allowed, separated by commas. If
the endpoint does not start with '@' or '>', it connects.

#### func  NewPAIR

```go
func NewPAIR(endpoints string) (*Sock, error)
```
NewPAIR creates a PAIR socket. The endpoint is empty, or starts with '@'
(connect) or '>' (bind). If the endpoint does not start with '@' or '>', it
connects.

#### func  NewPUB

```go
func NewPUB(endpoints string) (*Sock, error)
```
NewPUB creates a PUB socket. The endpoint is empty, or starts with '@' (connect)
or '>' (bind). Multiple endpoints are allowed, separated by commas. If the
endpoint does not start with '@' or '>', it binds.

#### func  NewPULL

```go
func NewPULL(endpoints string) (*Sock, error)
```
NewPULL creates a PULL socket. The endpoint is empty, or starts with '@'
(connect) or '>' (bind). Multiple endpoints are allowed, separated by commas. If
the endpoint does not start with '@' or '>', it binds.

#### func  NewPUSH

```go
func NewPUSH(endpoints string) (*Sock, error)
```
NewPUSH creates a PUSH socket. The endpoint is empty, or starts with '@'
(connect) or '>' (bind). Multiple endpoints are allowed, separated by commas. If
the endpoint does not start with '@' or '>', it connects.

#### func  NewREP

```go
func NewREP(endpoints string) (*Sock, error)
```
NewREP creates a REP socket. The endpoint is empty, or starts with '@' (connect)
or '>' (bind). Multiple endpoints are allowed, separated by commas. If the
endpoint does not start with '@' or '>', it binds.

#### func  NewREQ

```go
func NewREQ(endpoints string) (*Sock, error)
```
NewREQ creates a REQ socket. The endpoint is empty, or starts with '@' (connect)
or '>' (bind). Multiple endpoints are allowed, separated by commas. If the
endpoint does not start with '@' or '>', it connects.

#### func  NewROUTER

```go
func NewROUTER(endpoints string) (*Sock, error)
```
NewROUTER creates a ROUTER socket. The endpoint is empty, or starts with '@'
(connect) or '>' (bind). Multiple endpoints are allowed, separated by commas. If
the endpoint does not start with '@' or '>', it binds.

#### func  NewSTREAM

```go
func NewSTREAM(endpoints string) (*Sock, error)
```
NewSTREAM creates a STREAM socket. The endpoint is empty, or starts with '@'
(connect) or '>' (bind). Multiple endpoints are allowed, separated by commas. If
the endpoint does not start with '@' or '>', it connects.

#### func  NewSUB

```go
func NewSUB(endpoints string, subscribe string) (*Sock, error)
```
NewSUB creates a SUB socket. The enpoint is empty, or starts with '@' (connect)
or '>' (bind). Multiple endpoints are allowed, separated by commas. If the
endpoint does not start with '@' or '>', it connects. The second argument is a
comma delimited list of topics to subscribe to.

#### func  NewSock

```go
func NewSock(t Type) *Sock
```
NewSock creates a new socket. The caller source and line number are passed so
CZMQ can report socket leaks intelligently.

#### func  NewXPUB

```go
func NewXPUB(endpoints string) (*Sock, error)
```
NewXPUB creates an XPUB socket. The endpoint is empty, or starts with '@'
(connect) or '>' (bind). Multiple endpoints are allowed, separated by commas. If
the endpoint does not start with '@' or '>', it binds.

#### func  NewXSUB

```go
func NewXSUB(endpoints string) (*Sock, error)
```
NewXSUB creates an XSUB socket. The endpoint is empty, or starts with '@'
(connect) or '>' (bind). Multiple endpoints are allowed, separated by commas. If
the endpoint does not start with '@' or '>', it connects.

#### func (*Sock) Affinity

```go
func (z *Sock) Affinity() int
```
Affinity returns the current value of the socket's affinity option

#### func (*Sock) Backlog

```go
func (z *Sock) Backlog() int
```
Backlog returns the current value of the socket's backlog option

#### func (*Sock) Bind

```go
func (s *Sock) Bind(endpoint string) (int, error)
```
Bind binds a socket to an endpoint. On success returns the port number used for
tcp transports, or 0 for other transports. On failure returns a -1 for port, and
an error.

#### func (*Sock) Connect

```go
func (s *Sock) Connect(endpoint string) error
```
Connect connects a socket to an endpoint returns an error if the connect failed.

#### func (*Sock) CurvePublickey

```go
func (z *Sock) CurvePublickey() string
```
CurvePublickey returns the current value of the socket's curve_publickey option

#### func (*Sock) CurveSecretkey

```go
func (z *Sock) CurveSecretkey() string
```
CurveSecretkey returns the current value of the socket's curve_secretkey option

#### func (*Sock) CurveServer

```go
func (z *Sock) CurveServer() int
```
CurveServer returns the current value of the socket's curve_server option

#### func (*Sock) CurveServerkey

```go
func (z *Sock) CurveServerkey() string
```
CurveServerkey returns the current value of the socket's curve_serverkey option

#### func (*Sock) Destroy

```go
func (s *Sock) Destroy()
```
Destroy destroys the underlying zsock_t.

#### func (*Sock) Disconnect

```go
func (s *Sock) Disconnect(endpoint string) error
```
Disconnect disconnects a socket from an endpoint. If returns an error if the
endpoint was not found

#### func (*Sock) Events

```go
func (z *Sock) Events() int
```
Events returns the current value of the socket's events option

#### func (*Sock) Fd

```go
func (z *Sock) Fd() int
```
Fd returns the current value of the socket's fd option

#### func (*Sock) GssapiPlaintext

```go
func (z *Sock) GssapiPlaintext() int
```
GssapiPlaintext returns the current value of the socket's gssapi_plaintext
option

#### func (*Sock) GssapiPrincipal

```go
func (z *Sock) GssapiPrincipal() string
```
GssapiPrincipal returns the current value of the socket's gssapi_principal
option

#### func (*Sock) GssapiServer

```go
func (z *Sock) GssapiServer() int
```
GssapiServer returns the current value of the socket's gssapi_server option

#### func (*Sock) GssapiServicePrincipal

```go
func (z *Sock) GssapiServicePrincipal() string
```
GssapiServicePrincipal returns the current value of the socket's
gssapi_service_principal option

#### func (*Sock) Identity

```go
func (z *Sock) Identity() string
```
Identity returns the current value of the socket's identity option

#### func (*Sock) Immediate

```go
func (z *Sock) Immediate() int
```
Immediate returns the current value of the socket's immediate option

#### func (*Sock) Ipv4only

```go
func (z *Sock) Ipv4only() int
```
Ipv4only returns the current value of the socket's ipv4only option

#### func (*Sock) Ipv6

```go
func (z *Sock) Ipv6() int
```
Ipv6 returns the current value of the socket's ipv6 option

#### func (*Sock) LastEndpoint

```go
func (z *Sock) LastEndpoint() string
```
LastEndpoint returns the current value of the socket's last_endpoint option

#### func (*Sock) Linger

```go
func (z *Sock) Linger() int
```
Linger returns the current value of the socket's linger option

#### func (*Sock) Maxmsgsize

```go
func (z *Sock) Maxmsgsize() int
```
Maxmsgsize returns the current value of the socket's maxmsgsize option

#### func (*Sock) Mechanism

```go
func (z *Sock) Mechanism() int
```
Mechanism returns the current value of the socket's mechanism option

#### func (*Sock) MulticastHops

```go
func (z *Sock) MulticastHops() int
```
MulticastHops returns the current value of the socket's multicast_hops option

#### func (*Sock) PlainPassword

```go
func (z *Sock) PlainPassword() string
```
PlainPassword returns the current value of the socket's plain_password option

#### func (*Sock) PlainServer

```go
func (z *Sock) PlainServer() int
```
PlainServer returns the current value of the socket's plain_server option

#### func (*Sock) PlainUsername

```go
func (z *Sock) PlainUsername() string
```
PlainUsername returns the current value of the socket's plain_username option

#### func (*Sock) Pollin

```go
func (s *Sock) Pollin() bool
```
Pollin returns true if there is a POLLIN event on the socket

#### func (*Sock) Pollout

```go
func (s *Sock) Pollout() bool
```
Pollout returns true if there is a POLLOUT event on the socket

#### func (*Sock) Rate

```go
func (z *Sock) Rate() int
```
Rate returns the current value of the socket's rate option

#### func (*Sock) Rcvbuf

```go
func (z *Sock) Rcvbuf() int
```
Rcvbuf returns the current value of the socket's rcvbuf option

#### func (*Sock) Rcvhwm

```go
func (z *Sock) Rcvhwm() int
```
Rcvhwm returns the current value of the socket's rcvhwm option

#### func (*Sock) Rcvmore

```go
func (z *Sock) Rcvmore() int
```
Rcvmore returns the current value of the socket's rcvmore option

#### func (*Sock) Rcvtimeo

```go
func (z *Sock) Rcvtimeo() int
```
Rcvtimeo returns the current value of the socket's rcvtimeo option

#### func (*Sock) ReconnectIvl

```go
func (z *Sock) ReconnectIvl() int
```
ReconnectIvl returns the current value of the socket's reconnect_ivl option

#### func (*Sock) ReconnectIvlMax

```go
func (z *Sock) ReconnectIvlMax() int
```
ReconnectIvlMax returns the current value of the socket's reconnect_ivl_max
option

#### func (*Sock) RecoveryIvl

```go
func (z *Sock) RecoveryIvl() int
```
RecoveryIvl returns the current value of the socket's recovery_ivl option

#### func (*Sock) RecvBytes

```go
func (s *Sock) RecvBytes() ([]byte, Flag, error)
```
RecvBytes reads a frame from the socket and returns it as a byte array, Returns
an error if the call fails.

#### func (*Sock) RecvMessage

```go
func (s *Sock) RecvMessage() ([][]byte, error)
```
RecvMessage receives a full message from the socket and returns it as an array
of byte arrays.

#### func (*Sock) RecvString

```go
func (s *Sock) RecvString() (string, error)
```
RecvString reads a frame from the socket and returns it as a string, Returns an
error if the call fails.

#### func (*Sock) SendBytes

```go
func (s *Sock) SendBytes(data []byte, flags Flag) error
```
SendBytes sends a byte array via the socket. For the flags value, use 0 for a
single message, or SNDMORE if it is a multi-part message

#### func (*Sock) SendMessage

```go
func (s *Sock) SendMessage(parts ...interface{}) error
```
SendMessage is a variadic function that currently accepts ints, strings, and
bytes, and sends them as an atomic multi frame message over zeromq as a series
of byte arrays. In the case of numeric data, the resulting byte array is a
textual representation of the number (e.g., 100 turns to "100"). This may be
changed to network byte ordered representation in the near future - I have not
decided yet!

#### func (*Sock) SendString

```go
func (s *Sock) SendString(data string, flags Flag) error
```
SendString sends a string via the socket. For the flags value, use 0 for a
single message, or SNDMORE if it is a multi-part message

#### func (*Sock) SetAffinity

```go
func (z *Sock) SetAffinity(val int)
```
SetAffinity sets the affinity option for the socket

#### func (*Sock) SetBacklog

```go
func (z *Sock) SetBacklog(val int)
```
SetBacklog sets the backlog option for the socket

#### func (*Sock) SetConflate

```go
func (z *Sock) SetConflate(val int)
```
SetConflate sets the conflate option for the socket

#### func (*Sock) SetCurvePublickey

```go
func (z *Sock) SetCurvePublickey(val string)
```
SetCurvePublickey sets the curve_publickey option for the socket

#### func (*Sock) SetCurveSecretkey

```go
func (z *Sock) SetCurveSecretkey(val string)
```
SetCurveSecretkey sets the curve_secretkey option for the socket

#### func (*Sock) SetCurveServer

```go
func (z *Sock) SetCurveServer(val int)
```
SetCurveServer sets the curve_server option for the socket

#### func (*Sock) SetCurveServerkey

```go
func (z *Sock) SetCurveServerkey(val string)
```
SetCurveServerkey sets the curve_serverkey option for the socket

#### func (*Sock) SetDelayAttachOnConnect

```go
func (z *Sock) SetDelayAttachOnConnect(val int)
```
SetDelayAttachOnConnect sets the delay_attach_on_connect option for the socket

#### func (*Sock) SetGssapiPlaintext

```go
func (z *Sock) SetGssapiPlaintext(val int)
```
SetGssapiPlaintext sets the gssapi_plaintext option for the socket

#### func (*Sock) SetGssapiPrincipal

```go
func (z *Sock) SetGssapiPrincipal(val string)
```
SetGssapiPrincipal sets the gssapi_principal option for the socket

#### func (*Sock) SetGssapiServer

```go
func (z *Sock) SetGssapiServer(val int)
```
SetGssapiServer sets the gssapi_server option for the socket

#### func (*Sock) SetGssapiServicePrincipal

```go
func (z *Sock) SetGssapiServicePrincipal(val string)
```
SetGssapiServicePrincipal sets the gssapi_service_principal option for the
socket

#### func (*Sock) SetIdentity

```go
func (z *Sock) SetIdentity(val string)
```
SetIdentity sets the identity option for the socket

#### func (*Sock) SetImmediate

```go
func (z *Sock) SetImmediate(val int)
```
SetImmediate sets the immediate option for the socket

#### func (*Sock) SetIpv4only

```go
func (z *Sock) SetIpv4only(val int)
```
SetIpv4only sets the ipv4only option for the socket

#### func (*Sock) SetIpv6

```go
func (z *Sock) SetIpv6(val int)
```
SetIpv6 sets the ipv6 option for the socket

#### func (*Sock) SetLinger

```go
func (z *Sock) SetLinger(val int)
```
SetLinger sets the linger option for the socket

#### func (*Sock) SetMaxmsgsize

```go
func (z *Sock) SetMaxmsgsize(val int)
```
SetMaxmsgsize sets the maxmsgsize option for the socket

#### func (*Sock) SetMulticastHops

```go
func (z *Sock) SetMulticastHops(val int)
```
SetMulticastHops sets the multicast_hops option for the socket

#### func (*Sock) SetPlainPassword

```go
func (z *Sock) SetPlainPassword(val string)
```
SetPlainPassword sets the plain_password option for the socket

#### func (*Sock) SetPlainServer

```go
func (z *Sock) SetPlainServer(val int)
```
SetPlainServer sets the plain_server option for the socket

#### func (*Sock) SetPlainUsername

```go
func (z *Sock) SetPlainUsername(val string)
```
SetPlainUsername sets the plain_username option for the socket

#### func (*Sock) SetProbeRouter

```go
func (z *Sock) SetProbeRouter(val int)
```
SetProbeRouter sets the probe_router option for the socket

#### func (*Sock) SetRate

```go
func (z *Sock) SetRate(val int)
```
SetRate sets the rate option for the socket

#### func (*Sock) SetRcvbuf

```go
func (z *Sock) SetRcvbuf(val int)
```
SetRcvbuf sets the rcvbuf option for the socket

#### func (*Sock) SetRcvhwm

```go
func (z *Sock) SetRcvhwm(val int)
```
SetRcvhwm sets the rcvhwm option for the socket

#### func (*Sock) SetRcvtimeo

```go
func (z *Sock) SetRcvtimeo(val int)
```
SetRcvtimeo sets the rcvtimeo option for the socket

#### func (*Sock) SetReconnectIvl

```go
func (z *Sock) SetReconnectIvl(val int)
```
SetReconnectIvl sets the reconnect_ivl option for the socket

#### func (*Sock) SetReconnectIvlMax

```go
func (z *Sock) SetReconnectIvlMax(val int)
```
SetReconnectIvlMax sets the reconnect_ivl_max option for the socket

#### func (*Sock) SetRecoveryIvl

```go
func (z *Sock) SetRecoveryIvl(val int)
```
SetRecoveryIvl sets the recovery_ivl option for the socket

#### func (*Sock) SetReqCorrelate

```go
func (z *Sock) SetReqCorrelate(val int)
```
SetReqCorrelate sets the req_correlate option for the socket

#### func (*Sock) SetReqRelaxed

```go
func (z *Sock) SetReqRelaxed(val int)
```
SetReqRelaxed sets the req_relaxed option for the socket

#### func (*Sock) SetRouterHandover

```go
func (z *Sock) SetRouterHandover(val int)
```
SetRouterHandover sets the router_handover option for the socket

#### func (*Sock) SetRouterMandatory

```go
func (z *Sock) SetRouterMandatory(val int)
```
SetRouterMandatory sets the router_mandatory option for the socket

#### func (*Sock) SetRouterRaw

```go
func (z *Sock) SetRouterRaw(val int)
```
SetRouterRaw sets the router_raw option for the socket

#### func (*Sock) SetSndbuf

```go
func (z *Sock) SetSndbuf(val int)
```
SetSndbuf sets the sndbuf option for the socket

#### func (*Sock) SetSndhwm

```go
func (z *Sock) SetSndhwm(val int)
```
SetSndhwm sets the sndhwm option for the socket

#### func (*Sock) SetSndtimeo

```go
func (z *Sock) SetSndtimeo(val int)
```
SetSndtimeo sets the sndtimeo option for the socket

#### func (*Sock) SetSubscribe

```go
func (z *Sock) SetSubscribe(val string)
```
SetSubscribe sets the subscribe option for the socket

#### func (*Sock) SetTcpAcceptFilter

```go
func (z *Sock) SetTcpAcceptFilter(val string)
```
SetTcpAcceptFilter sets the tcp_accept_filter option for the socket

#### func (*Sock) SetTcpKeepalive

```go
func (z *Sock) SetTcpKeepalive(val int)
```
SetTcpKeepalive sets the tcp_keepalive option for the socket

#### func (*Sock) SetTcpKeepaliveCnt

```go
func (z *Sock) SetTcpKeepaliveCnt(val int)
```
SetTcpKeepaliveCnt sets the tcp_keepalive_cnt option for the socket

#### func (*Sock) SetTcpKeepaliveIdle

```go
func (z *Sock) SetTcpKeepaliveIdle(val int)
```
SetTcpKeepaliveIdle sets the tcp_keepalive_idle option for the socket

#### func (*Sock) SetTcpKeepaliveIntvl

```go
func (z *Sock) SetTcpKeepaliveIntvl(val int)
```
SetTcpKeepaliveIntvl sets the tcp_keepalive_intvl option for the socket

#### func (*Sock) SetTos

```go
func (z *Sock) SetTos(val int)
```
SetTos sets the tos option for the socket

#### func (*Sock) SetUnsubscribe

```go
func (z *Sock) SetUnsubscribe(val string)
```
SetUnsubscribe sets the unsubscribe option for the socket

#### func (*Sock) SetXpubVerbose

```go
func (z *Sock) SetXpubVerbose(val int)
```
SetXpubVerbose sets the xpub_verbose option for the socket

#### func (*Sock) SetZapDomain

```go
func (z *Sock) SetZapDomain(val string)
```
SetZapDomain sets the zap_domain option for the socket

#### func (*Sock) Sndbuf

```go
func (z *Sock) Sndbuf() int
```
Sndbuf returns the current value of the socket's sndbuf option

#### func (*Sock) Sndhwm

```go
func (z *Sock) Sndhwm() int
```
Sndhwm returns the current value of the socket's sndhwm option

#### func (*Sock) Sndtimeo

```go
func (z *Sock) Sndtimeo() int
```
Sndtimeo returns the current value of the socket's sndtimeo option

#### func (*Sock) TcpAcceptFilter

```go
func (z *Sock) TcpAcceptFilter() string
```
TcpAcceptFilter returns the current value of the socket's tcp_accept_filter
option

#### func (*Sock) TcpKeepalive

```go
func (z *Sock) TcpKeepalive() int
```
TcpKeepalive returns the current value of the socket's tcp_keepalive option

#### func (*Sock) TcpKeepaliveCnt

```go
func (z *Sock) TcpKeepaliveCnt() int
```
TcpKeepaliveCnt returns the current value of the socket's tcp_keepalive_cnt
option

#### func (*Sock) TcpKeepaliveIdle

```go
func (z *Sock) TcpKeepaliveIdle() int
```
TcpKeepaliveIdle returns the current value of the socket's tcp_keepalive_idle
option

#### func (*Sock) TcpKeepaliveIntvl

```go
func (z *Sock) TcpKeepaliveIntvl() int
```
TcpKeepaliveIntvl returns the current value of the socket's tcp_keepalive_intvl
option

#### func (*Sock) Tos

```go
func (z *Sock) Tos() int
```
Tos returns the current value of the socket's tos option

#### func (*Sock) Type

```go
func (z *Sock) Type() int
```
Type returns the current value of the socket's type option

#### func (*Sock) Unbind

```go
func (s *Sock) Unbind(endpoint string) error
```
Unbind unbinds a socket from an endpoint. If returns an error if the endpoint
was not found

#### func (*Sock) ZapDomain

```go
func (z *Sock) ZapDomain() string
```
ZapDomain returns the current value of the socket's zap_domain option

#### type Type

```go
type Type int
```
