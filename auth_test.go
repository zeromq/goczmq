package goczmq

import (
	"fmt"
	"testing"
)

func TestAuthNoDeny(t *testing.T) {
	server := NewSock(PULL)
	defer server.Destroy()

	server.SetZapDomain("global")
	server.Bind("tcp://*:9999")

	client := NewSock(PUSH)
	defer client.Destroy()

	auth := NewAuth()
	defer auth.Destroy()

	err := auth.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	port, err := server.Bind("tcp://127.0.0.1:*")
	if port <= 0 {
		t.Error("port should be > 0, is %d", port)
	}

	err = client.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Errorf("client connect error: %s", err)
	}

	client.SendString("Hello, World", 0)

	poller, err := NewPoller(server)
	if err != nil {
		t.Errorf("NwPoller failed: %s", err)
	}

	defer poller.Destroy()

	s := poller.Wait(200)
	if s == nil {
		t.Error("should be message waiting and there is none")
	}

	msg, err := s.RecvString()
	if err != nil {
		t.Error(err)
	}

	if msg != "Hello, World" {
		t.Error("message not sent properly")
	}
}

func TestAuthDeny(t *testing.T) {
	server := NewSock(PULL)
	defer server.Destroy()

	server.SetZapDomain("global")
	server.Bind("tcp://*:9999")

	client := NewSock(PUSH)
	defer client.Destroy()

	auth := NewAuth()
	defer auth.Destroy()

	err := auth.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	auth.Deny("127.0.0.1")

	port, err := server.Bind("tcp://127.0.0.1:*")
	if port <= 0 {
		t.Error("port should be > 0, is %d", port)
	}

	err = client.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Errorf("client connect error: %s", err)
	}

	client.SendString("Hello, World", 0)

	poller, err := NewPoller(server)
	if err != nil {
		t.Errorf("NwPoller failed: %s", err)
	}

	defer poller.Destroy()

	s := poller.Wait(200)
	if s != nil {
		t.Error("poller received message should not have!")
	}
}
