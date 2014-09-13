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

	bmsg, f, err := pullSock.RecvBytes()
	if err != nil {
		t.Errorf("pullSock.RecvBytes failed: %s", err)
	}

	if f == MORE {
		t.Errorf("MORE set and should not be")
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

func TestMessage(t *testing.T) {
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

	err = pushSock.SendMessage("The", "Answer", []byte("Is"), 42)
	if err != nil {
		t.Errorf("pushSock.SendMessages failed: %s", err)
	}

	multiMsg, err := pullSock.RecvMessage()
	if err != nil {
		t.Errorf("pullSock.RecvMessage failed: %s", err)
	}

	if len(multiMsg) != 4 {
		t.Errorf("pullSock.recvMessage expected 4 message frames, got %d", len(multiMsg))
	}

	if string(multiMsg[0]) != "The" {
		t.Errorf("expected 'The', received '%s'", string(multiMsg[0]))
	}

	if string(multiMsg[1]) != "Answer" {
		t.Errorf("expected 'Answer', received '%s'", string(multiMsg[1]))
	}

	if string(multiMsg[2]) != "Is" {
		t.Errorf("expected 'Is', received '%s'", string(multiMsg[2]))
	}

	if string(multiMsg[3]) != "42" {
		t.Errorf("expected '42', received '%s'", string(multiMsg[3]))
	}
}
