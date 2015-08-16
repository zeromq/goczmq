package goczmq

import (
	"fmt"
	"testing"
)

func TestBeacon(t *testing.T) {
	// Create a Beacon
	speaker := NewBeacon()

	var err error

	if testing.Verbose() {
		err := speaker.Verbose()
		if err != nil {
			t.Errorf("VERBOSE error: %s", err)
		}
	}

	_, err = speaker.Configure(9999)
	if err != nil {
		t.Errorf("CONFIGURE error: %s", err)
	}

	listener := NewBeacon()

	if testing.Verbose() {
		err = listener.Verbose()
		if err != nil {
			t.Errorf("VERBOSE error: %s", err)
		}
	}

	_, err = listener.Configure(9999)
	if err != nil {
		t.Errorf("CONFIGURE error: %s", err)
	}

	err = listener.Subscribe("HI")
	if err != nil {
		t.Errorf("SubSCRIBE error: %s", err)
	}

	speaker.Publish("HI", 100)

	msg := listener.Recv(500)
	t.Logf("Address: %s", string(msg[0]))
	t.Logf("Beacon: %s", string(msg[1]))

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
