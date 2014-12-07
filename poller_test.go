package goczmq

import (
	"testing"
)

func TestPoller(t *testing.T) {
	pullSock1, err := NewPULL("inproc://poller_pull1")
	if err != nil {
		t.Errorf("NewPULL failed: %s", err)
	}
	defer pullSock1.Destroy()

	poller, err := NewPoller(pullSock1)
	if err != nil {
		t.Errorf("NewPoller failed: %s", err)
	}
	defer poller.Destroy()

	if len(poller.socks) != 1 {
		t.Errorf("Expected number of socks to be 1, was %d", len(poller.socks))
	}

	pullSock2, err := NewPULL("inproc://poller_pull2")
	if err != nil {
		t.Errorf("NewPULL failed: %s", err)
	}
	defer pullSock2.Destroy()

	err = poller.Add(pullSock2)
	if err != nil {
		t.Errorf("poller Add failed: %s", err)
	}

	if len(poller.socks) != 2 {
		t.Errorf("Expected number of socks to be 2, was %d", len(poller.socks))
	}

	poller.Destroy()
	poller, err = NewPoller(pullSock1, pullSock2)
	if err != nil {
		t.Errorf("NewPoller failed: %s", err)
	}

	if len(poller.socks) != 2 {
		t.Errorf("Expected number of zsocks to be 2, was %d", len(poller.socks))
	}

	if poller.socks[0].zsockT != pullSock1.zsockT || poller.socks[1].zsockT != pullSock2.zsockT {
		t.Error("Expected each passed zsock to be in the poller")
	}

	pushSock, err := NewPUSH("inproc://poller_pull1")
	if err != nil {
		t.Errorf("NewPUSH failed: %s", err)
	}
	defer pushSock.Destroy()

	err = pushSock.SendString("Hello", 0)
	if err != nil {
		t.Errorf("SendMessage failed: %s", err)
	}

	s := poller.Wait(0)
	if s == nil {
		t.Errorf("Wait did not return waiting socket")
	}

	msg, err := s.RecvString()
	if err != nil {
		t.Errorf("RecvMessage failed: %s", err)
	}

	if msg != "Hello" {
		t.Errorf("Expected 'Hello', received %s", msg)
	}

	pushSock2, err := NewPUSH("inproc://poller_pull2")
	if err != nil {
		t.Errorf("NewPUSH failed: %s", err)
	}

	err = pushSock2.SendString("World", 0)
	if err != nil {
		t.Errorf("SendMessage failed: %s", err)
	}

	s = poller.Wait(0)
	if s == nil {
		t.Errorf("Wait did not return waiting socket")
	}

	msg, err = s.RecvString()
	if err != nil {
		t.Errorf("RecvMessage failed: %s", err)
	}

	if msg != "World" {
		t.Errorf("Expected 'World', received %s", msg)
	}

	poller.Remove(pullSock2)
	if len(poller.socks) != 1 {
		t.Errorf("socks len should be 1 after removing pushsock, is %d", len(poller.socks))
	}
}
