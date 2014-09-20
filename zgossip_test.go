package goczmq

import (
	"testing"
)

func TestZgossip(t *testing.T) {
	server := NewZgossip("server")

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
