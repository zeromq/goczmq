# goczmq
--
    import "github.com/taotetek/goczmq"

Package goczmq is a go interface to CZMQ

go:generate gsl sockopts.xml

## Usage

```go
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
	MORE     = Flag(C.ZFRAME_MORE)
	REUSE    = Flag(C.ZFRAME_REUSE)
	DONTWAIT = Flag(C.ZFRAME_DONTWAIT)

	CURVE_ALLOW_ANY = "*"
)
```

```go
var (
	ErrActorCmd   = errors.New("error sending actor command")
	ErrSockAttach = errors.New("error attaching zsock")
)
```

#### type Auth

```go
type Auth struct {
}
```

Auth wraps a CZMQ zauth zactor

#### func  NewAuth

```go
func NewAuth() *Auth
```
NewAuth creates a new Auth actor.

#### func (*Auth) Allow

```go
func (a *Auth) Allow(address string) error
```
Allow removes a previous Deny

#### func (*Auth) Curve

```go
func (a *Auth) Curve(allowed string) error
```
Curve sets auth method to curve

#### func (*Auth) Deny

```go
func (a *Auth) Deny(address string) error
```
Deny adds an address to a socket's deny list

#### func (*Auth) Destroy

```go
func (a *Auth) Destroy()
```
Destroy destroys the auth actor.

#### func (*Auth) Plain

```go
func (a *Auth) Plain(directory string) error
```
Plain sets auth method to plain

#### func (*Auth) Verbose

```go
func (a *Auth) Verbose() error
```
Verbose sets the auth actor to log information to stdout.

#### type Beacon

```go
type Beacon struct {
}
```

Beacon wraps a CZMQ zbeacon zactor

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

#### type Cert

```go
type Cert struct {
}
```

Cert holds a czmq zcert_t

#### func  NewCert

```go
func NewCert() *Cert
```
NewCert creates a new empty Cert instance

#### func  NewCertFromFile

```go
func NewCertFromFile(filename string) (*Cert, error)
```
NewCertFromFile Load loads a Cert from files

#### func  NewCertFromKeys

```go
func NewCertFromKeys(public []byte, secret []byte) (*Cert, error)
```
NewCertFromKeys creates a new Cert from a public and private key

#### func (*Cert) Apply

```go
func (c *Cert) Apply(s *Sock)
```
Apply sets the public and private keys for a socket

#### func (*Cert) Destroy

```go
func (c *Cert) Destroy()
```
Destroy destroys Cert instance

#### func (*Cert) Dup

```go
func (c *Cert) Dup() *Cert
```
Dup duplicates a Cert

#### func (*Cert) Equal

```go
func (c *Cert) Equal(compare *Cert) bool
```
Equal checks two Certs for equality

#### func (*Cert) Meta

```go
func (c *Cert) Meta(key string) string
```
Meta returns a meta data item from a Cert given a key

#### func (*Cert) Print

```go
func (c *Cert) Print()
```
Print prints a Cert to stdout

#### func (*Cert) PublicText

```go
func (c *Cert) PublicText() string
```
PublicText returns the public key as a string

#### func (*Cert) Save

```go
func (c *Cert) Save(filename string) error
```
Save saves the public and secret key to filename and filename_secret

#### func (*Cert) SavePublic

```go
func (c *Cert) SavePublic(filename string) error
```
SavePublic saves the public key to a file

#### func (*Cert) SaveSecret

```go
func (c *Cert) SaveSecret(filename string) error
```
SaveSecret saves the secret key to a file

#### func (*Cert) SetMeta

```go
func (c *Cert) SetMeta(key string, value string)
```
SetMeta sets meta data for a Cert

#### type Channeler

```go
type Channeler struct {
	SendChan   chan<- [][]byte
	RecvChan   <-chan [][]byte
	AttachChan chan<- string
	ErrChan    <-chan error
}
```

Channeler serializes all access to a socket through a send, receive and close
channel. It starts two threads, on is used for receiving from the zeromq socket.
The other is used to listen to the receive channel, and send everything back to
the socket thrad for sending using an additional inproc socket. The channeler
takes ownership of the passed socket and will destroy it when the close channel
is closed.

#### func  NewChanneler

```go
func NewChanneler(sock *Sock, sendErrors bool) *Channeler
```
NewChanneler initialized a new channeler for the passed socket If sendErrors is
true, errors will be sent on the error channel If it is false, any error will
cause a panic

#### func (*Channeler) Close

```go
func (c *Channeler) Close()
```
Close closes the close channel sigaling the channeler to shut down

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
func (g *Gossip) Bind(endpoint string) error
```
Bind binds the gossip service to a specified endpoint

#### func (*Gossip) Destroy

```go
func (g *Gossip) Destroy()
```
Destroy destroys the gossip actor.

#### func (*Gossip) Verbose

```go
func (g *Gossip) Verbose() error
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
Remove removes a Sock from the poller

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

Proxy actors switch messages between a frontend and backend socket.

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
func (p *Proxy) SetBackend(sockType int, endpoint string) error
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
func (p *Proxy) SetFrontend(sockType int, endpoint string) error
```
SetFrontend accepts a socket type and endpoint, and sends a message to the
zactor thread telling it to set up a socket bound to the endpoint.

#### func (*Proxy) Verbose

```go
func (p *Proxy) Verbose() error
```
Verbose sets the proxy to log information to stdout.

#### type ReadChunker

```go
type ReadChunker struct {
}
```

ReadChunker accepts a socket and a chunkSize, and implements the ReaderFrom
interface.

#### func  NewReadChunker

```go
func NewReadChunker(s *Sock, cs int64) *ReadChunker
```
NewReadChunker takes a socket and a chunkSize and returns a new ReadChunker
instance

#### func (*ReadChunker) Destroy

```go
func (c *ReadChunker) Destroy()
```
Destroy calls destroy on the underlying socket to clean it up

#### func (*ReadChunker) ReadFrom

```go
func (c *ReadChunker) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom reads from an io.Reader into a []byte of chunkSize. It writes each
chunk of data as a two frame message. The first frame is used to indicate if
this is the last message or if there are more messages containing file data to
come.

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
NewDEALER creates a DEALER socket and calls Attach. The socket will Connect by
default.

#### func  NewPAIR

```go
func NewPAIR(endpoints string) (*Sock, error)
```
NewPAIR creates a PAIR socket and calls Attach. The socket will Connect by
default.

#### func  NewPUB

```go
func NewPUB(endpoints string) (*Sock, error)
```
NewPUB creates a PUB socket and calls Attach. The socket will Bind by default.

#### func  NewPULL

```go
func NewPULL(endpoints string) (*Sock, error)
```
NewPULL creates a PULL socket and calls Attach. The socket will Bind by default.

#### func  NewPUSH

```go
func NewPUSH(endpoints string) (*Sock, error)
```
NewPUSH creates a PUSH socket and calls Attach. The socket will Connect by
default.

#### func  NewREP

```go
func NewREP(endpoints string) (*Sock, error)
```
NewREP creates a REP socket and calls Attach. The socket will Bind by default.

#### func  NewREQ

```go
func NewREQ(endpoints string) (*Sock, error)
```
NewREQ creates a REQ socket and calls Attach. The socket will Connect by
default.

#### func  NewROUTER

```go
func NewROUTER(endpoints string) (*Sock, error)
```
NewROUTER creates a ROUTER socket and calls Attach. The socket will Bind by
default.

#### func  NewSTREAM

```go
func NewSTREAM(endpoints string) (*Sock, error)
```
NewSTREAM creates a STREAM socket and calls Attach. The socket will Connect by
default.

#### func  NewSUB

```go
func NewSUB(endpoints string, subscribe string) (*Sock, error)
```
NewSUB creates a SUB socket and calls Attach. 'subscribe' is a comma delimited
list of topics to subscribe to. The socket will Connect by default.

#### func  NewSock

```go
func NewSock(t int) *Sock
```
NewSock creates a new socket. The caller source and line number are passed so
CZMQ can report socket leaks intelligently.

#### func  NewXPUB

```go
func NewXPUB(endpoints string) (*Sock, error)
```
NewXPUB creates an XPUB socket and calls Attach. The socket will Bind by
default.

#### func  NewXSUB

```go
func NewXSUB(endpoints string) (*Sock, error)
```
NewXSUB creates an XSUB socket and calls Attach. The socket will Connect by
default.

#### func (*Sock) Affinity

```go
func (s *Sock) Affinity() int
```
Affinity returns the current value of the socket's affinity option

#### func (*Sock) Attach

```go
func (s *Sock) Attach(endpoints string, serverish bool) error
```
Attach attaches a socket to zero or more endpoints. If endpoints is not null,
parses as list of ZeroMQ endpoints, separated by commas, and prefixed by '@' (to
bind the socket) or '>' (to attach the socket). If the endpoint does not start
with '@' or '>', the serverish argument determines whether it is used to bind
(serverish = true) or connect (serverish = false)

#### func (*Sock) Backlog

```go
func (s *Sock) Backlog() int
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
func (s *Sock) CurvePublickey() string
```
CurvePublickey returns the current value of the socket's curve_publickey option

#### func (*Sock) CurveSecretkey

```go
func (s *Sock) CurveSecretkey() string
```
CurveSecretkey returns the current value of the socket's curve_secretkey option

#### func (*Sock) CurveServer

```go
func (s *Sock) CurveServer() int
```
CurveServer returns the current value of the socket's curve_server option

#### func (*Sock) CurveServerkey

```go
func (s *Sock) CurveServerkey() string
```
CurveServerkey returns the current value of the socket's curve_serverkey option

#### func (*Sock) Destroy

```go
func (s *Sock) Destroy()
```
Destroy destroys the underlying zsockT.

#### func (*Sock) Disconnect

```go
func (s *Sock) Disconnect(endpoint string) error
```
Disconnect disconnects a socket from an endpoint. If returns an error if the
endpoint was not found

#### func (*Sock) Events

```go
func (s *Sock) Events() int
```
Events returns the current value of the socket's events option

#### func (*Sock) Fd

```go
func (s *Sock) Fd() int
```
Fd returns the current value of the socket's fd option

#### func (*Sock) GetType

```go
func (s *Sock) GetType() int
```
GetType returns the socket's type

#### func (*Sock) GssapiPlaintext

```go
func (s *Sock) GssapiPlaintext() int
```
GssapiPlaintext returns the current value of the socket's gssapi_plaintext
option

#### func (*Sock) GssapiPrincipal

```go
func (s *Sock) GssapiPrincipal() string
```
GssapiPrincipal returns the current value of the socket's gssapi_principal
option

#### func (*Sock) GssapiServer

```go
func (s *Sock) GssapiServer() int
```
GssapiServer returns the current value of the socket's gssapi_server option

#### func (*Sock) GssapiServicePrincipal

```go
func (s *Sock) GssapiServicePrincipal() string
```
GssapiServicePrincipal returns the current value of the socket's
gssapi_service_principal option

#### func (*Sock) Identity

```go
func (s *Sock) Identity() string
```
Identity returns the current value of the socket's identity option

#### func (*Sock) Immediate

```go
func (s *Sock) Immediate() int
```
Immediate returns the current value of the socket's immediate option

#### func (*Sock) Ipv4only

```go
func (s *Sock) Ipv4only() int
```
Ipv4only returns the current value of the socket's ipv4only option

#### func (*Sock) Ipv6

```go
func (s *Sock) Ipv6() int
```
Ipv6 returns the current value of the socket's ipv6 option

#### func (*Sock) LastEndpoint

```go
func (s *Sock) LastEndpoint() string
```
LastEndpoint returns the current value of the socket's last_endpoint option

#### func (*Sock) Linger

```go
func (s *Sock) Linger() int
```
Linger returns the current value of the socket's linger option

#### func (*Sock) Maxmsgsize

```go
func (s *Sock) Maxmsgsize() int
```
Maxmsgsize returns the current value of the socket's maxmsgsize option

#### func (*Sock) Mechanism

```go
func (s *Sock) Mechanism() int
```
Mechanism returns the current value of the socket's mechanism option

#### func (*Sock) MulticastHops

```go
func (s *Sock) MulticastHops() int
```
MulticastHops returns the current value of the socket's multicast_hops option

#### func (*Sock) PlainPassword

```go
func (s *Sock) PlainPassword() string
```
PlainPassword returns the current value of the socket's plain_password option

#### func (*Sock) PlainServer

```go
func (s *Sock) PlainServer() int
```
PlainServer returns the current value of the socket's plain_server option

#### func (*Sock) PlainUsername

```go
func (s *Sock) PlainUsername() string
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
func (s *Sock) Rate() int
```
Rate returns the current value of the socket's rate option

#### func (*Sock) Rcvbuf

```go
func (s *Sock) Rcvbuf() int
```
Rcvbuf returns the current value of the socket's rcvbuf option

#### func (*Sock) Rcvhwm

```go
func (s *Sock) Rcvhwm() int
```
Rcvhwm returns the current value of the socket's rcvhwm option

#### func (*Sock) Rcvmore

```go
func (s *Sock) Rcvmore() int
```
Rcvmore returns the current value of the socket's rcvmore option

#### func (*Sock) Rcvtimeo

```go
func (s *Sock) Rcvtimeo() int
```
Rcvtimeo returns the current value of the socket's rcvtimeo option

#### func (*Sock) ReconnectIvl

```go
func (s *Sock) ReconnectIvl() int
```
ReconnectIvl returns the current value of the socket's reconnect_ivl option

#### func (*Sock) ReconnectIvlMax

```go
func (s *Sock) ReconnectIvlMax() int
```
ReconnectIvlMax returns the current value of the socket's reconnect_ivl_max
option

#### func (*Sock) RecoveryIvl

```go
func (s *Sock) RecoveryIvl() int
```
RecoveryIvl returns the current value of the socket's recovery_ivl option

#### func (*Sock) RecvFrame

```go
func (s *Sock) RecvFrame() ([]byte, Flag, error)
```
RecvFrame reads a frame from the socket and returns it as a byte array, along
with a more flag and and error (if there is an error)

#### func (*Sock) RecvFrameNoWait

```go
func (s *Sock) RecvFrameNoWait() ([]byte, Flag, error)
```
RecvFrameNoWait receives a frame from the socket and returns it as a byte array
if one is waiting. Returns an empty frame, a 0 more flag and an error if one is
not immediately available

#### func (*Sock) RecvMessage

```go
func (s *Sock) RecvMessage() ([][]byte, error)
```
RecvMessage receives a full message from the socket and returns it as an array
of byte arrays.

#### func (*Sock) RecvMessageNoWait

```go
func (s *Sock) RecvMessageNoWait() ([][]byte, error)
```
RecvMessageNoWait receives a full message from the socket and returns it as an
array of byte arrays if one is waiting. Returns an empty message and an error if
one is not immediately available

#### func (*Sock) SendFrame

```go
func (s *Sock) SendFrame(data []byte, flags Flag) error
```
SendFrame sends a byte array via the socket. For the flags value, use 0 for a
single message, or SNDMORE if it is a multi-part message

#### func (*Sock) SendMessage

```go
func (s *Sock) SendMessage(parts [][]byte) error
```
SendMessage accepts an array of byte arrays and sends it as a multi-part
message.

#### func (*Sock) SetAffinity

```go
func (s *Sock) SetAffinity(val int)
```
SetAffinity sets the affinity option for the socket

#### func (*Sock) SetBacklog

```go
func (s *Sock) SetBacklog(val int)
```
SetBacklog sets the backlog option for the socket

#### func (*Sock) SetConflate

```go
func (s *Sock) SetConflate(val int)
```
SetConflate sets the conflate option for the socket

#### func (*Sock) SetCurvePublickey

```go
func (s *Sock) SetCurvePublickey(val string)
```
SetCurvePublickey sets the curve_publickey option for the socket

#### func (*Sock) SetCurveSecretkey

```go
func (s *Sock) SetCurveSecretkey(val string)
```
SetCurveSecretkey sets the curve_secretkey option for the socket

#### func (*Sock) SetCurveServer

```go
func (s *Sock) SetCurveServer(val int)
```
SetCurveServer sets the curve_server option for the socket

#### func (*Sock) SetCurveServerkey

```go
func (s *Sock) SetCurveServerkey(val string)
```
SetCurveServerkey sets the curve_serverkey option for the socket

#### func (*Sock) SetDelayAttachOnConnect

```go
func (s *Sock) SetDelayAttachOnConnect(val int)
```
SetDelayAttachOnConnect sets the delay_attach_on_connect option for the socket

#### func (*Sock) SetGssapiPlaintext

```go
func (s *Sock) SetGssapiPlaintext(val int)
```
SetGssapiPlaintext sets the gssapi_plaintext option for the socket

#### func (*Sock) SetGssapiPrincipal

```go
func (s *Sock) SetGssapiPrincipal(val string)
```
SetGssapiPrincipal sets the gssapi_principal option for the socket

#### func (*Sock) SetGssapiServer

```go
func (s *Sock) SetGssapiServer(val int)
```
SetGssapiServer sets the gssapi_server option for the socket

#### func (*Sock) SetGssapiServicePrincipal

```go
func (s *Sock) SetGssapiServicePrincipal(val string)
```
SetGssapiServicePrincipal sets the gssapi_service_principal option for the
socket

#### func (*Sock) SetIdentity

```go
func (s *Sock) SetIdentity(val string)
```
SetIdentity sets the identity option for the socket

#### func (*Sock) SetImmediate

```go
func (s *Sock) SetImmediate(val int)
```
SetImmediate sets the immediate option for the socket

#### func (*Sock) SetIpv4only

```go
func (s *Sock) SetIpv4only(val int)
```
SetIpv4only sets the ipv4only option for the socket

#### func (*Sock) SetIpv6

```go
func (s *Sock) SetIpv6(val int)
```
SetIpv6 sets the ipv6 option for the socket

#### func (*Sock) SetLinger

```go
func (s *Sock) SetLinger(val int)
```
SetLinger sets the linger option for the socket

#### func (*Sock) SetMaxmsgsize

```go
func (s *Sock) SetMaxmsgsize(val int)
```
SetMaxmsgsize sets the maxmsgsize option for the socket

#### func (*Sock) SetMulticastHops

```go
func (s *Sock) SetMulticastHops(val int)
```
SetMulticastHops sets the multicast_hops option for the socket

#### func (*Sock) SetPlainPassword

```go
func (s *Sock) SetPlainPassword(val string)
```
SetPlainPassword sets the plain_password option for the socket

#### func (*Sock) SetPlainServer

```go
func (s *Sock) SetPlainServer(val int)
```
SetPlainServer sets the plain_server option for the socket

#### func (*Sock) SetPlainUsername

```go
func (s *Sock) SetPlainUsername(val string)
```
SetPlainUsername sets the plain_username option for the socket

#### func (*Sock) SetProbeRouter

```go
func (s *Sock) SetProbeRouter(val int)
```
SetProbeRouter sets the probe_router option for the socket

#### func (*Sock) SetRate

```go
func (s *Sock) SetRate(val int)
```
SetRate sets the rate option for the socket

#### func (*Sock) SetRcvbuf

```go
func (s *Sock) SetRcvbuf(val int)
```
SetRcvbuf sets the rcvbuf option for the socket

#### func (*Sock) SetRcvhwm

```go
func (s *Sock) SetRcvhwm(val int)
```
SetRcvhwm sets the rcvhwm option for the socket

#### func (*Sock) SetRcvtimeo

```go
func (s *Sock) SetRcvtimeo(val int)
```
SetRcvtimeo sets the rcvtimeo option for the socket

#### func (*Sock) SetReconnectIvl

```go
func (s *Sock) SetReconnectIvl(val int)
```
SetReconnectIvl sets the reconnect_ivl option for the socket

#### func (*Sock) SetReconnectIvlMax

```go
func (s *Sock) SetReconnectIvlMax(val int)
```
SetReconnectIvlMax sets the reconnect_ivl_max option for the socket

#### func (*Sock) SetRecoveryIvl

```go
func (s *Sock) SetRecoveryIvl(val int)
```
SetRecoveryIvl sets the recovery_ivl option for the socket

#### func (*Sock) SetReqCorrelate

```go
func (s *Sock) SetReqCorrelate(val int)
```
SetReqCorrelate sets the req_correlate option for the socket

#### func (*Sock) SetReqRelaxed

```go
func (s *Sock) SetReqRelaxed(val int)
```
SetReqRelaxed sets the req_relaxed option for the socket

#### func (*Sock) SetRouterHandover

```go
func (s *Sock) SetRouterHandover(val int)
```
SetRouterHandover sets the router_handover option for the socket

#### func (*Sock) SetRouterMandatory

```go
func (s *Sock) SetRouterMandatory(val int)
```
SetRouterMandatory sets the router_mandatory option for the socket

#### func (*Sock) SetRouterRaw

```go
func (s *Sock) SetRouterRaw(val int)
```
SetRouterRaw sets the router_raw option for the socket

#### func (*Sock) SetSndbuf

```go
func (s *Sock) SetSndbuf(val int)
```
SetSndbuf sets the sndbuf option for the socket

#### func (*Sock) SetSndhwm

```go
func (s *Sock) SetSndhwm(val int)
```
SetSndhwm sets the sndhwm option for the socket

#### func (*Sock) SetSndtimeo

```go
func (s *Sock) SetSndtimeo(val int)
```
SetSndtimeo sets the sndtimeo option for the socket

#### func (*Sock) SetSubscribe

```go
func (s *Sock) SetSubscribe(val string)
```
SetSubscribe sets the subscribe option for the socket

#### func (*Sock) SetTcpAcceptFilter

```go
func (s *Sock) SetTcpAcceptFilter(val string)
```
SetTcpAcceptFilter sets the tcp_accept_filter option for the socket

#### func (*Sock) SetTcpKeepalive

```go
func (s *Sock) SetTcpKeepalive(val int)
```
SetTcpKeepalive sets the tcp_keepalive option for the socket

#### func (*Sock) SetTcpKeepaliveCnt

```go
func (s *Sock) SetTcpKeepaliveCnt(val int)
```
SetTcpKeepaliveCnt sets the tcp_keepalive_cnt option for the socket

#### func (*Sock) SetTcpKeepaliveIdle

```go
func (s *Sock) SetTcpKeepaliveIdle(val int)
```
SetTcpKeepaliveIdle sets the tcp_keepalive_idle option for the socket

#### func (*Sock) SetTcpKeepaliveIntvl

```go
func (s *Sock) SetTcpKeepaliveIntvl(val int)
```
SetTcpKeepaliveIntvl sets the tcp_keepalive_intvl option for the socket

#### func (*Sock) SetTos

```go
func (s *Sock) SetTos(val int)
```
SetTos sets the tos option for the socket

#### func (*Sock) SetUnsubscribe

```go
func (s *Sock) SetUnsubscribe(val string)
```
SetUnsubscribe sets the unsubscribe option for the socket

#### func (*Sock) SetXpubVerbose

```go
func (s *Sock) SetXpubVerbose(val int)
```
SetXpubVerbose sets the xpub_verbose option for the socket

#### func (*Sock) SetZapDomain

```go
func (s *Sock) SetZapDomain(val string)
```
SetZapDomain sets the zap_domain option for the socket

#### func (*Sock) Sndbuf

```go
func (s *Sock) Sndbuf() int
```
Sndbuf returns the current value of the socket's sndbuf option

#### func (*Sock) Sndhwm

```go
func (s *Sock) Sndhwm() int
```
Sndhwm returns the current value of the socket's sndhwm option

#### func (*Sock) Sndtimeo

```go
func (s *Sock) Sndtimeo() int
```
Sndtimeo returns the current value of the socket's sndtimeo option

#### func (*Sock) TcpAcceptFilter

```go
func (s *Sock) TcpAcceptFilter() string
```
TcpAcceptFilter returns the current value of the socket's tcp_accept_filter
option

#### func (*Sock) TcpKeepalive

```go
func (s *Sock) TcpKeepalive() int
```
TcpKeepalive returns the current value of the socket's tcp_keepalive option

#### func (*Sock) TcpKeepaliveCnt

```go
func (s *Sock) TcpKeepaliveCnt() int
```
TcpKeepaliveCnt returns the current value of the socket's tcp_keepalive_cnt
option

#### func (*Sock) TcpKeepaliveIdle

```go
func (s *Sock) TcpKeepaliveIdle() int
```
TcpKeepaliveIdle returns the current value of the socket's tcp_keepalive_idle
option

#### func (*Sock) TcpKeepaliveIntvl

```go
func (s *Sock) TcpKeepaliveIntvl() int
```
TcpKeepaliveIntvl returns the current value of the socket's tcp_keepalive_intvl
option

#### func (*Sock) Tos

```go
func (s *Sock) Tos() int
```
Tos returns the current value of the socket's tos option

#### func (*Sock) Type

```go
func (s *Sock) Type() int
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
func (s *Sock) ZapDomain() string
```
ZapDomain returns the current value of the socket's zap_domain option

#### type WriteChunker

```go
type WriteChunker struct {
}
```

WriteChunker accepts a socket and implements the WriterTo interface

#### func  NewWriteChunker

```go
func NewWriteChunker(s *Sock) *WriteChunker
```
NewWriteChunker takes a socket and returns a new WriteChunker instance

#### func (*WriteChunker) Destroy

```go
func (c *WriteChunker) Destroy()
```
Destroy calls destroy on the underlying socket to clean it up

#### func (*WriteChunker) WriteTo

```go
func (c *WriteChunker) WriteTo(w io.Writer) (int64, error)
```
WriteTo to reads each chunk message one at a time and writes them to an
io.Writer.
