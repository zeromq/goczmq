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

func TestPubSubChanneler(t *testing.T) {
	pub := NewXPubChanneler("inproc://channelerpubsub")
	defer pub.Destroy()

	sub := NewSubChanneler("inproc://channelerpubsub", "a,b")
	defer sub.Destroy()

	confirmXPubSubscriptions(t, pub, 2)

	pub.SendChan <- [][]byte{[]byte("a"), []byte("message")}
	select {
	case resp := <-sub.RecvChan:
		topic, message := string(resp[0]), string(resp[1])
		if want, got := "a", topic; want != got {
			t.Errorf("want '%s', got '%s'", want, got)
		}
		if want, got := "message", message; want != got {
			t.Errorf("want '%s', got '%s'", want, got)
		}
	case <-time.After(time.Second * 2):
		t.Errorf("timeout")
	}

	pub.SendChan <- [][]byte{[]byte("X"), []byte("message")}
	pub.SendChan <- [][]byte{[]byte("b"), []byte("message")}
	select {
	case resp := <-sub.RecvChan:
		topic, message := string(resp[0]), string(resp[1])
		if want, got := "b", topic; want != got {
			t.Errorf("want '%s', got '%s'", want, got)
		}
		if want, got := "message", message; want != got {
			t.Errorf("want '%s', got '%s'", want, got)
		}
	case <-time.After(time.Second * 1):
		t.Errorf("timeout")
	}

	sub.Subscribe("c")
	sub.Unsubscribe("a")

	confirmXPubSubscriptions(t, pub, 2)

	pub.SendChan <- [][]byte{[]byte("a"), []byte("message")}
	pub.SendChan <- [][]byte{[]byte("c"), []byte("message")}
	select {
	case resp := <-sub.RecvChan:
		topic, message := string(resp[0]), string(resp[1])
		if want, got := "c", topic; want != got {
			t.Errorf("want '%s', got '%s'", want, got)
		}
		if want, got := "message", message; want != got {
			t.Errorf("want '%s', got '%s'", want, got)
		}
	case <-time.After(time.Second * 1):
		t.Errorf("timeout")
	}
}

func TestPubSubChannelerNoInitialSubscription(t *testing.T) {
	pub := NewXPubChanneler("inproc://channelerpubsub2")
	defer pub.Destroy()

	sub := NewSubChanneler("inproc://channelerpubsub2")
	defer sub.Destroy()

	sub.Subscribe("a")

	confirmXPubSubscriptions(t, pub, 1)

	pub.SendChan <- [][]byte{[]byte("a"), []byte("message")}
	select {
	case resp := <-sub.RecvChan:
		topic, message := string(resp[0]), string(resp[1])
		if want, got := "a", topic; want != got {
			t.Errorf("want '%s', got '%s'", want, got)
		}
		if want, got := "message", message; want != got {
			t.Errorf("want '%s', got '%s'", want, got)
		}
	case <-time.After(time.Second * 2):
		t.Errorf("timeout")
	}
}

func TestPubSubChannelerOptionError(t *testing.T) {
	sub := NewSubChanneler("inproc://channelerpubsub2", 32)
	defer sub.Destroy()
	err := <-sub.ErrChan
	expected := fmt.Errorf("Don't know how to handle a %T argument to NewSubChanneler", 32)
	assertEqual(t, expected.Error(), err.Error())
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

func TestDealerRouterChannelerAttachError(t *testing.T) {
	dealer := NewDealerChanneler("bad endpoint")
	defer dealer.Destroy()

	dealer.SendChan <- [][]byte{[]byte("hello")}

	select {
	case resp := <-dealer.RecvChan:
		if want, got := "world", string(resp[0]); want != got {
			t.Errorf("want '%s', got '%s'", want, got)
		}
	case err := <-dealer.ErrChan:
		assertEqual(t, ErrSockAttach, err)
	}
}

func TestDealerRouterChannelerEmptyEndpointsError(t *testing.T) {
	dealer := NewDealerChanneler("")
	defer dealer.Destroy()

	dealer.SendChan <- [][]byte{[]byte("hello")}

	select {
	case resp := <-dealer.RecvChan:
		if want, got := "world", string(resp[0]); want != got {
			t.Errorf("want '%s', got '%s'", want, got)
		}
	case err := <-dealer.ErrChan:
		assertEqual(t, ErrSockAttachEmptyEndpoints, err)
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

func confirmXPubSubscriptions(t *testing.T, pub *Channeler, count int) {
	for i := 0; i < count; i++ {
		select {
		case <-pub.RecvChan:
		case <-time.After(time.Second * 2):
			t.Errorf("timeout")
		}
	}
}
