package goczmq

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestAuthNoDeny(t *testing.T) {
	server := NewSock(PULL)
	defer server.Destroy()

	server.SetZapDomain("global")

	client := NewSock(PUSH)
	defer client.Destroy()

	auth := NewAuth()
	defer auth.Destroy()

	err := auth.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	port, err := server.Bind("tcp://127.0.0.1:*")
	if port <= 0 {
		t.Error("port should be > 0, is %d", port)
	}

	err = client.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Errorf("client connect error: %s", err)
	}

	client.SendString("Hello, World", 0)

	poller, err := NewPoller(server)
	if err != nil {
		t.Errorf("NwPoller failed: %s", err)
	}

	defer poller.Destroy()

	s := poller.Wait(200)
	if s == nil {
		t.Error("should be message waiting and there is none")
	}

	msg, err := s.RecvString()
	if err != nil {
		t.Error(err)
	}

	if msg != "Hello, World" {
		t.Error("message not sent properly")
	}
}

func TestAuthPlainDeny(t *testing.T) {
	// create an auth service
	auth := NewAuth()
	defer auth.Destroy()

	// set verbose
	err := auth.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	// allow localhost connections
	err = auth.Allow("127.0.0.1")
	if err != nil {
		t.Errorf("ALLOW error: %s", err)
	}

	// create a server socket
	server := NewSock(PULL)
	defer server.Destroy()

	// set plain authentication
	server.SetPlainServer(1)

	// bind to an ephemeral port
	port, err := server.Bind("tcp://127.0.0.1:*")
	if port <= 0 {
		t.Error("port should be > 0, is %d", port)
	}

	// create a client
	client := NewSock(PUSH)
	defer client.Destroy()

	// set a username and password
	client.SetPlainUsername("admin")
	client.SetPlainPassword("bogus")

	// try to connect
	err = client.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Errorf("client connect error: %s", err)
	}

	// try to send a message.  since we have not told the auth
	// service where a password file is, this should fail
	client.SendString("Hello, World", 0)

	poller, err := NewPoller(server)
	if err != nil {
		t.Errorf("NwPoller failed: %s", err)
	}
	defer poller.Destroy()

	// poll for a message.  there should not be one.
	s := poller.Wait(200)
	if s != nil {
		t.Error("poller received message should not have!")
	}
}

func TestAuthPlainAllow(t *testing.T) {
	pwfile, err := os.Create("./password_test.txt")
	if err != nil {
		t.Fatalf("could not create password test file")
	}
	defer func() {
		os.Remove("./password_test.txt")
	}()

	w := bufio.NewWriter(pwfile)
	w.Write([]byte("admin=Password\n"))
	w.Flush()
	pwfile.Close()

	// create an auth service
	auth := NewAuth()
	defer auth.Destroy()

	// set verbose
	err = auth.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	// allow localhost connections
	err = auth.Allow("127.0.0.1")
	if err != nil {
		t.Errorf("ALLOW error: %s", err)
	}

	err = auth.Plain("./password_test.txt")
	if err != nil {
		t.Errorf("PLAIN error: %s", err)
	}

	// create a server socket
	server := NewSock(PULL)
	defer server.Destroy()

	// set plain authentication
	server.SetPlainServer(1)

	// bind to an ephemeral port
	port, err := server.Bind("tcp://127.0.0.1:*")
	if port <= 0 {
		t.Error("port should be > 0, is %d", port)
	}

	// create a client
	client := NewSock(PUSH)
	defer client.Destroy()

	// set a username and password
	client.SetPlainUsername("admin")
	client.SetPlainPassword("Password")

	// try to connect
	err = client.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Errorf("client connect error: %s", err)
	}

	// try to send a message.  since we have not told the auth
	// service where a password file is, this should fail
	client.SendString("Hello, World", 0)

	poller, err := NewPoller(server)
	if err != nil {
		t.Errorf("NwPoller failed: %s", err)
	}
	defer poller.Destroy()

	// poll for a message.  there should be one.
	s := poller.Wait(200)
	if s == nil {
		t.Error("poller should have waiting message!")
	}

	msg, err := s.RecvString()
	if err != nil {
		t.Error(err)
	}

	if msg != "Hello, World" {
		t.Error("message not sent properly")
	}

}

func TestAuthCurveAllow(t *testing.T) {
	// create server socket and set global auth domain
	server := NewSock(PULL)
	defer server.Destroy()
	server.SetZapDomain("global")

	// create client socket
	client := NewSock(PUSH)
	defer client.Destroy()

	// create auth service
	auth := NewAuth()
	defer auth.Destroy()

	// set verbus
	err := auth.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	// create server cert pair and get public key
	// and apply cert to server socket
	serverCert := NewCert()
	serverKey := serverCert.PublicText()
	serverCert.Apply(server)
	server.SetCurveServer(1)

	// create client cert
	clientCert := NewCert()
	clientCert.Apply(client)
	client.SetCurveServerkey(serverKey)

	// allow any cert
	auth.Curve(CURVE_ALLOW_ANY)

	// bind the server
	port, err := server.Bind("tcp://127.0.0.1:*")
	if port <= 0 {
		t.Error("port should be > 0, is %d", port)
	}

	// connect the client
	err = client.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Errorf("client connect error: %s", err)
	}

	// try to send a message
	client.SendString("Hello, World", 0)

	// see if we got a message
	poller, err := NewPoller(server)
	if err != nil {
		t.Errorf("NwPoller failed: %s", err)
	}

	defer poller.Destroy()

	s := poller.Wait(200)
	if s == nil {
		t.Error("should be message waiting and there is none")
	}

	msg, err := s.RecvString()
	if err != nil {
		t.Error(err)
	}

	if msg != "Hello, World" {
		t.Error("message not sent properly")
	}
}
