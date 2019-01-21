package goczmq

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, value interface{}, message ...string) {
	if expected == value {
		return
	}
	msg := fmt.Sprintf("'%v' != '%v'. Expected '%v' but got '%v'", expected, value, expected, value)
	for _, m := range message {
		msg += fmt.Sprintf(". %s", m)
	}
	t.Fatal(msg)
}
