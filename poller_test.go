package goczmq

import (
	"fmt"
	"testing"
)

func TestPollerNewNoSocks(t *testing.T) {
	poller, err := NewPoller()
	if err != nil {
		t.Error(err)
	}
	defer poller.Destroy()

	pullSock1, err := NewPull("inproc://poller_new_no_socks")
	if err != nil {
		t.Error(err)
	}
	defer pullSock1.Destroy()

	err = poller.Add(pullSock1)
	if err != nil {
		t.Error(err)
	}

	pushSock, err := NewPush("inproc://poller_new_no_socks")
	if err != nil {
		t.Error(err)
	}
	defer pushSock.Destroy()

	err = pushSock.SendFrame([]byte("Hello"), FlagNone)
	if err != nil {
		t.Error(err)
	}

	s, err := poller.Wait(0)
	if err != nil {
		t.Error(err)
	}
	if want, have := pullSock1, s; want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	frame, _, err := s.RecvFrame()
	if err != nil {
		t.Error(err)
	}

	if want, have := "Hello", string(frame); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}
}

func TestPoller(t *testing.T) {
	pullSock1, err := NewPull("inproc://poller_pull1")
	if err != nil {
		t.Error(err)
	}
	defer pullSock1.Destroy()

	poller, err := NewPoller(pullSock1)
	if err != nil {
		t.Error(err)
	}
	defer poller.Destroy()

	if want, have := 1, len(poller.socks); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	pullSock2, err := NewPull("inproc://poller_pull2")
	if err != nil {
		t.Error(err)
	}
	defer pullSock2.Destroy()

	err = poller.Add(pullSock2)
	if err != nil {
		t.Error(err)
	}

	if want, have := 2, len(poller.socks); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	poller.Destroy()

	poller, err = NewPoller(pullSock1, pullSock2)
	if err != nil {
		t.Error(err)
	}

	if want, have := 2, len(poller.socks); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if want, have := pullSock1.zsockT, poller.socks[0].zsockT; want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if want, have := pullSock2.zsockT, poller.socks[1].zsockT; want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	pushSock, err := NewPush("inproc://poller_pull1")
	if err != nil {
		t.Error(err)
	}
	defer pushSock.Destroy()

	err = pushSock.SendFrame([]byte("Hello"), FlagNone)
	if err != nil {
		t.Error(err)
	}

	s, err := poller.Wait(0)
	if err != nil {
		t.Error(err)
	}
	if want, have := pullSock1, s; want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	frame, _, err := s.RecvFrame()
	if err != nil {
		t.Error(err)
	}

	if want, have := "Hello", string(frame); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	pushSock2, err := NewPush("inproc://poller_pull2")
	if err != nil {
		t.Error(err)
	}
	defer pushSock2.Destroy()

	err = pushSock2.SendFrame([]byte("World"), FlagNone)
	if err != nil {
		t.Error(err)
	}

	s, err = poller.Wait(0)
	if err != nil {
		t.Error(err)
	}
	if want, have := pullSock2, s; want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	frame, _, err = s.RecvFrame()
	if err != nil {
		t.Error(err)
	}

	if want, have := "World", string(frame); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	poller.Remove(pullSock2)
	if want, have := 1, len(poller.socks); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}
}

func TestPollerAfterDestroy(t *testing.T) {
	pullSock, err := NewPull("inproc://poller_pull")
	if err != nil {
		t.Error(err)
	}
	defer pullSock.Destroy()

	poller, err := NewPoller(pullSock)
	if err != nil {
		t.Error(err)
	}
	_, err = poller.Wait(0)
	if err != nil {
		t.Error(err)
	}

	poller.Destroy()
	_, err = poller.Wait(0)
	if err != ErrWaitAfterDestroy {
		t.Errorf("want %#v, have %#v", ErrWaitAfterDestroy, err)
	}
}

func ExamplePoller() {
	sock1, err := NewRouter("inproc://poller_example_1")
	if err != nil {
		panic(err)
	}
	defer sock1.Destroy()

	poller, err := NewPoller(sock1)
	if err != nil {
		panic(err)
	}

	sock2, err := NewRouter("inproc://poller_example_2")
	if err != nil {
		panic(err)
	}
	defer sock2.Destroy()

	err = poller.Add(sock2)
	if err != nil {
		panic(err)
	}

	// Poller.Wait(millis) returns first socket that has a waiting message
	poller.Wait(1)
}

func benchmarkPollerSendFrame(size int, b *testing.B) {
	pullSock := NewSock(Pull)
	defer pullSock.Destroy()

	_, err := pullSock.Bind(fmt.Sprintf("inproc://benchSockPoller%d", size))
	if err != nil {
		panic(err)
	}

	go func() {
		pushSock := NewSock(Push)
		defer pushSock.Destroy()
		err := pushSock.Connect(fmt.Sprintf("inproc://benchSockPoller%d", size))
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

	poller, err := NewPoller(pullSock)
	if err != nil {
		panic(err)
	}
	defer poller.Destroy()

	for i := 0; i < b.N; i++ {
		s, err := poller.Wait(-1)
		if err != nil {
			b.Error(err)
		}
		msg, _, err := s.RecvFrame()
		if err != nil {
			panic(err)
		}
		if len(msg) != size {
			panic("msg too small")
		}
		b.SetBytes(int64(size))
	}
}

func BenchmarkPollerSendFrame1k(b *testing.B)  { benchmarkPollerSendFrame(1024, b) }
func BenchmarkPollerSendFrame4k(b *testing.B)  { benchmarkPollerSendFrame(4096, b) }
func BenchmarkPollerSendFrame16k(b *testing.B) { benchmarkPollerSendFrame(16384, b) }
