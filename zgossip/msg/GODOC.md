# msg
--
    import "github.com/zeromq/goczmq/zgossip/msg"

Package msg is 100% generated. If you edit this file, you will lose your changes
at the next build cycle. DO NOT MAKE ANY CHANGES YOU WISH TO KEEP.

The correct places for commits are:

    - The XML model used for this code generation: zgossip_msg.xml
    - The code generation script that built this file: zproto_codec_goczmq

## Usage

```go
const (
	HelloId   uint8 = 1
	PublishId uint8 = 2
	PingId    uint8 = 3
	PongId    uint8 = 4
	InvalidId uint8 = 5
)
```

```go
const (
	Signature uint16 = 0xAAA0 | 0
)
```

#### type Hello

```go
type Hello struct {
}
```

Client says hello to server

#### func  NewHello

```go
func NewHello() *Hello
```
New creates new Hello message.

#### func (*Hello) Marshal

```go
func (h *Hello) Marshal() ([]byte, error)
```
Marshal serializes the message.

#### func (*Hello) RoutingId

```go
func (h *Hello) RoutingId() []byte
```
RoutingId returns the routingId for this message, routingId should be set
whenever talking to a ROUTER.

#### func (*Hello) Send

```go
func (h *Hello) Send(sock *goczmq.Sock) (err error)
```
Send sends marshaled data through 0mq socket.

#### func (*Hello) SetRoutingId

```go
func (h *Hello) SetRoutingId(routingId []byte)
```
SetRoutingId sets the routingId for this message, routingId should be set
whenever talking to a ROUTER.

#### func (*Hello) SetVersion

```go
func (h *Hello) SetVersion(version byte)
```
Setversion sets the version.

#### func (*Hello) String

```go
func (h *Hello) String() string
```
String returns print friendly name.

#### func (*Hello) Unmarshal

```go
func (h *Hello) Unmarshal(frames ...[]byte) error
```
Unmarshal unmarshals the message.

#### func (*Hello) Version

```go
func (h *Hello) Version() byte
```
version returns the version.

#### type Invalid

```go
type Invalid struct {
}
```

Server rejects command as invalid

#### func  NewInvalid

```go
func NewInvalid() *Invalid
```
New creates new Invalid message.

#### func (*Invalid) Marshal

```go
func (i *Invalid) Marshal() ([]byte, error)
```
Marshal serializes the message.

#### func (*Invalid) RoutingId

```go
func (i *Invalid) RoutingId() []byte
```
RoutingId returns the routingId for this message, routingId should be set
whenever talking to a ROUTER.

#### func (*Invalid) Send

```go
func (i *Invalid) Send(sock *goczmq.Sock) (err error)
```
Send sends marshaled data through 0mq socket.

#### func (*Invalid) SetRoutingId

```go
func (i *Invalid) SetRoutingId(routingId []byte)
```
SetRoutingId sets the routingId for this message, routingId should be set
whenever talking to a ROUTER.

#### func (*Invalid) SetVersion

```go
func (i *Invalid) SetVersion(version byte)
```
Setversion sets the version.

#### func (*Invalid) String

```go
func (i *Invalid) String() string
```
String returns print friendly name.

#### func (*Invalid) Unmarshal

```go
func (i *Invalid) Unmarshal(frames ...[]byte) error
```
Unmarshal unmarshals the message.

#### func (*Invalid) Version

```go
func (i *Invalid) Version() byte
```
version returns the version.

#### type Ping

```go
type Ping struct {
}
```

Client signals liveness

#### func  NewPing

```go
func NewPing() *Ping
```
New creates new Ping message.

#### func (*Ping) Marshal

```go
func (p *Ping) Marshal() ([]byte, error)
```
Marshal serializes the message.

#### func (*Ping) RoutingId

```go
func (p *Ping) RoutingId() []byte
```
RoutingId returns the routingId for this message, routingId should be set
whenever talking to a ROUTER.

#### func (*Ping) Send

```go
func (p *Ping) Send(sock *goczmq.Sock) (err error)
```
Send sends marshaled data through 0mq socket.

#### func (*Ping) SetRoutingId

```go
func (p *Ping) SetRoutingId(routingId []byte)
```
SetRoutingId sets the routingId for this message, routingId should be set
whenever talking to a ROUTER.

#### func (*Ping) SetVersion

```go
func (p *Ping) SetVersion(version byte)
```
Setversion sets the version.

#### func (*Ping) String

```go
func (p *Ping) String() string
```
String returns print friendly name.

#### func (*Ping) Unmarshal

```go
func (p *Ping) Unmarshal(frames ...[]byte) error
```
Unmarshal unmarshals the message.

#### func (*Ping) Version

```go
func (p *Ping) Version() byte
```
version returns the version.

#### type Pong

```go
type Pong struct {
}
```

Server responds to ping; note that pongs are not correlated with pings,

#### func  NewPong

```go
func NewPong() *Pong
```
New creates new Pong message.

#### func (*Pong) Marshal

```go
func (p *Pong) Marshal() ([]byte, error)
```
Marshal serializes the message.

#### func (*Pong) RoutingId

```go
func (p *Pong) RoutingId() []byte
```
RoutingId returns the routingId for this message, routingId should be set
whenever talking to a ROUTER.

#### func (*Pong) Send

```go
func (p *Pong) Send(sock *goczmq.Sock) (err error)
```
Send sends marshaled data through 0mq socket.

#### func (*Pong) SetRoutingId

```go
func (p *Pong) SetRoutingId(routingId []byte)
```
SetRoutingId sets the routingId for this message, routingId should be set
whenever talking to a ROUTER.

#### func (*Pong) SetVersion

```go
func (p *Pong) SetVersion(version byte)
```
Setversion sets the version.

#### func (*Pong) String

```go
func (p *Pong) String() string
```
String returns print friendly name.

#### func (*Pong) Unmarshal

```go
func (p *Pong) Unmarshal(frames ...[]byte) error
```
Unmarshal unmarshals the message.

#### func (*Pong) Version

```go
func (p *Pong) Version() byte
```
version returns the version.

#### type Publish

```go
type Publish struct {
	Key   string
	Value string
	Ttl   uint32
}
```

Client or server announces a new tuple

#### func  NewPublish

```go
func NewPublish() *Publish
```
New creates new Publish message.

#### func (*Publish) Marshal

```go
func (p *Publish) Marshal() ([]byte, error)
```
Marshal serializes the message.

#### func (*Publish) RoutingId

```go
func (p *Publish) RoutingId() []byte
```
RoutingId returns the routingId for this message, routingId should be set
whenever talking to a ROUTER.

#### func (*Publish) Send

```go
func (p *Publish) Send(sock *goczmq.Sock) (err error)
```
Send sends marshaled data through 0mq socket.

#### func (*Publish) SetRoutingId

```go
func (p *Publish) SetRoutingId(routingId []byte)
```
SetRoutingId sets the routingId for this message, routingId should be set
whenever talking to a ROUTER.

#### func (*Publish) SetVersion

```go
func (p *Publish) SetVersion(version byte)
```
Setversion sets the version.

#### func (*Publish) String

```go
func (p *Publish) String() string
```
String returns print friendly name.

#### func (*Publish) Unmarshal

```go
func (p *Publish) Unmarshal(frames ...[]byte) error
```
Unmarshal unmarshals the message.

#### func (*Publish) Version

```go
func (p *Publish) Version() byte
```
version returns the version.

#### type Transit

```go
type Transit interface {
	Marshal() ([]byte, error)
	Unmarshal(...[]byte) error
	String() string
	Send(*goczmq.Sock) error
	SetRoutingId([]byte)
	RoutingId() []byte
	SetVersion(byte)
	Version() byte
}
```

Transit is a codec interface

#### func  Clone

```go
func Clone(t Transit) Transit
```
Clone clones a message.

#### func  Recv

```go
func Recv(sock *goczmq.Sock) (t Transit, err error)
```
Recv receives marshaled data from a 0mq socket.

#### func  RecvNoWait

```go
func RecvNoWait(sock *goczmq.Sock) (t Transit, err error)
```
RecvNoWait receives marshaled data from 0mq socket. It won't wait for input.

#### func  Unmarshal

```go
func Unmarshal(frames ...[]byte) (t Transit, err error)
```
Unmarshal unmarshals data from raw frames.
