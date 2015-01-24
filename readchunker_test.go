package goczmq

import (
	"bytes"
	"io"
	"testing"
)

func TestChunkerEven(t *testing.T) {
	in := NewSock(PULL)
	defer in.Destroy()

	_, err := in.Bind("inproc://chunker-test")
	if err != nil {
		t.Fatalf("in.Bind: %s", err)
	}

	out := NewSock(PUSH)
	err = out.Connect("inproc://chunker-test")
	if err != nil {
		t.Fatalf("out.Connect: %s", err)
	}

	chunker := NewReadChunker(out, 5)
	defer chunker.Destroy()

	data := []byte("1234567890")
	buf := bytes.NewBuffer(data)

	n, err := chunker.ReadFrom(buf)

	if err != io.EOF {
		t.Errorf("expected io.EOF received %s", err)
	}

	if n != 10 {
		t.Errorf("expected 10 bytes read got %d", n)
	}

	msg, err := in.RecvMessage()
	if err != nil {
		t.Errorf("in.RecvMessage: %s", err)
	}

	if string(msg[0]) != "12345" {
		t.Errorf("first frame should be '12345', is %s", string(msg[0]))
	}

	if string(msg[1]) != "67890" {
		t.Errorf("second frame should be '67890', is %s", string(msg[1]))
	}
}

func TestChunkerRemainder(t *testing.T) {
	in := NewSock(PULL)
	defer in.Destroy()

	_, err := in.Bind("inproc://chunker-test")
	if err != nil {
		t.Fatalf("in.Bind: %s", err)
	}

	out := NewSock(PUSH)
	err = out.Connect("inproc://chunker-test")
	if err != nil {
		t.Fatalf("out.Connect: %s", err)
	}

	chunker := NewReadChunker(out, 6)
	defer chunker.Destroy()

	data := []byte("1234567890")
	buf := bytes.NewBuffer(data)

	n, err := chunker.ReadFrom(buf)

	if err != io.EOF {
		t.Errorf("expected io.EOF received %s", err)
	}

	if n != 10 {
		t.Errorf("expected 10 bytes read got %d", n)
	}

	msg, err := in.RecvMessage()
	if err != nil {
		t.Errorf("in.RecvMessage: %s", err)
	}

	if string(msg[0]) != "123456" {
		t.Errorf("first frame should be '12345', is %s", string(msg[0]))
	}

	if string(msg[1]) != "7890" {
		t.Errorf("second frame should be '67890', is %s", string(msg[1]))
	}
}
