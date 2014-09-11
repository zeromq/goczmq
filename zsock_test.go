package goczmq

import (
	"testing"
)

func TestZsock(t *testing.T) {

	pushSock := NewZsock(PUSH)
	pullSock := NewZsock(PULL)

	port, err := pullSock.Bind("inproc://test")
	if port != 0 {
		t.Errorf("port for Bind should be 0, is %d", port)
	}

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
