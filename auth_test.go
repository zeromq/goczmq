package goczmq

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"testing"
)

// Here we start an auth actor, and set it to allow
// connections from localhost.  We then connect from
// a local client and verify we can send a message.

func TestAuthIPAllow(t *testing.T) {
	// start an auth actor and set it to VERBOSE mode
	auth := NewAuth()
	defer auth.Destroy()

	err := auth.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	// set the auth actor to allow connections from localhost
	err = auth.Allow("127.0.0.1")
	if err != nil {
		t.Errorf("ALLOW error: %s", err)
	}

	// create a pull socket server
	server := NewSock(PULL)
	server.SetZapDomain("global")
	defer server.Destroy()

	// bind the socket and get the port it bound to
	port, err := server.Bind("tcp://127.0.0.1:*")
	if port <= 0 {
		t.Errorf("port should be > 0, is %d", port)
	}

	// create a push socket client
	client := NewSock(PUSH)
	defer client.Destroy()

	// connect the client to the server socket
	err = client.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Errorf("client connect error: %s", err)
	}

	// send a hello world message
	client.SendString("Hello, World", 0)

	// create a poller and add the server socket to it
	poller, err := NewPoller(server)
	if err != nil {
		t.Errorf("NewPoller failed: %s", err)
	}
	defer poller.Destroy()

	// poll the server socket. we should have a message waiting.
	s := poller.Wait(200)
	if s == nil {
		t.Error("should be message waiting and there is none")
	}

	// receive the message and check the contents
	msg, err := s.RecvString()
	if err != nil {
		t.Error(err)
	}

	if msg != "Hello, World" {
		t.Error("message not sent properly")
	}
}

// Here we create an auth actor and tell the server to use
// "PLAIN" auth (username / password).  We will use a
// password file, and test that it works.

func TestAuthPlain(t *testing.T) {
	// Create a password file, and create one account in it
	// using username "admin" and password "Password".
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

	// start an auth actor and set it to VERBOSE mode
	auth := NewAuth()
	defer auth.Destroy()

	err = auth.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	// set the auth actor to allow connections from localhost
	err = auth.Allow("127.0.0.1")
	if err != nil {
		t.Errorf("ALLOW error: %s", err)
	}

	// tell the auth actor to use PLAIN authentication and
	// use the password file.
	err = auth.Plain("./password_test.txt")
	if err != nil {
		t.Errorf("PLAIN error: %s", err)
	}

	// create a pull socket server and set it to plain authentication
	server := NewSock(PULL)
	defer server.Destroy()
	server.SetZapDomain("global")
	server.SetPlainServer(1)

	// bind the socket and get the port it is bound to
	port, err := server.Bind("tcp://127.0.0.1:*")
	if port <= 0 {
		t.Errorf("port should be > 0, is %d", port)
	}

	// create a push client that will use the correct password
	goodClient := NewSock(PUSH)
	defer goodClient.Destroy()
	goodClient.SetPlainUsername("admin")
	goodClient.SetPlainPassword("Password")

	// create a push client that will use a bad password
	badClient := NewSock(PUSH)
	defer badClient.Destroy()
	badClient.SetPlainUsername("admin")
	badClient.SetPlainPassword("BadPassword")

	// connect to the server
	err = goodClient.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Errorf("goodClient connect error: %s", err)
	}

	// connect to the server as the good client, and send a message.
	// then poll the server to verify the message arrived, and
	// receive it.
	goodClient.SendString("Hello, World", 0)

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

	// connect to the server as the bad client, and send a message.
	// then poll the server to verify the message did not arrive.
	err = badClient.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Errorf("badClient connect error: %s", err)
	}

	// try to send a message.  this should succeed.
	badClient.SendString("Hello, World", 0)

	// poll for a message.  there should be one.
	s = poller.Wait(200)
	if s != nil {
		t.Error("poller should not have waiting message!")
	}
}

func TestAuthCurveAllow(t *testing.T) {
	// create auth service
	auth := NewAuth()
	defer auth.Destroy()

	err := auth.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	// create server socket and a server cert pair,
	// and apply the cert to the server
	server := NewSock(PULL)
	defer server.Destroy()
	server.SetZapDomain("global")
	serverCert := NewCert()
	serverKey := serverCert.PublicText()
	serverCert.Apply(server)
	server.SetCurveServer(1)

	// create a client + cert, attach
	// the cert to the client and set the
	// clients server key
	goodClient := NewSock(PUSH)
	defer goodClient.Destroy()
	goodClientCert := NewCert()
	goodClientCert.Apply(goodClient)
	goodClient.SetCurveServerkey(serverKey)

	// create a client, and don't assign a
	// cert or server key. this client should
	// be rejected.
	badClient := NewSock(PUSH)
	defer badClient.Destroy()

	// allow any cert
	auth.Curve(CURVE_ALLOW_ANY)

	// bind the server
	port, err := server.Bind("tcp://127.0.0.1:*")
	if port <= 0 {
		t.Errorf("port should be > 0, is %d", port)
	}

	// connect the goodClient
	err = goodClient.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Errorf("goodClient connect error: %s", err)
	}

	// connect the bad client
	err = badClient.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Errorf("client connect error: %s", err)
	}

	// try to send a message from the good client
	goodClient.SendString("Hello, World", 0)

	// see if we got a message
	poller, err := NewPoller(server)
	if err != nil {
		t.Errorf("NwPoller failed: %s", err)
	}
	defer poller.Destroy()

	s := poller.Wait(2000)
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

	// try to send a message from the bad client
	badClient.SendString("Hello, Bad World", 0)

	// poll and verify there is no waiting message from
	// the bad client.
	s = poller.Wait(200)
	if s != nil {
		t.Error("bad client should not have been able to send message")
	}
}

func TestAuthCurveCertificate(t *testing.T) {
	// create certificate directory
	testpath := path.Join("testauth")
	err := os.Mkdir(testpath, 0777)
	if err != nil {
		t.Fatal("TestAuthCurveCertificate could not create test dir")
	}

	// create auth service
	auth := NewAuth()
	defer auth.Destroy()

	// set verbus
	err = auth.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	// create a server socket and server cert pair,
	// get the public key, and apply to the cert
	// to the server socket.
	server := NewSock(PULL)
	defer server.Destroy()
	server.SetZapDomain("global")
	serverCert := NewCert()
	serverKey := serverCert.PublicText()
	serverCert.Apply(server)
	server.SetCurveServer(1)

	// create a client push socket, create a cert
	// for it and apply it, and add the server
	// public key to the client.
	goodClient := NewSock(PUSH)
	defer goodClient.Destroy()
	goodClientCert := NewCert()
	defer goodClientCert.Destroy()
	goodClientCert.Apply(goodClient)
	goodClient.SetCurveServerkey(serverKey)

	// save the good client public cert
	certfile := path.Join("testauth", "goodClient.txt")
	goodClientCert.SavePublic(certfile)

	// create a client push socket, and a cert for it.
	// this cert will not be added to the auth list.
	badClient := NewSock(PUSH)
	defer badClient.Destroy()
	badClientCert := NewCert()
	defer badClientCert.Destroy()
	badClientCert.Apply(badClient)
	badClient.SetCurveServerkey(serverKey)

	// set auth to only allow public keys from the
	// cert directory
	err = auth.Curve(testpath)
	if err != nil {
		panic(err)
	}

	// bind the server
	port, err := server.Bind("tcp://127.0.0.1:*")
	if port <= 0 {
		t.Errorf("port should be > 0, is %d", port)
	}

	// connect the good client
	err = goodClient.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Errorf("client connect error: %s", err)
	}

	// connect the bad client
	err = badClient.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Errorf("client connect error: %s", err)
	}

	// try to send a message from the good client
	goodClient.SendString("Hello, Good World", 0)

	// create a poller to poll the server, and verify
	// the message from the good client was received.
	poller, err := NewPoller(server)
	if err != nil {
		t.Errorf("NewPoller failed: %s", err)
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

	if msg != "Hello, Good World" {
		t.Error("message not sent properly")
	}

	// try to send a message from the bad client
	badClient.SendString("Hello, Bad World", 0)

	// poll and verify there is no waiting message from
	// the bad client.
	s = poller.Wait(200)
	if s != nil {
		t.Error("bad client should not have been able to send message")
	}

	os.RemoveAll(testpath)
}
