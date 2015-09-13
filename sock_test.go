package goczmq

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"testing"
)

func TestSendFrame(t *testing.T) {
	pushSock := NewSock(Push)
	defer pushSock.Destroy()

	pullSock := NewSock(Pull)
	defer pullSock.Destroy()

	_, err := pullSock.Bind("inproc://test-send-frame")
	if err != nil {
		t.Errorf("repSock.Bind failed: %s", err)
	}

	err = pushSock.Connect("inproc://test-send-frame")
	if err != nil {
		t.Errorf("reqSock.Connect failed: %s", err)
	}

	err = pushSock.SendFrame([]byte("Hello"), FlagNone)
	if err != nil {
		t.Errorf("pushSock.SendFrame failed: %s", err)
	}

	frame, flag, err := pullSock.RecvFrame()
	if err != nil {
		t.Errorf("pullSock.RecvFrame failed: %s", err)
	}

	if bytes.Compare(frame, []byte("Hello")) != 0 {
		t.Errorf("expected 'Hello' received '%s'", frame)
	}

	frame, flag, err = pullSock.RecvFrameNoWait()
	if err == nil {
		t.Errorf("RecvFrameNoWait should return error if no frame waiting")
	}

	if flag != 0 {
		t.Errorf("flag shouled have been 0, is '%d'", flag)
	}

	err = pushSock.SendFrame([]byte("World"), FlagNone)
	if err != nil {
		t.Errorf("pushSock.SendFrame failed: %s", err)
	}

	frame, flag, err = pullSock.RecvFrameNoWait()
	if err != nil {
		t.Errorf("pullsock.RecvFrameNoWait failed: %s", err)
	}

	if flag != 0 {
		t.Errorf("flag shouled have been 0, is '%d'", flag)
	}

	if string(frame) != "World" {
		t.Errorf("expected 'World' received '%s'", frame)
	}
}

func TestSendEmptyFrame(t *testing.T) {
	pushSock := NewSock(Push)
	defer pushSock.Destroy()

	pullSock := NewSock(Pull)
	defer pullSock.Destroy()

	_, err := pullSock.Bind("inproc://test-send-empty")
	if err != nil {
		t.Errorf("repSock.Bind failed: %s", err)
	}

	err = pushSock.Connect("inproc://test-send-empty")
	if err != nil {
		t.Errorf("reqSock.Connect failed: %s", err)
	}

	empty := make([]byte, 0)
	err = pushSock.SendFrame(empty, FlagNone)
	if err != nil {
		t.Errorf("pushSock.SendFrame failed: %s", err)
	}

	frame, _, err := pullSock.RecvFrame()
	if err != nil {
		t.Errorf("pullSock.RecvFrame failed: %s", err)
	}
	if len(frame) != 0 {
		t.Errorf("frame should be empty but had len %d", len(frame))
	}

}

func TestSendMessage(t *testing.T) {
	pushSock := NewSock(Push)
	defer pushSock.Destroy()

	pullSock := NewSock(Pull)
	defer pullSock.Destroy()

	_, err := pullSock.Bind("inproc://test-send-msg")
	if err != nil {
		t.Errorf("repSock.Bind failed: %s", err)
	}

	err = pushSock.Connect("inproc://test-send-msg")
	if err != nil {
		t.Errorf("reqSock.Connect failed: %s", err)
	}

	pushSock.SendMessage([][]byte{[]byte("Hello")})
	msg, err := pullSock.RecvMessage()
	if err != nil {
		t.Errorf("pullsock.RecvMessage() failed: %s", err)
	}

	if bytes.Compare(msg[0], []byte("Hello")) != 0 {
		t.Errorf("expected 'Hello' received '%s'", msg)
	}

	msg, err = pullSock.RecvMessageNoWait()
	if err == nil {
		t.Errorf("RecvMessageNoWait should return error if no frame waiting")
	}

	pushSock.SendMessage([][]byte{[]byte("World")})
	msg, err = pullSock.RecvMessageNoWait()
	if err != nil {
		t.Errorf("pullsock.RecvMessage() failed: %s", err)
	}

	if bytes.Compare(msg[0], []byte("World")) != 0 {
		t.Errorf("expected 'World' received '%s'", msg)
	}
}

func TestPubSub(t *testing.T) {
	bogusPub, err := NewPub("bogus://bogus")
	if err == nil {
		t.Error("NewPub should have returned error and did not")
	}
	defer bogusPub.Destroy()

	bogusSub, err := NewSub("bogus://bogus", "")
	if err == nil {
		t.Error("NewSub should have returned error and did not")
	}
	defer bogusSub.Destroy()

	pub, err := NewPub("inproc://pub1,inproc://pub2")
	if err != nil {
		t.Errorf("NewPub failed: %s", err)
	}
	defer pub.Destroy()

	sub, err := NewSub("inproc://pub1,inproc://pub2", "")
	if err != nil {
		t.Errorf("NewSub failed: %s", err)
	}
	defer sub.Destroy()

	err = pub.SendFrame([]byte("test pub sub"), FlagNone)
	if err != nil {
		t.Errorf("SendFrame failed: %s", err)
	}

	frame, _, err := sub.RecvFrame()
	if err != nil {
		t.Errorf("RecvFrame failed: %s", err)
	}

	if string(frame) != "test pub sub" {
		t.Errorf("Expected 'test pub sub', received %s", frame)
	}
}

func TestReqRep(t *testing.T) {
	bogusReq, err := NewReq("bogus://bogus")
	if err == nil {
		t.Error("NewReq should have returned error and did not")
	}
	defer bogusReq.Destroy()

	bogusRep, err := NewRep("bogus://bogus")
	if err == nil {
		t.Error("NewRep should have returned error and did not")
	}
	defer bogusRep.Destroy()

	rep, err := NewRep("inproc://rep1,inproc://rep2")
	if err != nil {
		t.Errorf("NewRep failed: %s", err)
	}
	defer rep.Destroy()

	req, err := NewReq("inproc://rep1,inproc://rep2")
	if err != nil {
		t.Errorf("NewReq failed: %s", err)
	}
	defer req.Destroy()

	err = req.SendFrame([]byte("Hello"), FlagNone)
	if err != nil {
		t.Errorf("SendFrame failed: %s", err)
	}

	reqframe, _, err := rep.RecvFrame()
	if err != nil {
		t.Errorf("RecvFrame failed: %s", err)
	}

	if string(reqframe) != "Hello" {
		t.Errorf("Expected 'Hello', received '%s", string(reqframe))
	}

	err = rep.SendFrame([]byte("World"), FlagNone)
	if err != nil {
		t.Errorf("SendFrame failed: %s", err)
	}

	repframe, _, err := req.RecvFrame()
	if err != nil {
		t.Errorf("RecvFrame failed: %s", err)
	}

	if string(repframe) != "World" {
		t.Errorf("Expected 'World', received '%s", string(repframe))
	}

}

func TestPushPull(t *testing.T) {
	bogusPush, err := NewPush("bogus://bogus")
	if err == nil {
		t.Error("NewPush should have returned error and did not")
	}
	defer bogusPush.Destroy()

	bogusPull, err := NewPull("bogus://bogus")
	if err == nil {
		t.Error("NewPull should have returned error and did not")
	}
	defer bogusPull.Destroy()

	push, err := NewPush("inproc://push1,inproc://push2")
	if err != nil {
		t.Errorf("NewPush failed: %s", err)
	}
	defer push.Destroy()

	pull, err := NewPull("inproc://push1,inproc://push2")
	if err != nil {
		t.Errorf("NewPull failed: %s", err)
	}
	defer pull.Destroy()

	err = push.SendFrame([]byte("Hello"), FlagMore)
	if err != nil {
		t.Errorf("SendFrame failed: %s", err)
	}

	err = push.SendFrame([]byte("World"), FlagNone)
	if err != nil {
		t.Errorf("SendFrame failed: %s", err)
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

func TestRouterDealer(t *testing.T) {
	bogusDealer, err := NewDealer("bogus://bogus")
	if err == nil {
		t.Error("NewDealer should have returned error and did not")
	}
	defer bogusDealer.Destroy()

	bogusRouter, err := NewRouter("bogus://bogus")
	if err == nil {
		t.Error("NewRouter should have returned error and did not")
	}
	defer bogusRouter.Destroy()

	dealer, err := NewDealer("inproc://router1,inproc://router2")
	if err != nil {
		t.Errorf("NewDealer failed: %s", err)
	}
	defer dealer.Destroy()

	router, err := NewRouter("inproc://router1,inproc://router2")
	if err != nil {
		t.Errorf("NewRouter failed: %s", err)
	}
	defer router.Destroy()

	err = dealer.SendFrame([]byte("Hello"), FlagNone)
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

func TestXSubXPub(t *testing.T) {
	bogusXPub, err := NewXPub("bogus://bogus")
	if err == nil {
		t.Error("NewXPub should have returned error and did not")
	}
	defer bogusXPub.Destroy()

	bogusXSub, err := NewXSub("bogus://bogus")
	if err == nil {
		t.Error("NewXSub should have returned error and did not")
	}
	defer bogusXSub.Destroy()

	xpub, err := NewXPub("inproc://xpub1,inproc://xpub2")
	if err != nil {
		t.Errorf("NewXPub failed: %s", err)
	}
	defer xpub.Destroy()

	xsub, err := NewXSub("inproc://xpub1,inproc://xpub2")
	if err != nil {
		t.Errorf("NewXSub failed: %s", err)
	}
	defer xsub.Destroy()
}

func TestPair(t *testing.T) {
	bogusPair, err := NewPair("bogus://bogus")
	if err == nil {
		t.Error("NewPair should have returned error and did not")
	}
	defer bogusPair.Destroy()

	pair1, err := NewPair(">inproc://pair")
	if err != nil {
		t.Errorf("NewPair failed: %s", err)
	}
	defer pair1.Destroy()

	pair2, err := NewPair("@inproc://pair")
	if err != nil {
		t.Errorf("NewPair failed: %s", err)
	}
	defer pair2.Destroy()
}

func TestStream(t *testing.T) {
	bogusStream, err := NewStream("bogus://bogus")
	if err == nil {
		t.Error("NewStream should have returned error and did not")
	}
	defer bogusStream.Destroy()

	stream1, err := NewStream(">inproc://stream")
	if err != nil {
		t.Errorf("NewStream failed: %s", err)
	}
	defer stream1.Destroy()

	stream2, err := NewStream("@inproc://stream")
	if err != nil {
		t.Errorf("NewStream failed: %s", err)
	}
	defer stream2.Destroy()

}

func TestPollin(t *testing.T) {
	push, err := NewPush("inproc://pollin")
	if err != nil {
		t.Errorf("NewPush failed: %s", err)
	}
	defer push.Destroy()

	pull, err := NewPull("inproc://pollin")
	if err != nil {
		t.Errorf("NewPull failed: %s", err)
	}
	defer pull.Destroy()

	if pull.Pollin() {
		t.Errorf("Pollin returned true should be false")
	}

	err = push.SendFrame([]byte("Hello World"), FlagNone)
	if err != nil {
		t.Errorf("SendFrame failed: %s", err)
	}

	if !pull.Pollin() {
		t.Errorf("Pollin returned false should be true")
	}
}

func TestPollout(t *testing.T) {
	push := NewSock(Push)
	_, err := push.Bind("inproc://pollout")
	if err != nil {
		t.Errorf("failed binding test socket: %s", err)
	}
	defer push.Destroy()

	if push.Pollout() {
		t.Errorf("Pollout returned true should be false")
	}

	pull := NewSock(Pull)
	defer pull.Destroy()

	err = pull.Connect("inproc://pollout")
	if err != nil {
		t.Errorf("failed connecting test socket: %s", err)
	}

	if !push.Pollout() {
		t.Errorf("Pollout returned false should be true")
	}
}

func TestReader(t *testing.T) {
	pushSock := NewSock(Push)
	defer pushSock.Destroy()

	pullSock := NewSock(Pull)
	defer pullSock.Destroy()

	_, err := pullSock.Bind("inproc://test-read")
	if err != nil {
		t.Errorf("repSock.Bind failed: %s", err)
	}

	err = pushSock.Connect("inproc://test-read")
	if err != nil {
		t.Errorf("reqSock.Connect failed: %s", err)
	}

	err = pushSock.SendFrame([]byte("Hello"), FlagNone)
	if err != nil {
		t.Errorf("pushSock.SendFrame failed: %s", err)
	}

	b := make([]byte, 5)

	n, err := pullSock.Read(b)
	if n != 5 {
		t.Errorf("pullSock.Read expected 5 bytes read %d", n)
	}

	if err != nil {
		t.Errorf("pullSock.Read error: %s", err)
	}

	if bytes.Compare(b, []byte("Hello")) != 0 {
		t.Errorf("expected 'Hello' received '%s'", b)
	}

	err = pushSock.SendFrame([]byte("Hello"), FlagMore)
	if err != nil {
		t.Errorf("pushSock.SendFrame: %s", err)
	}

	err = pushSock.SendFrame([]byte(" World"), FlagNone)
	if err != nil {
		t.Errorf("pushSock.SendFrame: %s", err)
	}

	b = make([]byte, 8)
	n, err = pullSock.Read(b)
	if err != ErrSliceFull {
		t.Errorf("expected %s error, got %s", ErrSliceFull, err)
	}

	if bytes.Compare(b, []byte("Hello Wo")) != 0 {
		t.Errorf("expected 'Hello Wo' received '%s'", b)
	}
}

func TestReaderWithBufferSmallerThanFrame(t *testing.T) {
	dealerSock := NewSock(Dealer)
	defer dealerSock.Destroy()

	routerSock := NewSock(Router)
	defer routerSock.Destroy()

	_, err := routerSock.Bind("inproc://test-read-small-buf")
	if err != nil {
		t.Errorf("repSock.Bind failed: %s", err)
	}

	err = dealerSock.Connect("inproc://test-read-small-buf")
	if err != nil {
		t.Errorf("reqSock.Connect failed: %s", err)
	}

	err = dealerSock.SendFrame([]byte("Hello"), FlagNone)
	if err != nil {
		t.Errorf("dealerSock.SendFrame failed: %s", err)
	}

	b := make([]byte, 2)

	n, err := routerSock.Read(b)
	if n != 2 {
		t.Errorf("routerSock.Read expected 2 bytes read %d", n)
	}

	if err != ErrSliceFull {
		t.Errorf("routerSock.Read expected io.EOF got %s", err)
	}

	if bytes.Compare(b, []byte("He")) != 0 {
		t.Errorf("expected 'Hello' received '%s'", b)
	}
}

func TestReaderWithRouterDealer(t *testing.T) {
	dealerSock := NewSock(Dealer)
	defer dealerSock.Destroy()

	routerSock := NewSock(Router)
	defer routerSock.Destroy()

	_, err := routerSock.Bind("inproc://test-read-router")
	if err != nil {
		t.Errorf("repSock.Bind failed: %s", err)
	}

	err = dealerSock.Connect("inproc://test-read-router")
	if err != nil {
		t.Errorf("reqSock.Connect failed: %s", err)
	}

	err = dealerSock.SendFrame([]byte("Hello"), FlagNone)
	if err != nil {
		t.Errorf("dealerSock.SendFrame failed: %s", err)
	}

	b := make([]byte, 5)

	n, err := routerSock.Read(b)
	if n != 5 {
		t.Errorf("routerSock.Read expected 5 bytes read %d", n)
	}

	if err != nil {
		t.Errorf("routerSock.Read expected io.EOF got %s", err)
	}

	if bytes.Compare(b, []byte("Hello")) != 0 {
		t.Errorf("expected 'Hello' received '%s'", b)
	}

	err = dealerSock.SendFrame([]byte("Hello"), FlagMore)
	if err != nil {
		t.Errorf("dealerSock.SendFrame: %s", err)
	}

	err = dealerSock.SendFrame([]byte(" World"), FlagNone)
	if err != nil {
		t.Errorf("dealerSock.SendFrame: %s", err)
	}

	b = make([]byte, 8)
	n, err = routerSock.Read(b)
	if err != ErrSliceFull {
		t.Errorf("expected %s error, got %s", ErrSliceFull, err)
	}

	if bytes.Compare(b, []byte("Hello Wo")) != 0 {
		t.Errorf("expected 'Hello Wo' received '%s'", b)
	}

	n, err = routerSock.Write([]byte("World"))
	if err != nil {
		t.Errorf("routerSock.Write: %s", err)
	}
	if n != 5 {
		t.Errorf("expected 5 bytes sent got %d", n)
	}

	frame, _, err := dealerSock.RecvFrame()
	if err != nil {
		t.Errorf("dealer.RecvFrame: %s", err)
	}

	if bytes.Compare(frame, []byte("World")) != 0 {
		t.Errorf("expected 'World' received '%s'", b)
	}
}

func TestReaderWithRouterDealerAsync(t *testing.T) {
	dealerSock1 := NewSock(Dealer)
	defer dealerSock1.Destroy()

	dealerSock2 := NewSock(Dealer)
	defer dealerSock2.Destroy()

	routerSock := NewSock(Router)
	defer routerSock.Destroy()

	_, err := routerSock.Bind("inproc://test-read-router-async")
	if err != nil {
		t.Errorf("repSock.Bind failed: %s", err)
	}

	err = dealerSock1.Connect("inproc://test-read-router-async")
	if err != nil {
		t.Errorf("reqSock.Connect failed: %s", err)
	}

	err = dealerSock1.SendFrame([]byte("Hello From Client 1!"), FlagNone)
	if err != nil {
		t.Errorf("dealerSock.SendFrame failed: %s", err)
	}

	err = dealerSock2.Connect("inproc://test-read-router-async")
	if err != nil {
		t.Errorf("reqSock.Connect failed: %s", err)
	}

	err = dealerSock2.SendFrame([]byte("Hello From Client 2!"), FlagNone)
	if err != nil {
		t.Errorf("dealerSock.SendFrame failed: %s", err)
	}

	msg := make([]byte, 255)

	n, err := routerSock.Read(msg)
	if n != 20 {
		t.Errorf("routerSock.Read expected 20 bytes read %d", n)
	}

	client1ID := routerSock.GetLastClientID()

	if bytes.Compare(msg[:n], []byte("Hello From Client 1!")) != 0 {
		t.Errorf("expected 'Hello From Client 1!' received '%s'", string(msg[:n]))
	}

	n, err = routerSock.Read(msg)
	if n != 20 {
		t.Errorf("routerSock.Read expected 20 bytes read %d", n)
	}

	client2ID := routerSock.GetLastClientID()

	if bytes.Compare(msg[:n], []byte("Hello From Client 2!")) != 0 {
		t.Errorf("expected 'Hello From Client 2!' received '%s'", string(msg[:n]))
	}

	routerSock.SetLastClientID(client1ID)
	n, err = routerSock.Write([]byte("Hello Client 1!"))
	if err != nil {
		t.Errorf("routerSock.Write: %s", err)
	}

	frame, _, err := dealerSock1.RecvFrame()
	if err != nil {
		t.Errorf("dealer.RecvFrame: %s", err)
	}

	if bytes.Compare(frame, []byte("Hello Client 1!")) != 0 {
		t.Errorf("expected 'World' received '%s'", frame)
	}

	routerSock.SetLastClientID(client2ID)
	n, err = routerSock.Write([]byte("Hello Client 2!"))
	if err != nil {
		t.Errorf("routerSock.Write: %s", err)
	}

	frame, _, err = dealerSock2.RecvFrame()
	if err != nil {
		t.Errorf("dealer.RecvFrame: %s", err)
	}

	if bytes.Compare(frame, []byte("Hello Client 2!")) != 0 {
		t.Errorf("expected 'World' received '%s'", frame)
	}
}

type encodeMessage struct {
	Foo string
	Bar []byte
	Bat int
}

func TestPushPullEncodeDecode(t *testing.T) {
	push, err := NewPush("inproc://pushpullencode")
	if err != nil {
		t.Error(err)
	}
	defer push.Destroy()

	pull, err := NewPull("inproc://pushpullencode")
	if err != nil {
		t.Error(err)
	}
	defer pull.Destroy()

	encoder := gob.NewEncoder(push)
	decoder := gob.NewDecoder(pull)

	sent := encodeMessage{
		Foo: "the answer",
		Bar: []byte("is"),
		Bat: 42,
	}

	err = encoder.Encode(sent)
	if err != nil {
		t.Errorf("could not encode test message to buffer: %s", err)
	}

	var received encodeMessage
	err = decoder.Decode(&received)
	if err != nil {
		t.Errorf("could node decode test message from buffer: %s", err)
	}

	if received.Foo != sent.Foo {
		t.Errorf("expected '%s', got '%s'", sent.Foo, received.Foo)
	}

	if string(received.Bar) != string(sent.Bar) {
		t.Errorf("expected '%s', got '%s'", string(sent.Bar), string(received.Bar))
	}

	if received.Bat != sent.Bat {
		t.Errorf("expected '%d', got '%d'", sent.Bat, received.Bat)
	}
}

func TestDealerRouterEncodeDecode(t *testing.T) {
	router, err := NewRouter("inproc://dealerrouterencode")
	if err != nil {
		t.Error(err)
	}
	defer router.Destroy()

	dealer, err := NewDealer("inproc://dealerrouterencode")
	if err != nil {
		t.Error(err)
	}
	defer dealer.Destroy()

	rencoder := gob.NewEncoder(router)
	rdecoder := gob.NewDecoder(router)

	dencoder := gob.NewEncoder(dealer)
	ddecoder := gob.NewDecoder(dealer)

	question := encodeMessage{
		Foo: "what is",
		Bar: []byte("the answer"),
		Bat: 0,
	}

	err = dencoder.Encode(question)
	if err != nil {
		t.Errorf("could not encode test message: %s", err)
	}

	var received encodeMessage
	err = rdecoder.Decode(&received)
	if err != nil {
		t.Errorf("could not decode: %s", err)
	}

	if received.Foo != question.Foo {
		t.Errorf("expected '%s', got '%s'", question.Foo, received.Foo)
	}

	if string(received.Bar) != string(question.Bar) {
		t.Errorf("expected '%s', got '%s'", string(question.Bar), string(received.Bar))
	}

	if received.Bat != question.Bat {
		t.Errorf("expected '%d', got '%d'", question.Bat, received.Bat)
	}

	sent := encodeMessage{
		Foo: "the answer",
		Bar: []byte("is"),
		Bat: 42,
	}

	err = rencoder.Encode(sent)
	if err != nil {
		t.Errorf("could not encode test message: %s", err)
	}

	var answer encodeMessage
	err = ddecoder.Decode(&answer)
	if err != nil {
		t.Errorf("could not decode: %s", err)
	}

	if answer.Foo != sent.Foo {
		t.Errorf("expected '%s', got '%s'", sent.Foo, answer.Foo)
	}

	if string(answer.Bar) != string(sent.Bar) {
		t.Errorf("expected '%s', got '%s'", string(sent.Bar), string(answer.Bar))
	}

	if answer.Bat != sent.Bat {
		t.Errorf("expected '%d', got '%d'", sent.Bat, answer.Bat)
	}
}

func ExampleSock_output() {
	// create dealer socket
	dealer, err := NewDealer("inproc://example")
	if err != nil {
		panic(err)
	}
	defer dealer.Destroy()

	// create router socket
	router, err := NewRouter("inproc://example")
	if err != nil {
		panic(err)
	}
	defer router.Destroy()

	// send hello message
	dealer.SendFrame([]byte("Hello"), FlagNone)

	// receive hello message
	request, err := router.RecvMessage()
	if err != nil {
		panic(err)
	}

	// first frame is identify of client - let's append 'World'
	// to the message and route it back.
	request = append(request, []byte("World"))

	// send reply
	err = router.SendMessage(request)
	if err != nil {
		panic(err)
	}

	// receive reply
	reply, err := dealer.RecvMessage()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %s", string(reply[0]), string(reply[1]))
	// Output: Hello World
}

func benchmarkSockSendFrame(size int, b *testing.B) {
	pullSock := NewSock(Pull)
	defer pullSock.Destroy()

	_, err := pullSock.Bind("inproc://benchSock")
	if err != nil {
		panic(err)
	}

	go func() {
		pushSock := NewSock(Push)
		defer pushSock.Destroy()
		err := pushSock.Connect("inproc://benchSock")
		if err != nil {
			panic(err)
		}

		payload := make([]byte, size)
		for i := 0; i < b.N; i++ {
			err = pushSock.SendFrame(payload, FlagNone)
			if err != nil {
				panic(err)
			}
		}
	}()

	for i := 0; i < b.N; i++ {
		msg, _, err := pullSock.RecvFrame()
		if err != nil {
			panic(err)
		}
		if len(msg) != size {
			panic("msg too small")
		}
	}
}

func BenchmarkSockSendFrame1k(b *testing.B)  { benchmarkSockSendFrame(1024, b) }
func BenchmarkSockSendFrame4k(b *testing.B)  { benchmarkSockSendFrame(4096, b) }
func BenchmarkSockSendFrame16k(b *testing.B) { benchmarkSockSendFrame(16384, b) }
func BenchmarkSockSendFrame65k(b *testing.B) { benchmarkSockSendFrame(65536, b) }

func benchmarkSockReadWriter(size int, b *testing.B) {
	pullSock := NewSock(Pull)
	defer pullSock.Destroy()

	_, err := pullSock.Bind("inproc://benchSock")
	if err != nil {
		panic(err)
	}

	go func() {
		pushSock := NewSock(Push)
		defer pushSock.Destroy()
		err := pushSock.Connect("inproc://benchSock")
		if err != nil {
			panic(err)
		}

		payload := make([]byte, size)
		for i := 0; i < b.N; i++ {
			_, err = pushSock.Write(payload)
			if err != nil {
				panic(err)
			}
		}
	}()

	payload := make([]byte, size)
	for i := 0; i < b.N; i++ {
		n, err := pullSock.Read(payload)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n != size {
			panic("msg too small")
		}
	}
}

func BenchmarkSockReadWriter1k(b *testing.B)  { benchmarkSockReadWriter(1024, b) }
func BenchmarkSockReadWriter4k(b *testing.B)  { benchmarkSockReadWriter(4096, b) }
func BenchmarkSockReadWriter16k(b *testing.B) { benchmarkSockReadWriter(16384, b) }
func BenchmarkSockReadWriter65k(b *testing.B) { benchmarkSockReadWriter(65536, b) }

func BenchmarkEncodeDecode(b *testing.B) {
	pullSock := NewSock(Pull)
	defer pullSock.Destroy()

	decoder := gob.NewDecoder(pullSock)

	_, err := pullSock.Bind("inproc://benchSock")
	if err != nil {
		panic(err)
	}

	go func() {
		pushSock := NewSock(Push)
		defer pushSock.Destroy()
		err := pushSock.Connect("inproc://benchSock")
		if err != nil {
			panic(err)
		}

		encoder := gob.NewEncoder(pushSock)

		sent := encodeMessage{
			Foo: "the answer",
			Bar: make([]byte, 1024),
			Bat: 42,
		}

		for i := 0; i < b.N; i++ {
			err := encoder.Encode(sent)
			if err != nil {
				panic(err)
			}
		}
	}()

	var received encodeMessage
	for i := 0; i < b.N; i++ {
		err := decoder.Decode(&received)
		if err != nil {
			panic(err)
		}
	}
}
