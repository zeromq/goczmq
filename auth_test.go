package goczmq

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"testing"
)

func TestAuthIPAllow(t *testing.T) {
	auth := NewAuth()
	defer auth.Destroy()

	var err error

	if testing.Verbose() {
		err = auth.Verbose()
		if err != nil {
			t.Error(err)
		}
	}

	err = auth.Allow("127.0.0.1")
	if err != nil {
		t.Error(err)
	}

	server := NewSock(Pull)
	server.SetZapDomain("global")
	defer server.Destroy()

	port, err := server.Bind("tcp://127.0.0.1:*")
	if err != nil {
		t.Error(err)
	}

	client := NewSock(Push)
	defer client.Destroy()

	err = client.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Error(err)
	}

	client.SendFrame([]byte("Hello"), 1)
	client.SendFrame([]byte("World"), 0)

	poller, err := NewPoller(server)
	if err != nil {
		t.Error(err)
	}
	defer poller.Destroy()

	s := poller.Wait(200)
	if want, have := server, s; want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	msg, err := s.RecvMessage()
	if err != nil {
		t.Error(err)
	}

	if want, have := "Hello", string(msg[0]); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if want, have := "World", string(msg[1]); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}
}

func TestAuthPlain(t *testing.T) {
	pwfile, err := os.Create("./password_test.txt")
	if err != nil {
		t.Error(err)
	}

	defer func() {
		os.Remove("./password_test.txt")
	}()

	w := bufio.NewWriter(pwfile)
	w.Write([]byte("admin=Password\n"))
	w.Flush()
	pwfile.Close()

	auth := NewAuth()
	defer auth.Destroy()

	if testing.Verbose() {
		err = auth.Verbose()
		if err != nil {
			t.Error(err)
		}
	}

	err = auth.Allow("127.0.0.1")
	if err != nil {
		t.Error(err)
	}

	err = auth.Plain("./password_test.txt")
	if err != nil {
		t.Error(err)
	}

	server := NewSock(Pull)
	defer server.Destroy()
	server.SetZapDomain("global")
	server.SetPlainServer(1)

	port, err := server.Bind("tcp://127.0.0.1:*")
	if err != nil {
		t.Error(err)
	}

	goodClient := NewSock(Push)
	defer goodClient.Destroy()
	goodClient.SetPlainUsername("admin")
	goodClient.SetPlainPassword("Password")

	badClient := NewSock(Push)
	defer badClient.Destroy()
	badClient.SetPlainUsername("admin")
	badClient.SetPlainPassword("BadPassword")

	err = goodClient.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Error(err)
	}

	goodClient.SendFrame([]byte("Hello"), 1)
	goodClient.SendFrame([]byte("World"), 0)

	poller, err := NewPoller(server)
	if err != nil {
		t.Error(err)
	}
	defer poller.Destroy()

	s := poller.Wait(200)
	if want, have := server, s; want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	msg, err := s.RecvMessage()
	if err != nil {
		t.Error(err)
	}

	if want, have := "Hello", string(msg[0]); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if want, have := "World", string(msg[1]); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	err = badClient.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Error(err)
	}

	badClient.SendFrame([]byte("Hello"), 1)
	badClient.SendFrame([]byte("World"), 0)

	s = poller.Wait(200)
	if s != nil {
		t.Errorf("want %#v, have %#v", nil, s)
	}

	if want, have := "Hello", string(msg[0]); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if want, have := "World", string(msg[1]); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}
}

func TestAuthCurveAllowAny(t *testing.T) {
	auth := NewAuth()
	defer auth.Destroy()

	var err error

	if testing.Verbose() {
		err = auth.Verbose()
		if err != nil {
			t.Error(err)
		}
	}

	server := NewSock(Pull)
	defer server.Destroy()
	server.SetZapDomain("global")
	serverCert := NewCert()
	serverKey := serverCert.PublicText()
	serverCert.Apply(server)
	server.SetCurveServer(1)

	goodClient := NewSock(Push)
	defer goodClient.Destroy()
	goodClientCert := NewCert()
	goodClientCert.Apply(goodClient)
	goodClient.SetCurveServerkey(serverKey)

	badClient := NewSock(Push)
	defer badClient.Destroy()

	auth.Curve(CurveAllowAny)

	port, err := server.Bind("tcp://127.0.0.1:*")
	if err != nil {
		t.Error(err)
	}

	err = goodClient.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Error(err)
	}

	err = badClient.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Error(err)
	}

	goodClient.SendFrame([]byte("Hello"), 1)
	goodClient.SendFrame([]byte("World"), 0)

	poller, err := NewPoller(server)
	if err != nil {
		t.Error(err)
	}
	defer poller.Destroy()

	s := poller.Wait(2000)
	if want, have := server, s; want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	msg, err := s.RecvMessage()
	if err != nil {
		t.Error(err)
	}

	if want, have := "Hello", string(msg[0]); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if want, have := "World", string(msg[1]); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	badClient.SendFrame([]byte("Hello"), 1)
	badClient.SendFrame([]byte("Bad World"), 0)

	s = poller.Wait(200)
	if s != nil {
		t.Errorf("want %#v, have %#v", nil, s)
	}
}

func TestAuthCurveAllowCertificate(t *testing.T) {
	testpath := path.Join("testauth")
	err := os.Mkdir(testpath, 0777)
	if err != nil {
		t.Error(err)
	}

	auth := NewAuth()
	defer auth.Destroy()

	if testing.Verbose() {
		err = auth.Verbose()
		if err != nil {
			t.Error(err)
		}
	}

	server := NewSock(Pull)
	defer server.Destroy()
	server.SetZapDomain("global")
	serverCert := NewCert()
	serverKey := serverCert.PublicText()
	serverCert.Apply(server)
	server.SetCurveServer(1)

	goodClient := NewSock(Push)
	defer goodClient.Destroy()
	goodClientCert := NewCert()
	defer goodClientCert.Destroy()
	goodClientCert.Apply(goodClient)
	goodClient.SetCurveServerkey(serverKey)

	certfile := path.Join("testauth", "goodClient.txt")
	goodClientCert.SavePublic(certfile)

	badClient := NewSock(Push)
	defer badClient.Destroy()
	badClientCert := NewCert()
	defer badClientCert.Destroy()
	badClientCert.Apply(badClient)
	badClient.SetCurveServerkey(serverKey)

	err = auth.Curve(testpath)
	if err != nil {
		t.Error(err)
	}

	port, err := server.Bind("tcp://127.0.0.1:*")
	if err != nil {
		t.Error(err)
	}

	err = goodClient.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Error(err)
	}

	err = badClient.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Error(err)
	}

	goodClient.SendFrame([]byte("Hello, Good World!"), 0)

	poller, err := NewPoller(server)
	if err != nil {
		t.Error(err)
	}
	defer poller.Destroy()

	s := poller.Wait(200)
	if want, have := server, s; want != have {
		t.Errorf("want '%#v', have '%#v'", want, have)
	}

	msg, err := s.RecvMessage()
	if err != nil {
		t.Error(err)
	}

	if want, have := "Hello, Good World!", string(msg[0]); want != have {
		t.Errorf("want '%#v', have '%#v'", want, have)
	}

	badClient.SendFrame([]byte("Hello, Bad World"), 0)

	s = poller.Wait(200)
	if s != nil {
		t.Errorf("want '%#v', have '%#v", nil, s)
	}

	os.RemoveAll(testpath)
}

func ExampleAuth() {
	// create a server certificate
	serverCert := NewCert()
	defer serverCert.Destroy()

	// create a client certificate and save it
	clientCert := NewCert()
	defer clientCert.Destroy()
	clientCert.SavePublic("client_cert")
	defer func() { os.Remove("client_cert") }()

	// create an auth service
	auth := NewAuth()
	defer auth.Destroy()

	// tell the auth service the client cert is allowed
	auth.Curve("client_cert")

	// create a server socket and set it to
	// use the "global" auth domain
	server := NewSock(Push)
	defer server.Destroy()
	server.SetZapDomain("global")

	// set the server cert as the server cert
	// for the socket we created and set it
	// to be a curve server
	serverCert.Apply(server)
	server.SetCurveServer(1)

	// bind our server to an endpoint
	server.Bind("tcp://*:9898")

	// create a client socket
	client := NewSock(Pull)
	defer client.Destroy()

	// assign the client cert we made to the client
	clientCert.Apply(client)

	// set the server cert as the server cert
	// for the client. for the client to be
	// allowed to connect, it needs to know
	// the servers public cert.
	client.SetCurveServerkey(serverCert.PublicText())

	// connect
	client.Connect("tcp://127.0.0.1:9898")
}
