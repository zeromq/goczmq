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

	speaker.Publish("HI", 100)

	msg := listener.Recv(500)

	listener.Destroy()
	speaker.Destroy()
}

func ExampleBeacon(t *testing.T) {
	beacon := NewBeacon()
	defer beacon.Destroy()

	address, err := beacon.Configure(9999)
	if err != nil {
		panic(err)
	}
	fmt.Printf("started beacon on: %s", address)
	beacon.Publish("HI", 100)
}
