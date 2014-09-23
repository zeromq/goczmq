package goczmq

import (
	"testing"
)

func TestZpoller(t *testing.T) {
	pullSock1, err := NewPULL("inproc://poller_pull1")
	if err != nil {
		t.Errorf("NewPULL failed: %s", err)
	}

	poller, err := NewZpoller(pullSock1)
	if err != nil {
		t.Errorf("NewZpoller failed: %s", err)
	}

	if len(poller.zsocks) != 1 {
		t.Errorf("Expected number of zsocks to be 1, was %d", len(poller.zsocks))
	}

	pullSock2, err := NewPULL("inproc://poller_pull2")
	if err != nil {
		t.Errorf("NewPULL failed: %s", err)
	}

	err = poller.Add(pullSock2)
	if err != nil {
		t.Errorf("poller Add failed: %s", err)
	}

	if len(poller.zsocks) != 2 {
		t.Errorf("Expected number of zsocks to be 2, was %d", len(poller.zsocks))
	}

	pushSock, err := NewPUSH("inproc://poller_pull1")
	if err != nil {
		t.Errorf("NewPUSH failed: %s", err)
	}

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
	if len(poller.zsocks) != 1 {
		t.Errorf("zsocks len should be 1 after removing pushsock, is %d", len(poller.zsocks))
	}
	poller.Destroy()
}
