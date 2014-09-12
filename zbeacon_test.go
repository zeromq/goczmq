package goczmq

import (
	"testing"
)

func TestZbeacon(t *testing.T) {
	// Create a Zbeacon
	speaker := NewZbeacon()

	err := speaker.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	_, err = speaker.Configure(9999)
	if err != nil {
		t.Errorf("CONFIGURE error: %s", err)
	}

	listener := NewZbeacon()
	err = listener.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	_, err = listener.Configure(9999)
	if err != nil {
		t.Errorf("CONFIGURE error: %s", err)
	}

	speaker.Publish("HI", 100)
	listener.Destroy()
	speaker.Destroy()
}
