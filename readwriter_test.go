package goczmq

import (
	"bytes"
	"fmt"
	"io"
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

	pullReadWriter, err := NewReadWriter(pullSock)
	if err != nil {
		t.Errorf("NewReadWriter failed: %s", err)
	}
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

	err = pushSock.SendFrame([]byte("Hello World"), FlagNone)
	if err != nil {
		t.Errorf("pushSock.SendFrame: %s", err)
	}

	b = make([]byte, 8)
	n, err = pullReadWriter.Read(b)

	if bytes.Compare(b, []byte("Hello Wo")) != 0 {
		t.Errorf("expected 'Hello Wo' received '%s'", b)
	}

	n, err = pullReadWriter.Read(b)
	if n != 3 {
		t.Errorf("expected to read 3 bytes, read %d", n)
	}

	if err != nil {
		t.Errorf("pullReadWriter.Read: %s", err)
	}

	if bytes.Compare(b[:n], []byte("rld")) != 0 {
		t.Errorf("expected 'Hello Wo' received '%s'", b)
	}

	pullReadWriter.SetTimeout(1)
	n, err = pullReadWriter.Read(b)
	if n != 0 {
		t.Errorf("expected to read 0 bytes, read %d", n)
	}

	if err != io.EOF {
		t.Errorf("expected io.EOF on timeout, got %v", err)
	}
}

func TestReadWriterWithBufferSmallerThanFrame(t *testing.T) {
	endpoint := "inproc://testReadWriterSmallBuf"

	pushSock, err := NewPush(endpoint)
	if err != nil {
		t.Errorf("NewPush failed: %s", err)
	}
	defer pushSock.Destroy()

	pullSock, err := NewPull(endpoint)
	if err != nil {
		t.Errorf("NewPull failed: %s", err)
	}

	pullReadWriter, err := NewReadWriter(pullSock)
	if err != nil {
		t.Errorf("NewReadWriter failed: %s", err)
	}

	defer pullReadWriter.Destroy()

	err = pushSock.SendFrame([]byte("Hello"), FlagNone)
	if err != nil {
		t.Errorf("pushSock.SendFrame failed: %s", err)
	}

	b := make([]byte, 3)

	n, err := pullReadWriter.Read(b)
	if n != 3 {
		t.Errorf("pullReadWriter.Read expected 3 bytes read %d", n)
	}

	if err != nil {
		t.Errorf("pullReadWriter.Read: %s", err)
	}

	if bytes.Compare(b, []byte("Hel")) != 0 {
		t.Errorf("expected 'Hel' received '%s'", b)
	}

	n, err = pullReadWriter.Read(b)
	if n != 2 {
		t.Errorf("pullReadWriter.Read expected 3 bytes read %d", n)
	}

	if bytes.Compare(b[:n], []byte("lo")) != 0 {
		t.Errorf("expected 'lo' received '%s'", b)
	}
}

func TestReadWriterWithRouterDealer(t *testing.T) {
	endpoint := "inproc://testReaderWithRouterDealer"

	dealerSock, err := NewDealer(endpoint)
	if err != nil {
		t.Errorf("NewDealer failure: %s", err)
	}
	defer dealerSock.Destroy()

	routerSock, err := NewRouter(endpoint)
	if err != nil {
		t.Errorf("NewDealer failure: %s", err)
	}
	defer routerSock.Destroy()

	routerReadWriter, err := NewReadWriter(routerSock)
	if err != nil {
		t.Errorf("NewReadWriter failed: %s", err)
	}
	defer routerReadWriter.Destroy()

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

func TestReadWriterWithRouterDealerAsync(t *testing.T) {
	endpoint := "inproc://testReadWriterWithRouterDealerAsync"

	dealerSock1, err := NewDealer(endpoint)
	if err != nil {
		t.Errorf("NewDealer failed: %s", err)
	}
	defer dealerSock1.Destroy()

	err = dealerSock1.Connect("inproc://test-read-router-async")
	if err != nil {
		t.Errorf("reqSock.Connect failed: %s", err)
	}

	dealerSock2, err := NewDealer(endpoint)
	if err != nil {
		t.Errorf("NewDealer failed: %s", err)
	}
	defer dealerSock2.Destroy()

	err = dealerSock2.Connect("inproc://test-read-router-async")
	if err != nil {
		t.Errorf("reqSock.Connect failed: %s", err)
	}

	routerSock, err := NewRouter(endpoint)
	if err != nil {
		t.Errorf("NewRouter failed: %s", err)
	}

	routerReadWriter, err := NewReadWriter(routerSock)
	if err != nil {
		t.Errorf("NewReadWriter failed: %s", err)
	}
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

func benchmarkReadWriter(size int, b *testing.B) {
	endpoint := fmt.Sprintf("inproc://benchSockReadWriter%d", size)

	pullSock, err := NewPull(endpoint)
	if err != nil {
		panic(err)
	}
	pullReader, err := NewReadWriter(pullSock)
	if err != nil {
		panic(err)
	}
	defer pullReader.Destroy()

	go func() {
		pushSock, err := NewPush(endpoint)
		if err != nil {
			panic(err)
		}
		pushWriter, err := NewReadWriter(pushSock)
		if err != nil {
			panic(err)
		}
		defer pushWriter.Destroy()

		payload := make([]byte, size)
		for i := 0; i < b.N; i++ {
			_, err = pushWriter.Write(payload)
			if err != nil {
				panic(err)
			}
		}
	}()

	payload := make([]byte, size)
	for i := 0; i < b.N; i++ {
		n, err := pullReader.Read(payload)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n != size {
			panic("msg too small")
		}
		b.SetBytes(int64(size))
	}
}

func BenchmarkReadWriter1k(b *testing.B)  { benchmarkSockReadWriter(1024, b) }
func BenchmarkReadWriter4k(b *testing.B)  { benchmarkSockReadWriter(4096, b) }
func BenchmarkReadWriter16k(b *testing.B) { benchmarkSockReadWriter(16384, b) }
