package goczmq

import (
	"fmt"
	"testing"
)

func assertEvent(t *testing.T, monitor *Monitor, expectedEvent string) {
	poller, _ := NewPoller(monitor.Socket())
	defer poller.Destroy()

	socket, _ := poller.Wait(100)

	if socket == nil {
		t.Error("No messages received on monitor socket for 1 second")
		return
	}

	msg, _ := socket.RecvMessage()

	if len(msg) != 3 {
		t.Errorf("Expected message with 3 frames, got %v", len(msg))
	}

	eventName := string(msg[0])

	if eventName != expectedEvent {
		t.Errorf("Expected %v event, got %v", expectedEvent, eventName)
		return
	}
}

func TestMonitor(t *testing.T) {
	client := NewSock(Dealer)
	defer client.Destroy()

	clientmon := NewMonitor(client)
	defer clientmon.Destroy()

	clientmon.Verbose()

	clientmon.Listen("LISTENING")
	clientmon.Listen("ACCEPTED")
	clientmon.Start()

	server := NewSock(Dealer)
	defer server.Destroy()

	servermon := NewMonitor(server)
	defer servermon.Destroy()

	servermon.Listen("CONNECTED")
	servermon.Listen("DISCONNECTED")
	servermon.Start()

	port, _ := client.Bind("tcp://127.0.0.1:*")
	assertEvent(t, clientmon, "LISTENING")

	server.Connect(fmt.Sprint("tcp://127.0.0.1:", port))
	assertEvent(t, servermon, "CONNECTED")

	assertEvent(t, clientmon, "ACCEPTED")
}
