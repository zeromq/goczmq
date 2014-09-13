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
		t.Errorf("repSock.Bind failed: %s", err)
	}

	err = pushSock.Connect("inproc://test")
	if err != nil {
		t.Errorf("reqSock.Connect failed: %s", err)
	}

	err = pushSock.SendBytes([]byte("Hello"), 0)
	if err != nil {
		t.Errorf("pushSock.SendBytes failed: %s", err)
	}

	bmsg, err := pullSock.RecvBytes()
	if err != nil {
		t.Errorf("pullSock.RecvBytes failed: %s", err)
	}

	if string(bmsg) != "Hello" {
		t.Errorf("expected 'Hello' received '%s'", string(bmsg))
	}

	err = pushSock.SendString("World", 0)
	if err != nil {
		t.Errorf("pushSock.SendString failed: %s", err)
	}

	smsg, err := pullSock.RecvString()
	if err != nil {
		t.Errorf("pullSock.RecvString failed: %s", err)
	}

	if smsg != "World" {
		t.Errorf("expected 'World' received '%s'", smsg)
	}

	pushSock.Destroy()
	pullSock.Destroy()
}
