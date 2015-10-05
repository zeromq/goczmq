package goczmq

import "testing"

func TestGossip(t *testing.T) {

	server1 := NewGossip("server1")
	defer server1.Destroy()

	var err error

	if testing.Verbose() {
		err = server1.Verbose()
		if err != nil {
			t.Error(err)
		}
	}

	err = server1.Bind("inproc://server1")
	if err != nil {
		t.Error(err)
	}

	server2 := NewGossip("server2")
	defer server2.Destroy()

	if testing.Verbose() {
		err = server2.Verbose()
		if err != nil {
			t.Error(err)
		}
	}

	err = server2.Bind("inproc://server2")
	if err != nil {
		t.Error(err)
	}

	err = server2.Connect("inproc://server1")
	if err != nil {
		t.Error(err)
	}

	client1 := NewGossip("client1")
	defer client1.Destroy()

	if testing.Verbose() {
		err = client1.Verbose()
		if err != nil {
			t.Error(err)
		}
	}

	err = client1.Bind("inproc://client1")
	if err != nil {
		t.Error(err)
	}

	err = client1.Publish("client1-00", "0000")
	if err != nil {
		t.Error(err)
	}

	err = client1.Publish("client1-11", "0000")
	if err != nil {
		t.Error(err)
	}

	err = client1.Publish("client1-22", "0000")
	if err != nil {
		t.Error(err)
	}

	err = client1.Connect("inproc://server1")
	if err != nil {
		t.Error(err)
	}

	client2 := NewGossip("client2")
	defer client2.Destroy()

	if testing.Verbose() {
		err = client2.Verbose()
		if err != nil {
			t.Error(err)
		}
	}

	err = client2.Bind("inproc://client2")
	if err != nil {
		t.Error(err)
	}

	err = client2.Publish("client2-00", "0000")
	if err != nil {
		t.Error(err)
	}

	err = client2.Publish("client2-11", "0000")
	if err != nil {
		t.Error(err)
	}

	err = client2.Publish("client2-22", "0000")
	if err != nil {
		t.Error(err)
	}

	err = client2.Connect("inproc://server1")
	if err != nil {
		t.Error(err)
	}

	client3 := NewGossip("client3")
	defer client3.Destroy()

	if testing.Verbose() {
		err = client3.Verbose()
		if err != nil {
			t.Error(err)
		}
	}

	err = client3.Connect("inproc://server2")
	if err != nil {
		t.Error(err)
	}

	client4 := NewGossip("client4")
	defer client4.Destroy()

	if testing.Verbose() {
		err = client4.Verbose()
		if err != nil {
			t.Error(err)
		}
	}

	err = client4.Connect("inproc://server2")
	if err != nil {
		t.Error(err)
	}
}

func ExampleGossip() {
	gossiper := NewGossip("client")

	err := gossiper.Bind("inproc://gossip_example")
	if err != nil {
		panic(err)
	}

	err = gossiper.Publish("key", "value")
	if err != nil {
		panic(err)
	}

	gossiper.Destroy()
}
