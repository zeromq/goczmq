package goczmq

import (
	"fmt"
	"math/rand"
	"strings"
)

type Channeler struct {
	id          int64
	subscribe   string
	commandAddr string
	proxyAddr   string
	commandChan chan<- string
	SendChan    chan<- [][]byte
	RecvChan    <-chan [][]byte
}

func (c *Channeler) Destroy() {
	c.commandChan <- "destroy"
}

func (c *Channeler) actor(recvChan chan<- [][]byte, t int, endpoints string) {
	pipe, err := NewPair(fmt.Sprintf(">%s", c.commandAddr))
	if err != nil {
		panic(err)
	}
	defer pipe.Destroy()

	pull, err := NewPull(c.proxyAddr)
	if err != nil {
		panic(err)
	}
	defer pull.Destroy()

	sock := NewSock(t)
	defer sock.Destroy()
	switch t {
	case Pub, Rep, Pull, Router, XPub:
		err = sock.Attach(endpoints, true)
		if err != nil {
			panic(err)
		}

	case Req, Push, Dealer, Pair, Stream:
		err = sock.Attach(endpoints, false)
		if err != nil {
			panic(err)
		}
	case Sub, XSub:
		subscriptions := strings.Split(c.subscribe, ",")
		for _, topic := range subscriptions {
			sock.SetSubscribe(topic)
		}

		err = sock.Attach(endpoints, false)
		if err != nil {
			panic(err)
		}

	default:
		panic(ErrInvalidSockType)
	}

	poller, err := NewPoller(sock, pull, pipe)
	if err != nil {
		panic(err)
	}
	defer poller.Destroy()

	for {
		s := poller.Wait(-1)
		switch s {
		case pipe:
			cmd, err := pipe.RecvMessage()
			if err != nil {
				panic(err)
			}

			switch string(cmd[0]) {
			case "destroy":
				disconnect := strings.Split(endpoints, ",")
				for _, endpoint := range disconnect {
					sock.Disconnect(endpoint)
				}
				pipe.SendMessage([][]byte{[]byte("ok")})
				goto ExitActor
			}
		case sock:
			msg, err := s.RecvMessage()
			if err != nil {
				panic(err)
			}
			recvChan <- msg
		case pull:
			msg, err := pull.RecvMessage()
			if err != nil {
				panic(err)
			}

			err = sock.SendMessage(msg)
			if err != nil {
				panic(err)
			}
		}
	}
ExitActor:
}

func (c *Channeler) channeler(commandChan <-chan string, sendChan <-chan [][]byte) {
	push, err := NewPush(c.proxyAddr)
	if err != nil {
		panic(err)
	}
	defer push.Destroy()

	pipe, err := NewPair(fmt.Sprintf("@%s", c.commandAddr))
	if err != nil {
		panic(err)
	}
	defer pipe.Destroy()

	for {
		select {
		case cmd := <-commandChan:
			switch cmd {
			case "destroy":
				err = pipe.SendFrame([]byte("destroy"), FlagNone)
				if err != nil {
					panic(err)
				}
				_, err = pipe.RecvMessage()
				if err != nil {
					panic(err)
				}
				goto ExitChanneler
			}
		case msg := <-sendChan:
			err := push.SendMessage(msg)
			if err != nil {
				panic(err)
			}
		}
	}
ExitChanneler:
}

func newChanneler(t int, endpoints, subscribe string) *Channeler {
	commandChan := make(chan string)
	sendChan := make(chan [][]byte)
	recvChan := make(chan [][]byte)

	c := &Channeler{
		id:          rand.Int63(),
		subscribe:   subscribe,
		commandChan: commandChan,
		SendChan:    sendChan,
		RecvChan:    recvChan,
	}

	c.commandAddr = fmt.Sprintf("inproc://actorcontrol%d", c.id)
	c.proxyAddr = fmt.Sprintf("inproc://proxy%d", c.id)

	go c.channeler(commandChan, sendChan)
	go c.actor(recvChan, t, endpoints)

	return c
}

func NewPubChanneler(endpoints string) *Channeler {
	return newChanneler(Pub, endpoints, "")
}

func NewSubChanneler(endpoints, subscribe string) *Channeler {
	return newChanneler(Sub, endpoints, subscribe)
}

func NewRepChanneler(endpoints string) *Channeler {
	return newChanneler(Rep, endpoints, "")
}

func NewReqChanneler(endpoints string) *Channeler {
	return newChanneler(Req, endpoints, "")
}

func NewPullChanneler(endpoints string) *Channeler {
	return newChanneler(Pull, endpoints, "")
}

func NewPushChanneler(endpoints string) *Channeler {
	return newChanneler(Push, endpoints, "")
}

func NewRouterChanneler(endpoints string) *Channeler {
	return newChanneler(Router, endpoints, "")
}

func NewDealerChanneler(endpoints string) *Channeler {
	return newChanneler(Dealer, endpoints, "")
}

func NewXPubChanneler(endpoints string) *Channeler {
	return newChanneler(XPub, endpoints, "")
}

func NewXSubChanneler(endpoints, subscribe string) *Channeler {
	return newChanneler(XSub, endpoints, subscribe)
}

func NewPairChanneler(endpoints string) *Channeler {
	return newChanneler(Pair, endpoints, "")
}

func NewStreamChanneler(endpoints string) *Channeler {
	return newChanneler(Stream, endpoints, "")
}
