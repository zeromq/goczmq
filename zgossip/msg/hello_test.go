package msg

import (
	"testing"

	"github.com/taotetek/goczmq"
)

// Yay! Test function.
func TestHello(t *testing.T) {

	// Create pair of sockets we can send through

	// Output socket
	output := goczmq.NewSock(goczmq.DEALER)
	defer output.Destroy()

	routingId := "Shout"
	output.SetIdentity(routingId)
	_, err := output.Bind("inproc://selftest-hello")
	if err != nil {
		t.Fatal(err)
	}
	defer output.Unbind("inproc://selftest-hello")

	// Input socket
	input := goczmq.NewSock(goczmq.ROUTER)
	defer input.Destroy()

	err = input.Connect("inproc://selftest-hello")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Disconnect("inproc://selftest-hello")

	// Create a Hello message and send it through the wire
	hello := NewHello()

	err = hello.Send(output)
	if err != nil {
		t.Fatal(err)
	}
	transit, err := Recv(input)
	if err != nil {
		t.Fatal(err)
	}

	tr := transit.(*Hello)

	err = tr.Send(input)
	if err != nil {
		t.Fatal(err)
	}

	transit, err = Recv(output)
	if err != nil {
		t.Fatal(err)
	}

	if routingId != string(tr.RoutingId()) {
		t.Fatalf("expected %s, got %s", routingId, string(tr.RoutingId()))
	}
}
