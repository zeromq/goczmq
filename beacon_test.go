package goczmq

import (
	"fmt"
	"testing"
)

// TestBeacon tests sending a UDP beacon to a listener
func TestBeacon(t *testing.T) {
	speaker := NewBeacon()

	var err error
	if testing.Verbose() {
		err := speaker.Verbose()
		if err != nil {
			t.Error(err)
		}
	}

	_, err = speaker.Configure(9999)
	if err != nil {
		t.Error(err)
	}

	listener := NewBeacon()

	if testing.Verbose() {
		err = listener.Verbose()
		if err != nil {
			t.Error(err)
		}
	}

	_, err = listener.Configure(9999)
	if err != nil {
		t.Error(err)
	}

	err = listener.Subscribe("HI")
	if err != nil {
		t.Error(err)
	}

	err = speaker.Publish("HI", 100)
	if err != nil {
		t.Error(err)
	}

	msg := listener.Recv(500)
	if len(msg) == 2 {
		t.Logf("Address: %s", string(msg[0]))
		t.Logf("Beacon: %s", string(msg[1]))
	}

	listener.Destroy()
	speaker.Destroy()
}

func ExampleBeacon() {
	beacon := NewBeacon()
	defer beacon.Destroy()

	address, err := beacon.Configure(9999)
	if err != nil {
		panic(err)
	}

	fmt.Printf("started beacon on: %s", address)

	err = beacon.Publish("HI", 100)
	if err != nil {
		panic(err)
	}
}
