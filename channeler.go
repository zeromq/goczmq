package goczmq

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync/atomic"
)

// Channeler serializes all access to a socket through a send, receive
// and close channel.  It starts two threads, on is used for receiving
// from the zeromq socket.  The other is used to listen to the receive
// channel, and send everything back to the socket thrad for sending
// using an additional inproc socket.  The channeler takes ownership
// of the passed socket and will destroy it when the close channel
// is closed.
type Channeler struct {
	sock   *Sock
	id     int64
	closed int32

	closeChan  chan<- struct{}
	SendChan   chan<- [][]byte
	RecvChan   <-chan [][]byte
	AttachChan chan<- string
	ErrChan    <-chan error
}

// NewChanneler initialized a new channeler for the passed socket
// If sendErrors is true, errors will be sent on the error channel
// If it is false, any error will cause a panic
func NewChanneler(sock *Sock, sendErrors bool) *Channeler {
	closeChan := make(chan struct{})
	sendChan := make(chan [][]byte)
	recvChan := make(chan [][]byte)
	attachChan := make(chan string)

	var errChan chan error
	if sendErrors {
		errChan = make(chan error)
	}

	c := &Channeler{
		sock:       sock,
		id:         rand.Int63(),
		closeChan:  closeChan,
		SendChan:   sendChan,
		RecvChan:   recvChan,
		AttachChan: attachChan,
		ErrChan:    errChan,
	}

	go c.loopSend(closeChan, sendChan, attachChan)
	go c.loopMain(sendChan, recvChan, attachChan, errChan)

	runtime.SetFinalizer(c, func(c *Channeler) { c.Close() })
	return c
}

// Close closes the close channel sigaling the channeler to shut down
func (c *Channeler) Close() {
	if atomic.CompareAndSwapInt32(&c.closed, 0, 1) {
		close(c.closeChan)
	}
}

func (c *Channeler) loopSend(closeChan <-chan struct{}, sendChan <-chan [][]byte, attachChan <-chan string) {
	push, err := NewPush(fmt.Sprintf(">inproc://goczmq-channeler-%d", c.id))
	if err != nil {
		panic(err)
	}
	defer push.Destroy()

	for {
		select {
		case <-closeChan:
			_ = push.SendFrame([]byte("close"), 0)
			return
		case msg := <-sendChan:
			push.SendFrame([]byte("msg"), 1)
			var f int
			numFrames := len(msg)
			for i, val := range msg {
				if i == numFrames-1 {
					f = 0
				} else {
					f = More
				}

				_ = push.SendFrame(val, f)
			}
		case endpoint := <-attachChan:
			_ = push.SendFrame([]byte("attach"), 1)
			_ = push.SendFrame([]byte(endpoint), 0)
		}
	}
}

func (c *Channeler) loopMain(sendChan chan<- [][]byte, recvChan chan<- [][]byte, attachChan chan<- string, errChan chan<- error) {
	// Close all channels when we exit
	defer close(recvChan)
	defer close(sendChan)
	defer close(attachChan)
	if errChan != nil {
		defer close(errChan)
	}

	pull, err := NewPull(fmt.Sprintf("@inproc://goczmq-channeler-%d", c.id))
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
					if errChan != nil {
						errChan <- err
					}
				}
			case "attach":
				var err error
				switch int(c.sock.Type()) {
				case Pub, Rep, Router, Push, XPub:
					err = c.sock.Attach(string(msg[1]), true)
				default:
					err = c.sock.Attach(string(msg[1]), false)
				}

				if errChan != nil {
					errChan <- err
				}
			}

		case c.sock:
			recvChan <- msg
		}
	}
}
