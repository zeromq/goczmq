package goczmq

import "testing"

func TestPushPullChanneler(t *testing.T) {
	push := NewPushChanneler("inproc://channelerpushpull")
	defer push.Destroy()

	pull := NewPullChanneler("inproc://channelerpushpull")
	defer pull.Destroy()

	push.SendChan <- [][]byte{[]byte("hello")}
	resp := <-pull.RecvChan
	if string(resp[0]) != "hello" {
		t.Errorf("failed")
	}

	push.SendChan <- [][]byte{[]byte("world")}
	resp = <-pull.RecvChan
	if string(resp[0]) != "world" {
		t.Errorf("failed")
	}
}

func TestDealerRouterChanneler(t *testing.T) {
	dealer := NewDealerChanneler("inproc://channelerdealerrouter")
	defer dealer.Destroy()

	router := NewRouterChanneler("inproc://channelerdealerrouter")
	defer router.Destroy()

	dealer.SendChan <- [][]byte{[]byte("hello")}
	resp := <-router.RecvChan
	if string(resp[1]) != "hello" {
		t.Errorf("failed")
	}

	resp[1] = []byte("world")
	router.SendChan <- resp

	resp = <-dealer.RecvChan
	if string(resp[0]) != "world" {
		t.Errorf("failed")
	}
}

func BenchmarkChanneler(b *testing.B) {
	pull := NewPullChanneler("inproc://benchchanneler")
	defer pull.Destroy()

	go func() {
		push := NewPushChanneler("inproc://benchchanneler")
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
	}
}
