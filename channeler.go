package goczmq

import (
	"fmt"
	"math/rand"
	"runtime"
)

// Channeler serializes all access to a socket through a send, receive and close channel
// It starts two threads, one is used for receiving from the zeromq socket
// The other is used to listen to the receive channel, and send everything back to the socket thread for sending using an additional inproc socket
// The channeler takes ownership of the passed socket and will destroy it when the close channel is closed
type Channeler struct {
	sock *Sock
	id   int64

	close chan<- struct{}

	Send    chan<- [][]byte
	Receive <-chan [][]byte
	Connect chan<- string
	Error   <-chan error
}

// NewChanneler initialized a new channeler for the passed socket
// If sendErrors is true, errors will be sent on the error channel
// If it is false, any error will cause a panic
func NewChanneler(sock *Sock, sendErrors bool) *Channeler {
	close := make(chan struct{})
	send := make(chan [][]byte)
	receive := make(chan [][]byte)
	connect := make(chan string)
	var err chan error
	if sendErrors {
		err = make(chan error)
	}

	c := &Channeler{
		sock:    sock,
		id:      rand.Int63(),
		close:   close,
		Send:    send,
		Receive: receive,
		Connect: connect,
		Error:   err,
	}

	go c.loopSend(close, send, connect)
	go c.loopMain(send, receive, connect, err)

	runtime.SetFinalizer(c, func(c *Channeler) { c.Close() })
	return c
}

func (c *Channeler) Close() {
	close(c.close)
}

func (c *Channeler) loopSend(closeChan <-chan struct{}, send <-chan [][]byte, connect <-chan string) {
	push, err := NewPUSH(fmt.Sprintf(">inproc://goczmq-channeler-%d", c.id))
	if err != nil {
		panic(err)
	}
	defer push.Destroy()

	for {
		select {
		case <-closeChan:
			_ = push.SendMessage("close")
			return
		case msg := <-send:
			push.SendMessage("msg", msg)
		case endpoint := <-connect:
			push.SendMessage("connect", endpoint)
		}
	}
}

func (c *Channeler) loopMain(send chan<- [][]byte, receive chan<- [][]byte, connect chan<- string, error chan<- error) {
	// Close all channels when we exit
	defer close(receive)
	defer close(send)
	defer close(connect)

	pull, err := NewPULL(fmt.Sprintf("@inproc://goczmq-channeler-%d", c.id))
	if err != nil {
		panic(err)
	}
	defer pull.Destroy()

	poller, err := NewPoller(pull, c.sock)
	if err != nil {
		panic(err)
	}
	defer poller.Destroy()

	for {
		s := poller.Wait(-1)
		if s == nil {
			continue
		}

		msg, err := s.RecvMessage()
		if err != nil {
			panic(err)
		}

		switch s {
		case pull:
			switch string(msg[0]) {
			case "close":
				return
			case "msg":
				if err := c.sock.SendMessage(msg[1:]); err != nil {
					if error != nil {
						error <- err
					} else {
						panic(err)
					}
				}
			case "connect":
				if err := c.sock.Connect(string(msg[1])); err != nil {
					if error != nil {
						error <- err
					} else {
						panic(err)
					}
				}
			}

		case c.sock:
			receive <- msg
		}
	}
}
