package goczmq

import (
	"testing"
)

func TestZbeacon(t *testing.T) {
	// Create a Zbeacon
	beacon := NewZbeacon()

	// Destroy the Zbeacon
	beacon.Destroy()
}
