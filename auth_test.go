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
	server.SetOption(SockSetZapDomain("global"))
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

	err = client.SendFrame([]byte("Hello"), 1)
	if err != nil {
		t.Error(err)
	}

	err = client.SendFrame([]byte("World"), 0)
	if err != nil {
		t.Error(err)
	}

	poller, err := NewPoller(server)
	if err != nil {
		t.Error(err)
	}
	defer poller.Destroy()

	s, err := poller.Wait(200)
	if err != nil {
		t.Error(err)
	}
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
		err = os.Remove("./password_test.txt")
		if err != nil {
			t.Error(err)
		}
	}()

	w := bufio.NewWriter(pwfile)
	_, err = w.Write([]byte("admin=Password\n"))
	if err != nil {
		t.Error(err)
	}

	err = w.Flush()
	if err != nil {
		t.Error(err)
	}

	err = pwfile.Close()
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

	err = auth.Allow("127.0.0.1")
	if err != nil {
		t.Error(err)
	}

	err = auth.Plain("./password_test.txt")
	if err != nil {
		t.Error(err)
	}

	server := NewSock(Pull, SockSetZapDomain("global"), SockSetPlainServer(1))
	defer server.Destroy()

	port, err := server.Bind("tcp://127.0.0.1:*")
	if err != nil {
		t.Error(err)
	}

	goodClient := NewSock(Push, SockSetPlainUsername("admin"), SockSetPlainPassword("Password"))
	defer goodClient.Destroy()

	badClient := NewSock(Push, SockSetPlainUsername("admin"), SockSetPlainPassword("BadPassword"))
	defer badClient.Destroy()

	err = goodClient.Connect(fmt.Sprintf("tcp://127.0.0.1:%d", port))
	if err != nil {
		t.Error(err)
	}

	err = goodClient.SendFrame([]byte("Hello"), 1)
	if err != nil {
		t.Error(err)
	}

	err = goodClient.SendFrame([]byte("World"), 0)
	if err != nil {
		t.Error(err)
	}

	poller, err := NewPoller(server)
	if err != nil {
		t.Error(err)
	}
	defer poller.Destroy()

	s, err := poller.Wait(200)
	if err != nil {
		t.Error(err)
	}
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

	err = badClient.SendFrame([]byte("Hello"), 1)
	if err != nil {
		t.Error(err)
	}

	err = badClient.SendFrame([]byte("World"), 0)
	if err != nil {
		t.Error(err)
	}

	s, err = poller.Wait(200)
	if err != nil {
		t.Error(err)
	}
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

	server := NewSock(Pull, SockSetZapDomain("global"))
	defer server.Destroy()
	serverCert := NewCert()
	serverKey := serverCert.PublicText()
	serverCert.Apply(server)
	server.SetOption(SockSetCurveServer(1))

	goodClient := NewSock(Push)
	defer goodClient.Destroy()
	goodClientCert := NewCert()
	goodClientCert.Apply(goodClient)
	goodClient.SetOption(SockSetCurveServerkey(serverKey))

	badClient := NewSock(Push)
	defer badClient.Destroy()

	err = auth.Curve(CurveAllowAny)
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

	err = goodClient.SendFrame([]byte("Hello"), 1)
	if err != nil {
		t.Error(err)
	}

	err = goodClient.SendFrame([]byte("World"), 0)
	if err != nil {
		t.Error(err)
	}

	poller, err := NewPoller(server)
	if err != nil {
		t.Error(err)
	}
	defer poller.Destroy()

	s, err := poller.Wait(2000)
	if err != nil {
		t.Error(err)
	}
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

	err = badClient.SendFrame([]byte("Hello"), 1)
	if err != nil {
		t.Error(err)
	}

	err = badClient.SendFrame([]byte("Bad World"), 0)
	if err != nil {
		t.Error(err)
	}

	s, err = poller.Wait(200)
	if err != nil {
		t.Error(err)
	}
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

	server := NewSock(Pull, SockSetZapDomain("global"))
	defer server.Destroy()
	serverCert := NewCert()
	serverKey := serverCert.PublicText()
	serverCert.Apply(server)
	server.SetOption(SockSetCurveServer(1))

	goodClient := NewSock(Push, SockSetCurveServerkey(serverKey))
	defer goodClient.Destroy()
	goodClientCert := NewCert()
	defer goodClientCert.Destroy()
	goodClientCert.Apply(goodClient)

	certfile := path.Join("testauth", "goodClient.txt")
	err = goodClientCert.SavePublic(certfile)
	if err != nil {
		t.Error(err)
	}

	badClient := NewSock(Push, SockSetCurveServerkey(serverKey))
	defer badClient.Destroy()
	badClientCert := NewCert()
	defer badClientCert.Destroy()
	badClientCert.Apply(badClient)

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

	err = goodClient.SendFrame([]byte("Hello, Good World!"), 0)
	if err != nil {
		t.Error(err)
	}

	poller, err := NewPoller(server)
	if err != nil {
		t.Error(err)
	}
	defer poller.Destroy()

	s, err := poller.Wait(200)
	if err != nil {
		t.Error(err)
	}
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

	err = badClient.SendFrame([]byte("Hello, Bad World"), 0)
	if err != nil {
		t.Error(err)
	}

	s, err = poller.Wait(200)
	if err != nil {
		t.Error(err)
	}
	if s != nil {
		t.Errorf("want '%#v', have '%#v", nil, s)
	}

	err = os.RemoveAll(testpath)
	if err != nil {
		t.Error(err)
	}
}

func ExampleAuth() {
	// create a server certificate
	serverCert := NewCert()
	defer serverCert.Destroy()

	// create a client certificate and save it
	clientCert := NewCert()
	defer clientCert.Destroy()
	err := clientCert.SavePublic("client_cert")
	if err != nil {
		panic(err)
	}

	defer func() {
		err := os.Remove("client_cert")
		if err != nil {
			if err != nil {
				panic(err)
			}
		}
	}()

	// create an auth service
	auth := NewAuth()
	defer auth.Destroy()

	// tell the auth service the client cert is allowed
	err = auth.Curve("client_cert")
	if err != nil {
		panic(err)
	}

	// create a server socket and set it to
	// use the "global" auth domain
	server := NewSock(Push, SockSetZapDomain("global"))
	defer server.Destroy()

	// set the server cert as the server cert
	// for the socket we created and set it
	// to be a curve server
	serverCert.Apply(server)
	server.SetOption(SockSetCurveServer(1))

	// bind our server to an endpoint
	_, err = server.Bind("tcp://*:9898")
	if err != nil {
		panic(err)
	}

	// create a client socket
	client := NewSock(Pull)
	defer client.Destroy()

	// assign the client cert we made to the client
	clientCert.Apply(client)

	// set the server cert as the server cert
	// for the client. for the client to be
	// allowed to connect, it needs to know
	// the servers public cert.
	client.SetOption(SockSetCurveServerkey(serverCert.PublicText()))

	// connect
	err = client.Connect("tcp://127.0.0.1:9898")
	if err != nil {
		panic(err)
	}
}
