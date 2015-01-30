package msg

import (
	"testing"

	"github.com/zeromq/goczmq"
)

// Yay! Test function.
func TestChunk(t *testing.T) {

	// Create pair of sockets we can send through

	// Output socket
	output := goczmq.NewSock(goczmq.DEALER)
	defer output.Destroy()

	routingId := "Shout"
	output.SetIdentity(routingId)
	_, err := output.Bind("inproc://selftest-chunk")
	if err != nil {
		t.Fatal(err)
	}
	defer output.Unbind("inproc://selftest-chunk")

	// Input socket
	input := goczmq.NewSock(goczmq.ROUTER)
	defer input.Destroy()

	err = input.Connect("inproc://selftest-chunk")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Disconnect("inproc://selftest-chunk")

	// Create a Chunk message and send it through the wire
	chunk := NewChunk()

	chunk.More = 123

	chunk.Payload = []byte("Captcha Diem")

	err = chunk.Send(output)
	if err != nil {
		t.Fatal(err)
	}
	transit, err := Recv(input)
	if err != nil {
		t.Fatal(err)
	}

	tr := transit.(*Chunk)

	if tr.More != 123 {
		t.Fatalf("expected %d, got %d", 123, tr.More)
	}

	if string(tr.Payload) != "Captcha Diem" {
		t.Fatalf("expected %s, got %s", "Captcha Diem", tr.Payload)
	}

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
