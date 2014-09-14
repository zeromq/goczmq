package goczmq

import (
	"testing"
)

func TestZsock(t *testing.T) {

	pushSock := NewZsock(PUSH)
	defer pushSock.Destroy()

	pullSock := NewZsock(PULL)
	defer pullSock.Destroy()

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
	defer pushSock.Destroy()

	pullSock := NewZsock(PULL)
	defer pullSock.Destroy()

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

func TestPUBSUB(t *testing.T) {
	_, err := NewPUB("bogus://bogus")
	if err == nil {
		t.Error("NewPUB should have returned error and did not")
	}

	_, err = NewSUB("bogus://bogus", "")
	if err == nil {
		t.Error("NewSUB should have returned error and did not")
	}

	pub, err := NewPUB("inproc://pub1,inproc://pub2")
	if err != nil {
		t.Errorf("NewPUB failed: %s", err)
	}
	defer pub.Destroy()

	sub, err := NewSUB("inproc://pub1,inproc://pub2", "")
	if err != nil {
		t.Errorf("NewSUB failed: %s", err)
	}
	defer sub.Destroy()

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

func TestREQREP(t *testing.T) {
	_, err := NewREQ("bogus://bogus")
	if err == nil {
		t.Error("NewREQ should have returned error and did not")
	}

	_, err = NewREP("bogus://bogus")
	if err == nil {
		t.Error("NewREP should have returned error and did not")
	}

	rep, err := NewREP("inproc://rep1,inproc://rep2")
	if err != nil {
		t.Errorf("NewREP failed: %s", err)
	}
	defer rep.Destroy()

	req, err := NewREQ("inproc://rep1,inproc://rep2")
	if err != nil {
		t.Errorf("NewREQ failed: %s", err)
	}
	defer req.Destroy()

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

func TestPUSHPULL(t *testing.T) {
	_, err := NewPUSH("bogus://bogus")
	if err == nil {
		t.Error("NewPUSH should have returned error and did not")
	}

	_, err = NewPULL("bogus://bogus")
	if err == nil {
		t.Error("NewPULL should have returned error and did not")
	}

	push, err := NewPUSH("inproc://push1,inproc://push2")
	if err != nil {
		t.Errorf("NewPUSH failed: %s", err)
	}
	defer push.Destroy()

	pull, err := NewPULL("inproc://push1,inproc://push2")
	if err != nil {
		t.Errorf("NewPULL failed: %s", err)
	}
	defer pull.Destroy()

	err = push.SendMessage("Hello", "World")
	if err != nil {
		t.Errorf("SendMessage failed: %s", err)
	}

	msg, err := pull.RecvMessage()
	if err != nil {
		t.Errorf("RecvMessage failed: %s", err)
	}

	if string(msg[0]) != "Hello" {
		t.Errorf("Expected 'Hello', received '%s", string(msg[0]))
	}

	if string(msg[1]) != "World" {
		t.Errorf("Expected 'World', received '%s", string(msg[0]))
	}
}

func TestROUTERDEALER(t *testing.T) {
	_, err := NewDEALER("bogus://bogus")
	if err == nil {
		t.Error("NewDEALER should have returned error and did not")
	}

	_, err = NewROUTER("bogus://bogus")
	if err == nil {
		t.Error("NewROUTER should have returned error and did not")
	}

	dealer, err := NewDEALER("inproc://router1,inproc://router2")
	if err != nil {
		t.Errorf("NewDEALER failed: %s", err)
	}
	defer dealer.Destroy()

	router, err := NewROUTER("inproc://router1,inproc://router2")
	if err != nil {
		t.Errorf("NewROUTER failed: %s", err)
	}
	defer router.Destroy()

	err = dealer.SendMessage("Hello")
	if err != nil {
		t.Errorf("SendMessage failed: %s", err)
	}

	msg, err := router.RecvMessage()
	if err != nil {
		t.Errorf("RecvMessage failed: %s", err)
	}
	if len(msg) != 2 {
		t.Error("message should have 2 frames")
	}

	if string(msg[1]) != "Hello" {
		t.Errorf("Expected 'Hello', received '%s", string(msg[0]))
	}

	msg[1] = []byte("World")

	err = router.SendMessage(msg)
	if err != nil {
		t.Errorf("SendMessage failed: %s", err)
	}

	msg, err = dealer.RecvMessage()
	if err != nil {
		t.Errorf("RecvMessage failed: %s", err)
	}

	if len(msg) != 1 {
		t.Error("message should have 1 frames")
	}

	if string(msg[0]) != "World" {
		t.Errorf("Expected 'World', received '%s", string(msg[0]))
	}
}

func TestXSUBXPUB(t *testing.T) {
	_, err := NewXPUB("bogus://bogus")
	if err == nil {
		t.Error("NewXPUB should have returned error and did not")
	}

	_, err = NewXSUB("bogus://bogus")
	if err == nil {
		t.Error("NewXSUB should have returned error and did not")
	}

	xpub, err := NewXPUB("inproc://xpub1,inproc://xpub2")
	if err != nil {
		t.Errorf("NewXPUB failed: %s", err)
	}
	defer xpub.Destroy()

	xsub, err := NewXSUB("inproc://xpub1,inproc://xpub2")
	if err != nil {
		t.Errorf("NewXSUB failed: %s", err)
	}
	defer xsub.Destroy()
}

func TestPAIR(t *testing.T) {
	_, err := NewPAIR("bogus://bogus")
	if err == nil {
		t.Error("NewPAIR should have returned error and did not")
	}

	pair1, err := NewPAIR(">inproc://pair")
	if err != nil {
		t.Errorf("NewPAIR failed: %s", err)
	}
	defer pair1.Destroy()

	pair2, err := NewPAIR("@inproc://pair")
	if err != nil {
		t.Errorf("NewPAIR failed: %s", err)
	}
	defer pair2.Destroy()
}

func TestSTREAM(t *testing.T) {
	_, err := NewSTREAM("bogus://bogus")
	if err == nil {
		t.Error("NewSTREAM should have returned error and did not")
	}

	stream1, err := NewSTREAM(">inproc://stream")
	if err != nil {
		t.Errorf("NewSTREAM failed: %s", err)
	}
	defer stream1.Destroy()

	stream2, err := NewSTREAM("@inproc://stream")
	if err != nil {
		t.Errorf("NewSTREAM failed: %s", err)
	}
	defer stream2.Destroy()

}
