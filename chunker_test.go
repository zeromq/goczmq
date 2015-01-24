package goczmq

import (
	"bytes"
	"io"
	"testing"
)

func TestChunker(t *testing.T) {
	in := NewSock(PULL)
	defer in.Destroy()

	_, err := in.Bind("inproc://chunker-test")
	if err != nil {
		t.Fatalf("in.Bind: %s", err)
	}

	writeChunker := NewWriteChunker(in)
	defer writeChunker.Destroy()

	out := NewSock(PUSH)
	err = out.Connect("inproc://chunker-test")
	if err != nil {
		t.Fatalf("out.Connect: %s", err)
	}

	readChunker := NewReadChunker(out, 5)
	defer readChunker.Destroy()

	data := []byte("1234567890")
	buf := bytes.NewBuffer(data)

	n, err := readChunker.ReadFrom(buf)
	if err != nil && err != io.EOF {
		t.Errorf("chunker.ReadFrom: %s", err)
	}

	if n != 10 {
		t.Errorf("expected 10 bytes, got %d", n)
	}

	buf = bytes.NewBuffer(make([]byte, 10))

	n, err = writeChunker.WriteTo(buf)
	if err != nil {
		t.Errorf("chunker.WriteTo: %s", err)
	}

	if n != 10 {
		t.Errorf("expected 10 bytes, got %d", n)
	}
}
