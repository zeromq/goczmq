package goczmq

import (
	"testing"
)

func TestZproxy(t *testing.T) {
	proxy := NewZproxy()
	proxy.Destroy()
}
