package goczmq

import (
	"testing"
)

func TestZproxy(t *testing.T) {
	// Create and configure our proxy
	proxy := NewProxy()
	defer proxy.Destroy()

	err := proxy.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	err = proxy.SetFrontend(Pull, "inproc://frontend")
	if err != nil {
		t.Errorf("FRONTEND error: %s", err)
	}

	err = proxy.SetBackend(Push, "inproc://backend")
	if err != nil {
		t.Errorf("BACKEND error: %s", err)
	}

	err = proxy.SetCapture("inproc://capture")
	if err != nil {
		t.Errorf("CAPTURE error: %s", err)
	}

	// connect application sockets to proxy
	faucet := NewSock(Push)
	err = faucet.Connect("inproc://frontend")
	if err != nil {
		t.Error(err)
	}
	defer faucet.Destroy()

	sink := NewSock(Pull)
	err = sink.Connect("inproc://backend")
	if err != nil {
		t.Error(err)
	}
	defer sink.Destroy()

	tap := NewSock(Pull)
	_, err = tap.Bind("inproc://capture")
	if err != nil {
		t.Error(err)
	}
	defer tap.Destroy()

	// send some messages and check they arrived
	faucet.SendFrame([]byte("Hello"), 0)
	faucet.SendFrame([]byte("World"), 0)

	// check the tap
	b, f, err := tap.RecvFrame()
	if err != nil {
		t.Error(err)
	}

	if f == More {
		t.Error("More set and should not be")
	}

	if string(b) != "Hello" {
		t.Errorf("tap expected %s, received %s", "Hello", string(b))
	}

	b, f, err = tap.RecvFrame()
	if err != nil {
		t.Error(err)
	}

	if f == More {
		t.Error("More set and should not be")
	}

	if string(b) != "World" {
		t.Errorf("tap expected %s, received %s", "World", string(b))
	}

	// check the sink
	b, f, err = sink.RecvFrame()
	if err != nil {
		t.Error(err)
	}

	if f == More {
		t.Error("More set and should not be")
	}

	if string(b) != "Hello" {
		t.Errorf("sink expected %s, received %s", "Hello", string(b))
	}

	b, f, err = sink.RecvFrame()
	if err != nil {
		t.Error(err)
	}

	if f == More {
		t.Error("More set and should not be")
	}

	if string(b) != "World" {
		t.Errorf("sink expected %s, received %s", "World", string(b))
	}

	// Test pause/resume functionality
	err = proxy.Pause()
	if err != nil {
		t.Error(err)
	}

	faucet.SendFrame([]byte("Belated Hello"), 0)

	if sink.Pollin() {
		t.Error("Paused proxy should not pass message but did")
	}

	if tap.Pollin() {
		t.Error("Paused proxy should not pass message but did")
	}

	err = proxy.Resume()
	if err != nil {
		t.Error(err)
	}

	b, f, err = sink.RecvFrame()
	if err != nil {
		t.Error(err)
	}

	if f == More {
		t.Error("More set and should not be")
	}

	if string(b) != "Belated Hello" {
		t.Errorf("sink expected %s, received %s", "Belated Hello", string(b))
	}

	b, f, err = tap.RecvFrame()
	if err != nil {
		t.Error(err)
	}

	if f == More {
		t.Error("More set and should not be")
	}

	if string(b) != "Belated Hello" {
		t.Errorf("tap expected %s, received %s", "Belated Hello", string(b))
	}

	proxy.Destroy()
}

func ProxyExample() {
	proxy := NewProxy()
	defer proxy.Destroy()

	// set front end address and socket type
	err := proxy.SetFrontend(Pull, "inproc://frontend")
	if err != nil {
		panic(err)
	}

	// set back end address and socket type
	err = proxy.SetBackend(Push, "inproc://backend")
	if err != nil {
		panic(err)
	}

	// set address for "tee"ing proxy traffic to
	err = proxy.SetCapture("inproc://capture")
	if err != nil {
		panic(err)
	}

	// we can pause the proxy
	err = proxy.Pause()
	if err != nil {
		panic(err)
	}

	// and we can resume it
	err = proxy.Resume()
	if err != nil {
		panic(err)
	}

	proxy.Destroy()
}
