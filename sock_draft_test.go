// +build draft

package goczmq

import (
	"fmt"
	"testing"
)

func TestScatterGather(t *testing.T) {
	bogusScatter, err := NewScatter("bogus://bogus")
	if err == nil {
		t.Error(err)
	}
	defer bogusScatter.Destroy()

	bogusGather, err := NewGather("bogus://bogus")
	if err == nil {
		t.Error(err)
	}
	defer bogusGather.Destroy()

	scatter, err := NewScatter("inproc://scatter1,inproc://scatter2")
	if err != nil {
		t.Error(err)
	}
	defer scatter.Destroy()

	gather, err := NewGather("inproc://scatter1,inproc://scatter2")
	if err != nil {
		t.Error(err)
	}
	defer gather.Destroy()

	err = scatter.SendFrame([]byte("Hello World"), FlagNone)
	if err != nil {
		t.Error(err)
	}

	frame, _, err := gather.RecvFrame()
	if err != nil {
		t.Error(err)
	}

	if want, have := "Hello World", string(frame); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}
}

func benchmarkScatterGather(size int, b *testing.B) {
	gatherSock := NewSock(Gather)
	defer gatherSock.Destroy()

	_, err := gatherSock.Bind(fmt.Sprintf("inproc://benchScatterGather%#v", size))
	if err != nil {
		panic(err)
	}

	scatterSock := NewSock(Scatter)
	defer scatterSock.Destroy()
	err = scatterSock.Connect(fmt.Sprintf("inproc://benchScatterGather%#v", size))
	if err != nil {
		panic(err)
	}

	go func() {

		payload := make([]byte, size)
		for i := 0; i < b.N; i++ {
			err = scatterSock.SendFrame(payload, FlagNone)
			if err != nil {
				panic(err)
			}
		}
	}()

	for i := 0; i < b.N; i++ {
		msg, _, err := gatherSock.RecvFrame()
		if err != nil {
			panic(err)
		}
		if len(msg) != size {
			panic("msg too small")
		}

		b.SetBytes(int64(size))
	}
}

func BenchmarkScatterGather1k(b *testing.B)  { benchmarkScatterGather(1024, b) }
func BenchmarkScatterGather4k(b *testing.B)  { benchmarkScatterGather(4096, b) }
func BenchmarkScatterGather16k(b *testing.B) { benchmarkScatterGather(16384, b) }

func TestClientServer(t *testing.T) {
	bogusClient, err := NewClient("bogus://bogus")
	if err == nil {
		t.Error(err)
	}
	defer bogusClient.Destroy()

	bogusServer, err := NewServer("bogus://bogus")
	if err == nil {
		t.Error(err)
	}
	defer bogusServer.Destroy()

	client, err := NewClient("inproc://server")
	if err != nil {
		t.Error(err)
	}
	defer client.Destroy()

	server, err := NewServer("inproc://server")
	if err != nil {
		t.Error(err)
	}
	defer server.Destroy()

	err = client.SendFrame([]byte("Hello World"), FlagNone)
	if err != nil {
		t.Error(err)
	}

	frame, routing_id, err := server.RecvServerFrame()
	if err != nil {
		t.Error(err)
	}
	t.Logf("routing_id %d", routing_id)

	if want, have := "Hello World", string(frame); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	err = server.SendServerFrame([]byte("Hi World"), routing_id)
	if err != nil {
		t.Error(err)
	}

	frame, _, err = client.RecvFrame()
	if err != nil {
		t.Error(err)
	}

	if want, have := "Hi World", string(frame); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}
}
