package goczmq

import (
	"bytes"
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

	err = pushSock.SendFrame([]byte("Hello"), 0)
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

	err = pushSock.SendFrame([]byte("World"), 0)
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
	err = pushSock.SendFrame(empty, 0)
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
	_, err := NewPub("bogus://bogus")
	if err == nil {
		t.Error("NewPub should have returned error and did not")
	}

	_, err = NewSub("bogus://bogus", "")
	if err == nil {
		t.Error("NewSub should have returned error and did not")
	}

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

	err = pub.SendFrame([]byte("test pub sub"), 0)
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
	_, err := NewReq("bogus://bogus")
	if err == nil {
		t.Error("NewReq should have returned error and did not")
	}

	_, err = NewRep("bogus://bogus")
	if err == nil {
		t.Error("NewRep should have returned error and did not")
	}

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

	err = req.SendFrame([]byte("Hello"), 0)
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

	err = rep.SendFrame([]byte("World"), 0)
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
	_, err := NewPush("bogus://bogus")
	if err == nil {
		t.Error("NewPush should have returned error and did not")
	}

	_, err = NewPull("bogus://bogus")
	if err == nil {
		t.Error("NewPull should have returned error and did not")
	}

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

	err = push.SendFrame([]byte("Hello"), 1)
	if err != nil {
		t.Errorf("SendFrame failed: %s", err)
	}

	err = push.SendFrame([]byte("World"), 0)
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
	_, err := NewDealer("bogus://bogus")
	if err == nil {
		t.Error("NewDealer should have returned error and did not")
	}

	_, err = NewRouter("bogus://bogus")
	if err == nil {
		t.Error("NewRouter should have returned error and did not")
	}

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

	err = dealer.SendFrame([]byte("Hello"), 0)
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
	_, err := NewXPub("bogus://bogus")
	if err == nil {
		t.Error("NewXPub should have returned error and did not")
	}

	_, err = NewXSub("bogus://bogus")
	if err == nil {
		t.Error("NewXSub should have returned error and did not")
	}

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
	_, err := NewPair("bogus://bogus")
	if err == nil {
		t.Error("NewPair should have returned error and did not")
	}

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
	_, err := NewStream("bogus://bogus")
	if err == nil {
		t.Error("NewStream should have returned error and did not")
	}

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

	err = push.SendFrame([]byte("Hello World"), 0)
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

	err = pushSock.SendFrame([]byte("Hello"), 0)
	if err != nil {
		t.Errorf("pushSock.SendFrame failed: %s", err)
	}

	b := make([]byte, 5)

	n, err := pullSock.Read(b)
	if n != 5 {
		t.Errorf("pullSock.Read expected 5 bytes read %d", n)
	}

	if err != io.EOF {
		t.Errorf("pullSock.Read expected io.EOF got %s", err)
	}

	if bytes.Compare(b, []byte("Hello")) != 0 {
		t.Errorf("expected 'Hello' received '%s'", b)
	}

	err = pushSock.SendFrame([]byte("Hello"), 1)
	if err != nil {
		t.Errorf("pushSock.SendFrame: %s", err)
	}

	err = pushSock.SendFrame([]byte(" World"), 0)
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

	err = dealerSock.SendFrame([]byte("Hello"), 0)
	if err != nil {
		t.Errorf("dealerSock.SendFrame failed: %s", err)
	}

	b := make([]byte, 5)

	n, err := routerSock.Read(b)
	if n != 5 {
		t.Errorf("routerSock.Read expected 5 bytes read %d", n)
	}

	if err != io.EOF {
		t.Errorf("routerSock.Read expected io.EOF got %s", err)
	}

	if bytes.Compare(b, []byte("Hello")) != 0 {
		t.Errorf("expected 'Hello' received '%s'", b)
	}

	err = dealerSock.SendFrame([]byte("Hello"), 1)
	if err != nil {
		t.Errorf("dealerSock.SendFrame: %s", err)
	}

	err = dealerSock.SendFrame([]byte(" World"), 0)
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

	err = dealerSock1.SendFrame([]byte("Hello From Client 1!"), 0)
	if err != nil {
		t.Errorf("dealerSock.SendFrame failed: %s", err)
	}

	err = dealerSock2.Connect("inproc://test-read-router-async")
	if err != nil {
		t.Errorf("reqSock.Connect failed: %s", err)
	}

	err = dealerSock2.SendFrame([]byte("Hello From Client 2!"), 0)
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

func ExampleSock_output() {
	dealer, err := NewDealer("inproc://example")
	if err != nil {
		panic(err)
	}

	router, err := NewRouter("inproc://example")
	if err != nil {
		panic(err)
	}
	dealer.SendFrame([]byte("Hello"), 0)

	request, err := router.RecvMessage()
	if err != nil {
		panic(err)
	}

	// first frame is identify of client - let's append 'World'
	// to the message and route it back.
	request = append(request, []byte("World"))
	err = router.SendMessage(request)
	if err != nil {
		panic(err)
	}

	reply, err := dealer.RecvMessage()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %s", string(reply[0]), string(reply[1]))
	// Output: Hello World
}

func benchmarkSendFrame(size int, b *testing.B) {
	pullSock := NewSock(Pull)
	defer pullSock.Destroy()

	_, err := pullSock.Bind("inproc://benchDealerRouter")
	if err != nil {
		panic(err)
	}

	go func() {
		pushSock := NewSock(Dealer)
		defer pushSock.Destroy()
		err := pushSock.Connect("inproc://benchDealerRouter")
		if err != nil {
			panic(err)
		}

		payload := make([]byte, size)
		for i := 0; i < b.N; i++ {
			err = pushSock.SendFrame(payload, 0)
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

func BenchmarkSendFrame1k(b *testing.B)  { benchmarkSendFrame(1024, b) }
func BenchmarkSendFrame4k(b *testing.B)  { benchmarkSendFrame(4096, b) }
func BenchmarkSendFrame16k(b *testing.B) { benchmarkSendFrame(16384, b) }
func BenchmarkSendFrame65k(b *testing.B) { benchmarkSendFrame(65536, b) }
