package goczmq

import (
	"testing"
)

func TestSndhwm(t *testing.T) {
	sock := NewZsock(DEALER)
	expected := 5000
	sock.SetSndhwm(expected)
	val := sock.Sndhwm()
	if val != expected {
		t.Errorf("Sndhwm returned %d should be '%d'", val, expected)
	}
}

func TestRcvhwm(t *testing.T) {
	sock := NewZsock(DEALER)
	expected := 5000
	sock.SetRcvhwm(expected)
	val := sock.Rcvhwm()
	if val != expected {
		t.Errorf("Recvhwm returned %d should be '%d'", val, expected)
	}

}

func TestIdentity(t *testing.T) {
	sock := NewZsock(DEALER)
	expected := "myidentity"
	sock.SetIdentity(expected)
	val := sock.Identity()
	if val != expected {
		t.Errorf("TestIdentity returned %s should be %s", val, expected)
	}
}
