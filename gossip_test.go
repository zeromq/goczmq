package goczmq

import (
	"testing"
)

func TestGossip(t *testing.T) {
	server := NewGossip("server")

	err := server.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	err = server.Bind("inproc://zgossip")
	if err != nil {
		t.Errorf("BIND error: %s", err)
	}

	server.Destroy()
}
