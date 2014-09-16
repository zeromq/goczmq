package goczmq

import (
	"testing"
)

func TestZpoller(t *testing.T) {
	pullSock1, err := NewPULL("inproc://poller_pull1")
	if err != nil {
		t.Errorf("NewPULL failed: %s", err)
	}

	pullSock2, err := NewPULL("inproc://poller_pull2")
	if err != nil {
		t.Errorf("NewPULL failed: %s", err)
	}

	poller, err := NewZpoller(pullSock1, pullSock2)
	if err != nil {
		t.Errorf("NewZpoller failed: %s", err)
	}

	pullSock3, err := NewPULL("inproc://poller_pull3")
	if err != nil {
		t.Errorf("NewPULL failed: %s", err)
	}

	err = poller.Add(pullSock3)
	if err != nil {
		t.Errorf("poller.Add failed: %s", err)
	}

	err = poller.Remove(pullSock3)
	if err != nil {
		t.Errorf("poller.Remove failed: %s", err)
	}

	poller.Destroy()
}
