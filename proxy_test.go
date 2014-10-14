package goczmq

import (
	"testing"
)

func TestZproxy(t *testing.T) {
	// Create and configure our proxy
	proxy := NewZproxy()

	err := proxy.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	err = proxy.SetFrontend(PULL, "inproc://frontend")
	if err != nil {
		t.Errorf("FRONTEND error: %s", err)
	}

	err = proxy.SetBackend(PUSH, "inproc://backend")
	if err != nil {
		t.Errorf("BACKEND error: %s", err)
	}

	err = proxy.SetCapture("inproc://capture")
	if err != nil {
		t.Errorf("CAPTURE error: %s", err)
	}

	// connect application sockets to proxy
	faucet := NewSock(PUSH)
	err = faucet.Connect("inproc://frontend")
	if err != nil {
		t.Error(err)
	}

	sink := NewSock(PULL)
	err = sink.Connect("inproc://backend")
	if err != nil {
		t.Error(err)
	}

	tap := NewSock(PULL)
	_, err = tap.Bind("inproc://capture")
	if err != nil {
		t.Error(err)
	}

	// send some messages and check they arrived
	faucet.SendBytes([]byte("Hello"), 0)
	faucet.SendBytes([]byte("World"), 0)

	// check the tap
	b, f, err := tap.RecvBytes()
	if err != nil {
		t.Error(err)
	}

	if f == MORE {
		t.Error("MORE set and should not be")
	}

	if string(b) != "Hello" {
		t.Errorf("tap expected %s, received %s", "Hello", string(b))
	}

	b, f, err = tap.RecvBytes()
	if err != nil {
		t.Error(err)
	}

	if f == MORE {
		t.Error("MORE set and should not be")
	}

	if string(b) != "World" {
		t.Errorf("tap expected %s, received %s", "World", string(b))
	}

	// check the sink
	b, f, err = sink.RecvBytes()
	if err != nil {
		t.Error(err)
	}

	if f == MORE {
		t.Error("MORE set and should not be")
	}

	if string(b) != "Hello" {
		t.Errorf("sink expected %s, received %s", "Hello", string(b))
	}

	b, f, err = sink.RecvBytes()
	if err != nil {
		t.Error(err)
	}

	if f == MORE {
		t.Error("MORE set and should not be")
	}

	if string(b) != "World" {
		t.Errorf("sink expected %s, received %s", "World", string(b))
	}

	// Test pause/resume functionality
	// FIXME: improve this test once we can receive with a nowait
	err = proxy.Pause()
	if err != nil {
		t.Error(err)
	}

	err = proxy.Resume()
	if err != nil {
		t.Error(err)
	}

	proxy.Destroy()
}
