package goczmq

import (
	"testing"
	"time"
)

func TestChanneler(t *testing.T) {
	d2 := NewSock(Pair)
	_, err := d2.Bind("inproc://channeler-test")
	if err != nil {
		t.Errorf("Error creating d2: %s", err)
		return
	}

	d1 := NewSock(Pair)
	if err != nil {
		t.Errorf("Error creating d1: %s", err)
		return
	}

	c := NewChanneler(d1, false)
	c.AttachChan <- "inproc://channeler-test"
	c.SendChan <- [][]byte{[]byte("ready")}

	m, err := d2.RecvMessage()
	if string(m[0]) != "ready" {
		t.Errorf("Expected 'ready' but got %s", m)
		return
	}

	// The channeler listens on d1, do a send on d2 and verify the receive
	// channel of the channeler gets it
	err = d2.SendFrame([]byte("Test"), 0)
	if err != nil {
		t.Errorf("d2.SendMessage failed: %s", err)
		return
	}

	select {
	case msg := <-c.RecvChan:
		if string(msg[0]) != "Test" {
			t.Error("Message received on receive channel mismatch")
			return
		}
	case <-time.After(time.Millisecond * 250):
		t.Error("Timeout while waiting for receive channel")
		return
	}

	// Send a message through the channeler and verify d2 gets it
	c.SendChan <- [][]byte{[]byte("Test")}
	poller, err := NewPoller(d2)
	if err != nil {
		t.Errorf("Error while creating poller: %s", err)
		return
	}

	s := poller.Wait(250)
	if s == nil {
		t.Error("Timeout while waiting for send channel")
		return
	}

	msg, err := d2.RecvMessage()
	if err != nil {
		t.Errorf("Error while receiving message on d2: %s", err)
		return
	}

	if string(msg[0]) != "Test" {
		t.Error("Message received on d2 mismatch")
		return
	}

	c.Close()
}

func ExampleChanneler() {
	router := NewSock(Router)
	routerChan := NewChanneler(router, false)
	routerChan.AttachChan <- "inproc://channel_example"

	routerChan.Close()
}

func benchmarkChanneler(size int, b *testing.B) {
	pullSock := NewSock(Pull)
	pullSock.Bind("inproc://benchChan")
	defer pullSock.Destroy()

	channeler := NewChanneler(pullSock, false)
	time.Sleep(10 * time.Millisecond)

	go func() {
		pushSock := NewSock(Push)
		defer pushSock.Destroy()
		err := pushSock.Connect("inproc://benchChan")
		if err != nil {
			panic(err)
		}

		payload := make([]byte, size)
		for i := 0; i < b.N; i++ {
			_, err = pushSock.Write(payload)
			if err != nil {
				panic(err)
			}
		}
	}()

	for i := 0; i < b.N; i++ {
		msg := <-channeler.RecvChan
		if len(msg[0]) != size {
			panic("msg too small")
		}
	}
	channeler.Close()
}

func BenchmarkChanneler1k(b *testing.B) { benchmarkChanneler(1024, b) }
