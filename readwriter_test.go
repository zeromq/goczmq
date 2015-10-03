package goczmq

import (
	"bytes"
	"testing"
)

func TestReadWriter(t *testing.T) {
	endpoint := "inproc://testReadWriter"

	pushSock, err := NewPush(endpoint)
	if err != nil {
		t.Errorf("NewPush failed: %s", err)
	}
	defer pushSock.Destroy()

	pullSock, err := NewPull(endpoint)
	if err != nil {
		t.Errorf("NewPull failed: %s", err)
	}

	pullReadWriter := NewReadWriter(pullSock)
	defer pullReadWriter.Destroy()

	err = pushSock.SendFrame([]byte("Hello"), FlagNone)
	if err != nil {
		t.Errorf("pushSock.SendFrame failed: %s", err)
	}

	b := make([]byte, 5)

	n, err := pullReadWriter.Read(b)
	if n != 5 {
		t.Errorf("pullReadWriter.Read expected 5 bytes read %d", n)
	}

	if err != nil {
		t.Errorf("pullReadWriter.Read error: %s", err)
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
	n, err = pullReadWriter.Read(b)
	if err != ErrSliceFull {
		t.Errorf("expected %s error, got %s", ErrSliceFull, err)
	}

	if bytes.Compare(b, []byte("Hello Wo")) != 0 {
		t.Errorf("expected 'Hello Wo' received '%s'", b)
	}
}

func TestReadWriterWithBufferSmallerThanFrame(t *testing.T) {
	endpoint := "inproc://testReadWriterSmallBuf"

	dealerSock, err := NewDealer(endpoint)
	if err != nil {
		t.Errorf("NewDealer failed: %s", err)
	}
	defer dealerSock.Destroy()

	routerSock, err := NewRouter(endpoint)
	if err != nil {
		t.Errorf("NewRouter failed: %s", err)
	}

	routerReadWriter := NewReadWriter(routerSock)
	defer routerReadWriter.Destroy()

	err = dealerSock.SendFrame([]byte("Hello"), FlagNone)
	if err != nil {
		t.Errorf("dealerSock.SendFrame failed: %s", err)
	}

	b := make([]byte, 2)

	n, err := routerReadWriter.Read(b)
	if n != 2 {
		t.Errorf("routerReadWriter.Read expected 2 bytes read %d", n)
	}

	if err != ErrSliceFull {
		t.Errorf("routerReadWriter.Read expected io.EOF got %s", err)
	}

	if bytes.Compare(b, []byte("He")) != 0 {
		t.Errorf("expected 'Hello' received '%s'", b)
	}
}

func TestReadWriterWithRouterDealer(t *testing.T) {
	endpoint := "inproc://testReadWriterWithRouterDealer"

	dealerSock, err := NewDealer(endpoint)
	if err != nil {
		t.Errorf("NewDealer failed: %s", err)
	}
	defer dealerSock.Destroy()

	routerSock, err := NewRouter(endpoint)
	if err != nil {
		t.Errorf("NewRouter failed: %s", err)
	}

	routerReadWriter := NewReadWriter(routerSock)
	defer routerReadWriter.Destroy()

	err = dealerSock.SendFrame([]byte("Hello"), FlagNone)
	if err != nil {
		t.Errorf("dealerSock.SendFrame failed: %s", err)
	}

	b := make([]byte, 5)

	n, err := routerReadWriter.Read(b)
	if n != 5 {
		t.Errorf("routerReadWriter.Read expected 5 bytes read %d", n)
	}

	if err != nil {
		t.Errorf("routerReadWriter.Read expected io.EOF got %s", err)
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
	n, err = routerReadWriter.Read(b)
	if err != ErrSliceFull {
		t.Errorf("expected %s error, got %s", ErrSliceFull, err)
	}

	if bytes.Compare(b, []byte("Hello Wo")) != 0 {
		t.Errorf("expected 'Hello Wo' received '%s'", b)
	}

	n, err = routerReadWriter.Write([]byte("World"))
	if err != nil {
		t.Errorf("routerReadWriter.Write: %s", err)
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

func TestReadWriterWithRouterDealerAsync(t *testing.T) {
	endpoint := "inproc://testReadWriterWithRouterDealerAsync"

	dealerSock1, err := NewDealer(endpoint)
	if err != nil {
		t.Errorf("NewDealer failed: %s", err)
	}
	defer dealerSock1.Destroy()

	dealerSock2, err := NewDealer(endpoint)
	if err != nil {
		t.Errorf("NewDealer failed: %s", err)
	}
	defer dealerSock2.Destroy()

	routerSock, err := NewRouter(endpoint)
	if err != nil {
		t.Errorf("NewRouter failed: %s", err)
	}

	routerReadWriter := NewReadWriter(routerSock)
	defer routerReadWriter.Destroy()

	err = dealerSock1.SendFrame([]byte("Hello From Client 1!"), FlagNone)
	if err != nil {
		t.Errorf("dealerSock.SendFrame failed: %s", err)
	}

	err = dealerSock2.SendFrame([]byte("Hello From Client 2!"), FlagNone)
	if err != nil {
		t.Errorf("dealerSock.SendFrame failed: %s", err)
	}

	msg := make([]byte, 255)

	n, err := routerReadWriter.Read(msg)
	if n != 20 {
		t.Errorf("routerReadWriter.Read expected 20 bytes read %d", n)
	}

	client1ID := routerReadWriter.GetLastClientID()

	if bytes.Compare(msg[:n], []byte("Hello From Client 1!")) != 0 {
		t.Errorf("expected 'Hello From Client 1!' received '%s'", string(msg[:n]))
	}

	n, err = routerReadWriter.Read(msg)
	if n != 20 {
		t.Errorf("routerReadWriter.Read expected 20 bytes read %d", n)
	}

	client2ID := routerReadWriter.GetLastClientID()

	if bytes.Compare(msg[:n], []byte("Hello From Client 2!")) != 0 {
		t.Errorf("expected 'Hello From Client 2!' received '%s'", string(msg[:n]))
	}

	routerReadWriter.SetLastClientID(client1ID)
	n, err = routerReadWriter.Write([]byte("Hello Client 1!"))
	if err != nil {
		t.Errorf("routerReadWriter.Write: %s", err)
	}

	frame, _, err := dealerSock1.RecvFrame()
	if err != nil {
		t.Errorf("dealer.RecvFrame: %s", err)
	}

	if bytes.Compare(frame, []byte("Hello Client 1!")) != 0 {
		t.Errorf("expected 'World' received '%s'", frame)
	}

	routerReadWriter.SetLastClientID(client2ID)
	n, err = routerReadWriter.Write([]byte("Hello Client 2!"))
	if err != nil {
		t.Errorf("routerReadWriter.Write: %s", err)
	}

	frame, _, err = dealerSock2.RecvFrame()
	if err != nil {
		t.Errorf("dealer.RecvFrame: %s", err)
	}

	if bytes.Compare(frame, []byte("Hello Client 2!")) != 0 {
		t.Errorf("expected 'World' received '%s'", frame)
	}
}
