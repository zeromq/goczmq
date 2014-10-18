package goczmq

import (
	"testing"
)

func TestAuth(t *testing.T) {
	auth := NewAuth()

	err := auth.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	auth.Destroy()
}
