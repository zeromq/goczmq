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

	err = pushSock.SendMessage([]byte("The"), "Answer", []byte("Is"), 42, []byte(""))
	if err != nil {
		t.Errorf("pushSock.SendMessages failed: %s", err)
	}

	multiMsg, err := pullSock.RecvMessage()
	if err != nil {
		t.Errorf("pullSock.RecvMessage failed: %s", err)
	}

	if len(multiMsg) != 5 {
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

	if string(multiMsg[4]) != string([]byte{0}) {
		t.Errorf("expected null byte, received '%s'", string(multiMsg[4]))
	}
}

func TestPubSub(t *testing.T) {
	_, err := NewZsockPub("bogus://bogus")
	if err == nil {
		t.Error("NewZsockPub should have returned error and did not")
	}

	_, err = NewZsockSub("bogus://bogus", "")
	if err == nil {
		t.Error("NewZsockPub should have returned error and did not")
	}

	pub, err := NewZsockPub("inproc://pub1,inproc://pub2")
	if err != nil {
		t.Errorf("NewZsockPub failed: %s", err)
	}

	sub, err := NewZsockSub("inproc://pub1,inproc://pub2", "")
	if err != nil {
		t.Errorf("NewZsockSub failed: %s", err)
	}

	err = pub.SendMessage("test pub sub")
	if err != nil {
		t.Errorf("SendMessage failed: %s", err)
	}

	msg, err := sub.RecvMessage()
	if err != nil {
		t.Errorf("RecvMessage failed: %s", err)
	}

	if string(msg[0]) != "test pub sub" {
		t.Errorf("Expected 'test pub sub', received %s", msg)
	}
}

func TestReqRep(t *testing.T) {
	_, err := NewZsockReq("bogus://bogus")
	if err == nil {
		t.Error("NewZsockPub should have returned error and did not")
	}

	_, err = NewZsockRep("bogus://bogus")
	if err == nil {
		t.Error("NewZsockPub should have returned error and did not")
	}

	rep, err := NewZsockRep("inproc://rep1,inproc://rep2")
	if err != nil {
		t.Errorf("NewZsockRep failed: %s", err)
	}

	req, err := NewZsockReq("inproc://rep1,inproc://rep2")
	if err != nil {
		t.Errorf("NewZsockReq failed: %s", err)
	}

	err = req.SendMessage("Hello")
	if err != nil {
		t.Errorf("SendMessage failed: %s", err)
	}

	reqmsg, err := rep.RecvMessage()
	if err != nil {
		t.Errorf("RecvMessage failed: %s", err)
	}

	if string(reqmsg[0]) != "Hello" {
		t.Errorf("Expected 'Hello', received '%s", string(reqmsg[0]))
	}

	err = rep.SendMessage("World")
	if err != nil {
		t.Errorf("SendMessage failed: %s", err)
	}

	repmsg, err := req.RecvMessage()
	if err != nil {
		t.Errorf("RecvMessage failed: %s", err)
	}

	if string(repmsg[0]) != "World" {
		t.Errorf("Expected 'World', received '%s", string(repmsg[0]))
	}

}
