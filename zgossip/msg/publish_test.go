package msg

import (
	"testing"

	"github.com/taotetek/goczmq"
)

// Yay! Test function.
func TestPublish(t *testing.T) {

	// Create pair of sockets we can send through

	// Output socket
	output := goczmq.NewSock(goczmq.DEALER)
	defer output.Destroy()

	routingId := "Shout"
	output.SetIdentity(routingId)
	_, err := output.Bind("inproc://selftest-publish")
	if err != nil {
		t.Fatal(err)
	}
	defer output.Unbind("inproc://selftest-publish")

	// Input socket
	input := goczmq.NewSock(goczmq.ROUTER)
	defer input.Destroy()

	err = input.Connect("inproc://selftest-publish")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Disconnect("inproc://selftest-publish")

	// Create a Publish message and send it through the wire
	publish := NewPublish()

	publish.Key = "Life is short but Now lasts for ever"

	publish.Value = "Life is short but Now lasts for ever"

	publish.Ttl = 123

	err = publish.Send(output)
	if err != nil {
		t.Fatal(err)
	}
	transit, err := Recv(input)
	if err != nil {
		t.Fatal(err)
	}

	tr := transit.(*Publish)

	if tr.Key != "Life is short but Now lasts for ever" {
		t.Fatalf("expected %s, got %s", "Life is short but Now lasts for ever", tr.Key)
	}

	if tr.Value != "Life is short but Now lasts for ever" {
		t.Fatalf("expected %s, got %s", "Life is short but Now lasts for ever", tr.Value)
	}

	if tr.Ttl != 123 {
		t.Fatalf("expected %d, got %d", 123, tr.Ttl)
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
