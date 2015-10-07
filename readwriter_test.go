package goczmq

import (
	"fmt"
	"io"
	"testing"
)

func TestReadWriter(t *testing.T) {
	endpoint := "inproc://testReadWriter"

	pushSock, err := NewPush(endpoint)
	if err != nil {
		t.Error(err)
	}
	defer pushSock.Destroy()

	pullSock, err := NewPull(endpoint)
	if err != nil {
		t.Error(err)
	}

	pullReadWriter, err := NewReadWriter(pullSock)
	if err != nil {
		t.Error(err)
	}
	defer pullReadWriter.Destroy()

	err = pushSock.SendFrame([]byte("Hello"), FlagNone)
	if err != nil {
		t.Error(err)
	}

	b := make([]byte, 5)

	n, err := pullReadWriter.Read(b)
	if err != nil {
		t.Error(err)
	}

	if want, got := 5, n; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := "Hello", string(b); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	err = pushSock.SendFrame([]byte("Hello World"), FlagNone)
	if err != nil {
		t.Error(err)
	}

	b = make([]byte, 8)
	n, err = pullReadWriter.Read(b)
	if err != nil {
		t.Error(err)
	}

	if want, got := "Hello Wo", string(b); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	n, err = pullReadWriter.Read(b)
	if err != nil {
		t.Error(err)
	}

	if want, got := 3, n; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := "rld", string(b[:n]); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	pullReadWriter.SetTimeout(1)
	n, err = pullReadWriter.Read(b)
	if want, got := io.EOF, err; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := 0, n; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}
}

func TestReadWriterWithBufferSmallerThanFrame(t *testing.T) {
	endpoint := "inproc://testReadWriterSmallBuf"

	pushSock, err := NewPush(endpoint)
	if err != nil {
		t.Error(err)
	}
	defer pushSock.Destroy()

	pullSock, err := NewPull(endpoint)
	if err != nil {
		t.Error(err)
	}

	pullReadWriter, err := NewReadWriter(pullSock)
	if err != nil {
		t.Error(err)
	}

	defer pullReadWriter.Destroy()

	err = pushSock.SendFrame([]byte("Hello"), FlagNone)
	if err != nil {
		t.Error(err)
	}

	b := make([]byte, 3)

	n, err := pullReadWriter.Read(b)
	if err != nil {
		t.Error(err)
	}

	if want, got := 3, n; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := "Hel", string(b); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	n, err = pullReadWriter.Read(b)
	if err != nil {
		t.Error(err)
	}

	if want, got := 2, n; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := "lo", string(b[:n]); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestReadWriterDoesNotSupportMultiPart(t *testing.T) {
	endpoint := "inproc://testReadWriterDoesNotSupportMultiPart"

	pushSock, err := NewPush(endpoint)
	if err != nil {
		t.Error(err)
	}
	defer pushSock.Destroy()

	pullSock, err := NewPull(endpoint)
	if err != nil {
		t.Error(err)
	}

	pullReadWriter, err := NewReadWriter(pullSock)
	if err != nil {
		t.Error(err)
	}
	defer pullReadWriter.Destroy()

	err = pushSock.SendFrame([]byte("Hello"), FlagMore)
	if err != nil {
		t.Error(err)
	}

	err = pushSock.SendFrame([]byte("World"), FlagNone)
	if err != nil {
		t.Error(err)
	}

	b := make([]byte, 5)

	n, err := pullReadWriter.Read(b)

	if want, got := ErrMultiPartUnsupported, err; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := 0, n; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}
}

func ExampleReadWriter_output() {
	push, err := NewPush("inproc://exampleReadWriter")
	if err != nil {
		panic(err)
	}

	pushWriter, err := NewReadWriter(push)
	if err != nil {
		panic(err)
	}

	defer pushWriter.Destroy()

	pull, err := NewPull("inproc://exampleReadWriter")
	if err != nil {
		panic(err)
	}

	pullReader, err := NewReadWriter(pull)
	if err != nil {
		panic(err)
	}

	defer pushReader.Destroy()

	n, err := pushWriter.Write([]byte("Hello World"))
	if err != nil {
		panic(err)
	}

	if n != 11 {
		panic("should have sent 11 bytes")
	}

	p := make([]byte, 255)

	n, err = pullReader.Read(p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", p[:n])
	// Output: Hello World
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
