package czmq

import (
	"testing"
)

func TestZsock(t *testing.T) {

	pushSock := NewZsock(PUSH)
	pullSock := NewZsock(PULL)

	err := pullSock.Bind("inproc://test")
	if err != nil {
		t.Error("repSock.Bind failed")
	}
	err = pushSock.Connect("inproc://test")
	if err != nil {
		t.Error("reqSock.Connect failed")
	}

	err = pushSock.SendBytes([]byte("Hello"), 0)
	if err != nil {
		t.Error("pushSock.SendBytes failed")
	}

	msg, err := pullSock.RecvBytes()
	if err != nil {
		t.Error("pullSock.RecvBytes: %s", err)
	}

	if string(msg) != "Hello" {
		t.Error("expected 'Hello' received '%s'", string(msg))
	}

	pushSock.Destroy()
	pullSock.Destroy()
}
