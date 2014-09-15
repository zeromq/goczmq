package goczmq

import (
	"testing"
)

func TestZpoller(t *testing.T) {
	pullSock, err := NewPULL("inproc://proxy_pull")
	if err != nil {
		t.Errorf("NewPULL failed: %s", err)
	}

	poller, err := NewZpoller(pullSock)
	if err != nil {
		t.Errorf("NewZpoller failed: %s", err)
	}

	poller.Destroy()
}
