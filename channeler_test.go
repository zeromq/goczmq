package goczmq

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPushPullChanneler(t *testing.T) {
	push := NewPushChanneler("inproc://channelerpushpull")
	defer push.Destroy()

	pull := NewPullChanneler("inproc://channelerpushpull")
	defer pull.Destroy()

	push.SendChan <- [][]byte{[]byte("hello")}
	resp := <-pull.RecvChan
	if want, got := "hello", string(resp[0]); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	push.SendChan <- [][]byte{[]byte("world")}
	resp = <-pull.RecvChan
	if want, got := "world", string(resp[0]); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestDealerRouterChanneler(t *testing.T) {
	dealer := NewDealerChanneler("inproc://channelerdealerrouter")
	defer dealer.Destroy()

	router := NewRouterChanneler("inproc://channelerdealerrouter")
	defer router.Destroy()

	dealer.SendChan <- [][]byte{[]byte("hello")}
	resp := <-router.RecvChan
	if want, got := "hello", string(resp[1]); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	resp[1] = []byte("world")
	router.SendChan <- resp

	resp = <-dealer.RecvChan
	if want, got := "world", string(resp[0]); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestDealerChannelerRecvChanIsClosedOnDestroy(t *testing.T) {
	test_router := NewRouterChanneler("inproc://channelerouter")

	done := make(chan bool, 1)
	go func(router *Channeler) {
		resp := <-router.RecvChan
		if resp != nil {
			t.Errorf("expected nil response")
		}
		done <- true
	}(test_router)

	go func(router *Channeler) {
		time.Sleep(100 * time.Millisecond)
			router.Destroy()
	}(test_router)

	select {
	case <-done:
		break
	case <-time.After(1 * time.Second):
		t.Errorf("Router channel is not closed")
	}

}

func ExampleChanneler_output() {
	// create a dealer channeler
	dealer := NewDealerChanneler("inproc://channelerdealerrouter")
	defer dealer.Destroy()

	// create a router channeler
	router := NewRouterChanneler("inproc://channelerdealerrouter")
	defer router.Destroy()

	// send a hello message
	dealer.SendChan <- [][]byte{[]byte("Hello")}

	// receive the hello message
	request := <-router.RecvChan

	// first frame is identity of client - let's append 'World'
	// to the message and route it back
	request = append(request, []byte("World"))

	// send the reply
	router.SendChan <- request

	// receive the reply
	reply := <-dealer.RecvChan

	fmt.Printf("%s %s", string(reply[0]), string(reply[1]))
	// Output: Hello World
}

func BenchmarkChanneler(b *testing.B) {
	r := rand.Int63()
	pull := NewPullChanneler(fmt.Sprintf("inproc://benchchanneler-%d", r))
	defer pull.Destroy()

	go func() {
		push := NewPushChanneler(fmt.Sprintf("inproc://benchchanneler-%d", r))
		defer push.Destroy()

		payload := make([]byte, 1024)

		for i := 0; i < b.N; i++ {
			push.SendChan <- [][]byte{payload}
		}
	}()

	for i := 0; i < b.N; i++ {
		msg := <-pull.RecvChan
		if len(msg[0]) != 1024 {
			panic("message is corrupt")
		}
		b.SetBytes(1024)
	}
}
