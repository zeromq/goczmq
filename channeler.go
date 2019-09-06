package goczmq

/*
#include "czmq.h"
void Sock_init() {zsys_init();}
*/
import "C"

import (
	"fmt"
	"strings"
)

// Channeler serializes all access to a socket through a send
// and receive channel.  It starts two threads, one is used for receiving
// from the zeromq socket.  The other is used to listen to the receive
// channel, and send everything back to the socket thread for sending
// using an additional inproc socket.
type Channeler struct {
	id          string
	sockType    int
	endpoints   string
	subscribe   *string
	commandAddr string
	proxyAddr   string
	commandChan chan<- string
	SendChan    chan<- [][]byte
	RecvChan    <-chan [][]byte
	ErrChan     <-chan error
	errChan     chan<- error
	destroyed   bool
}

// Destroy sends a message to the Channeler to shut it down
// and clean it up.
func (c *Channeler) Destroy() {
	if c.destroyed {
		return
	}
	c.commandChan <- "destroy"
}

// Subscribe to a Topic
func (c *Channeler) Subscribe(topic string) {
	if c.destroyed {
		return
	}
	c.commandChan <- fmt.Sprintf("subscribe %s", topic)
}

// Unsubscribe from a Topic
func (c *Channeler) Unsubscribe(topic string) {
	if c.destroyed {
		return
	}
	c.commandChan <- fmt.Sprintf("unsubscribe %s", topic)
}

// actor is a routine that handles communication with
// the zeromq socket.
func (c *Channeler) actor(recvChan chan<- [][]byte, options []SockOption) {
	pipe, err := NewPair(fmt.Sprintf(">%s", c.commandAddr))

	if err != nil {
		c.errChan <- err
		return
	}
	defer pipe.Destroy()
	defer close(recvChan)

	pull, err := NewPull(c.proxyAddr)
	if err != nil {
		c.errChan <- err
		return
	}
	defer pull.Destroy()

	sock := NewSock(c.sockType, options...)
	defer sock.Destroy()
	switch c.sockType {
	case Pub, Rep, Pull, Router, XPub:
		err = sock.Attach(c.endpoints, true)
		if err != nil {
			c.errChan <- err
			return
		}

	case Req, Push, Dealer, Pair, Stream, XSub:
		err = sock.Attach(c.endpoints, false)
		if err != nil {
			c.errChan <- err
			return
		}

	case Sub:
		if c.subscribe != nil {
			subscriptions := strings.Split(*c.subscribe, ",")
			for _, topic := range subscriptions {
				sock.SetOption(SockSetSubscribe(topic))
			}
		}

		err = sock.Attach(c.endpoints, false)
		if err != nil {
			c.errChan <- err
			return
		}

	default:
		c.errChan <- ErrInvalidSockType
		return
	}

	poller, err := NewPoller(sock, pull, pipe)
	if err != nil {
		c.errChan <- err
		goto ExitActor
	}
	defer poller.Destroy()

	for {
		s, err := poller.Wait(-1)
		if err != nil {
			c.errChan <- err
			goto ExitActor
		}
		switch s {
		case pipe:
			cmd, err := pipe.RecvMessage()
			if err != nil {
				c.errChan <- err
				goto ExitActor
			}

			switch string(cmd[0]) {
			case "destroy":
				disconnect := strings.Split(c.endpoints, ",")
				for _, endpoint := range disconnect {
					sock.Disconnect(endpoint)
				}
				pipe.SendMessage([][]byte{[]byte("ok")})
				goto ExitActor
			case "subscribe":
				topic := string(cmd[1])
				sock.SetOption(SockSetSubscribe(topic))
				pipe.SendMessage([][]byte{[]byte("ok")})
			case "unsubscribe":
				topic := string(cmd[1])
				sock.SetOption(SockSetUnsubscribe(topic))
				pipe.SendMessage([][]byte{[]byte("ok")})
			}

		case sock:
			msg, err := s.RecvMessage()
			if err != nil {
				c.errChan <- err
				goto ExitActor
			}
			recvChan <- msg

		case pull:
			msg, err := pull.RecvMessage()
			if err != nil {
				c.errChan <- err
				goto ExitActor
			}

			err = sock.SendMessage(msg)
			if err != nil {
				c.errChan <- err
				goto ExitActor
			}
		}
	}
ExitActor:
}

// channeler is a routine that handles the channel select loop
// and sends commands to the zeromq socket.
func (c *Channeler) channeler(commandChan <-chan string, sendChan <-chan [][]byte) {
	push, err := NewPush(c.proxyAddr)
	if err != nil {
		c.errChan <- err
		return
	}
	defer push.Destroy()

	pipe, err := NewPair(fmt.Sprintf("@%s", c.commandAddr))
	if err != nil {
		c.errChan <- err
		goto ExitChanneler
	}
	defer pipe.Destroy()

	for {
		select {
		case cmd := <-commandChan:
			switch cmd {
			case "destroy":
				c.destroyed = true
				err = pipe.SendFrame([]byte("destroy"), FlagNone)
				if err != nil {
					c.errChan <- err
				} else {
					_, err = pipe.RecvMessage()
					if err != nil {
						c.errChan <- err
					}
				}
				goto ExitChanneler
			default:
				parts := strings.Split(cmd, " ")
				numParts := len(parts)
				message := make([][]byte, numParts, numParts)
				for i, p := range parts {
					message[i] = []byte(p)
				}
				err := pipe.SendMessage(message)
				if err != nil {
					c.errChan <- err
				}
				_, err = pipe.RecvMessage()
				if err != nil {
					c.errChan <- err
				}
			}

		case msg := <-sendChan:
			err := push.SendMessage(msg)
			if err != nil {
				c.errChan <- err
			}
		}
	}
ExitChanneler:
}

// newChanneler accepts arguments from the socket type based
// constructors and creates a new Channeler instance
func newChanneler(sockType int, endpoints string, subscribe []string, options []SockOption) *Channeler {
	commandChan := make(chan string)
	sendChan := make(chan [][]byte)
	recvChan := make(chan [][]byte)
	errChan := make(chan error)

	C.Sock_init()
	c := &Channeler{
		id:          C.GoString(C.zuuid_str(C.zuuid_new())),
		endpoints:   endpoints,
		sockType:    sockType,
		commandChan: commandChan,
		SendChan:    sendChan,
		RecvChan:    recvChan,
		ErrChan:     errChan,
		errChan:     errChan,
	}
	c.commandAddr = fmt.Sprintf("inproc://actorcontrol_%s", c.id)
	c.proxyAddr = fmt.Sprintf("inproc://proxy_%s", c.id)

	if len(subscribe) > 0 {
		topics := strings.Join(subscribe, ",")
		c.subscribe = &topics
	}

	go c.channeler(commandChan, sendChan)
	go c.actor(recvChan, options)

	return c
}

// NewPubChanneler creats a new Channeler wrapping
// a Pub socket.  The socket will bind by default.
func NewPubChanneler(endpoints string, options ...SockOption) *Channeler {
	return newChanneler(Pub, endpoints, nil, options)
}

// NewSubChanneler creates a new Channeler wrapping
// a Sub socket. Along with an endpoint list
// it accepts a list of topics and/or socket options
// (discriminated by type). The socket will connect
// by default.
func NewSubChanneler(endpoints string, varargs ...interface{}) *Channeler {
	subscribe := []string{}
	options := []SockOption{}
	var err error

	for _, arg := range varargs {
		switch x := arg.(type) {
		case string:
			subscribe = append(subscribe, x)
		case SockOption:
			options = append(options, x)
		default:
			err = fmt.Errorf("Don't know how to handle a %T argument to NewSubChanneler", arg)
			break
		}
	}

	channeler := newChanneler(Sub, endpoints, subscribe, options)

	if err != nil {
		go func() { channeler.errChan <- err }()
	}
	return channeler
}

// NewRepChanneler creates a new Channeler wrapping
// a Rep socket. The socket will bind by default.
func NewRepChanneler(endpoints string, options ...SockOption) *Channeler {
	return newChanneler(Rep, endpoints, nil, options)
}

// NewReqChanneler creates a new Channeler wrapping
// a Req socket. The socket will connect by default.
func NewReqChanneler(endpoints string, options ...SockOption) *Channeler {
	return newChanneler(Req, endpoints, nil, options)
}

// NewPullChanneler creates a new Channeler wrapping
// a Pull socket. The socket will bind by default.
func NewPullChanneler(endpoints string, options ...SockOption) *Channeler {
	return newChanneler(Pull, endpoints, nil, options)
}

// NewPushChanneler creates a new Channeler wrapping
// a Push socket. The socket will connect by default.
func NewPushChanneler(endpoints string, options ...SockOption) *Channeler {
	return newChanneler(Push, endpoints, nil, options)
}

// NewRouterChanneler creates a new Channeler wrapping
// a Router socket. The socket will Bind by default.
func NewRouterChanneler(endpoints string, options ...SockOption) *Channeler {
	return newChanneler(Router, endpoints, nil, options)
}

// NewDealerChanneler creates a new Channeler wrapping
// a Dealer socket. The socket will connect by default.
func NewDealerChanneler(endpoints string, options ...SockOption) *Channeler {
	return newChanneler(Dealer, endpoints, nil, options)
}

// NewXPubChanneler creates a new Channeler wrapping
// an XPub socket. The socket will Bind by default.
func NewXPubChanneler(endpoints string, options ...SockOption) *Channeler {
	return newChanneler(XPub, endpoints, nil, options)
}

// NewXSubChanneler creates a new Channeler wrapping
// a XSub socket. The socket will connect by default.
func NewXSubChanneler(endpoints string, options ...SockOption) *Channeler {
	return newChanneler(XSub, endpoints, nil, options)
}

// NewPairChanneler creates a new Channeler wrapping
// a Pair socket. The socket will connect by default.
func NewPairChanneler(endpoints string, options ...SockOption) *Channeler {
	return newChanneler(Pair, endpoints, nil, options)
}

// NewStreamChanneler creates a new Channeler wrapping
// a Pair socket. The socket will connect by default.
func NewStreamChanneler(endpoints string, options ...SockOption) *Channeler {
	return newChanneler(Stream, endpoints, nil, options)
}
