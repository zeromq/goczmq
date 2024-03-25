package goczmq

import (
	"fmt"
	"github.com/nofeaturesonlybugs/z85"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProxy(t *testing.T) {
	// Create and configure our proxy
	proxy := NewProxy()
	defer proxy.Destroy()

	var err error

	if testing.Verbose() {
		err = proxy.Verbose()
		if err != nil {
			t.Error(err)
		}
	}

	err = proxy.SetFrontend(Pull, "inproc://frontend")
	if err != nil {
		t.Error(err)
	}

	err = proxy.SetBackend(Push, "inproc://backend")
	if err != nil {
		t.Error(err)
	}

	err = proxy.SetCapture("inproc://capture")
	if err != nil {
		t.Error(err)
	}

	// connect application sockets to proxy
	faucet := NewSock(Push)
	err = faucet.Connect("inproc://frontend")
	if err != nil {
		t.Error(err)
	}
	defer faucet.Destroy()

	sink := NewSock(Pull)
	err = sink.Connect("inproc://backend")
	if err != nil {
		t.Error(err)
	}
	defer sink.Destroy()

	tap := NewSock(Pull)
	_, err = tap.Bind("inproc://capture")
	if err != nil {
		t.Error(err)
	}
	defer tap.Destroy()

	// send some messages and check they arrived
	err = faucet.SendFrame([]byte("Hello"), FlagNone)
	if err != nil {
		t.Error(err)
	}

	err = faucet.SendFrame([]byte("World"), FlagNone)
	if err != nil {
		t.Error(err)
	}

	// check the tap
	b, f, err := tap.RecvFrame()
	if err != nil {
		t.Error(err)
	}

	if want, have := false, f == FlagMore; want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if want, have := "Hello", string(b); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	b, f, err = tap.RecvFrame()
	if err != nil {
		t.Error(err)
	}

	if want, have := false, f == FlagMore; want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if want, have := "World", string(b); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	b, f, err = sink.RecvFrame()
	if err != nil {
		t.Error(err)
	}

	if want, have := false, f == FlagMore; want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if want, have := "Hello", string(b); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	_, f, err = sink.RecvFrame()
	if err != nil {
		t.Error(err)
	}

	if f == FlagMore {
		t.Error("FlagMore set and should not be")
	}

	if want, have := false, f == FlagMore; want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	err = proxy.Pause()
	require.NoError(t, err)

	err = faucet.SendFrame([]byte("Belated Hello"), FlagNone)
	require.NoError(t, err)

	if want, have := false, sink.Pollin(); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if want, have := false, tap.Pollin(); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	err = proxy.Resume()
	require.NoError(t, err)

	b, f, err = sink.RecvFrame()
	require.NoError(t, err)

	if want, have := false, f == FlagMore; want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if want, have := "Belated Hello", string(b); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	b, f, err = tap.RecvFrame()
	require.NoError(t, err)

	if want, have := false, f == FlagMore; want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if want, have := "Belated Hello", string(b); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	proxy.Destroy()
}

func TestProxyCurve(t *testing.T) {
	serverPubKey := "j+KTl+V-G75#pBWwItQta7<5Rzs:N$1xFwjW2{C2"
	serverSecretKey := "*i(F-QJdIE04$AtHVoo.AwGjcM}0sN../[j)<)N}"
	serverPubKeyBinary, err := z85.PaddedDecode(serverPubKey)
	require.NoError(t, err)
	serverSecretKeyBinary, err := z85.PaddedDecode(serverSecretKey)
	require.NoError(t, err)

	serverCert, err := NewCertFromKeys(serverPubKeyBinary, serverSecretKeyBinary)
	require.NoError(t, err)
	clientCert := NewCert()

	// Create and configure our proxy
	proxy := NewProxy()
	defer proxy.Destroy()

	if testing.Verbose() {
		err = proxy.Verbose()
		require.NoError(t, err)
	}

	err = proxy.SetFrontendDomain("global")
	require.NoError(t, err)
	err = proxy.SetFrontendCurve(serverPubKey, serverSecretKey)
	require.NoError(t, err)
	err = proxy.SetFrontend(Pull, "inproc://curveFrontend")
	require.NoError(t, err)

	err = proxy.SetBackendDomain("global")
	require.NoError(t, err)
	err = proxy.SetBackendCurve(serverPubKey, serverSecretKey)
	require.NoError(t, err)
	err = proxy.SetBackend(Push, "inproc://curveBackend")
	require.NoError(t, err)

	// connect application sockets to proxy
	faucet := NewSock(Push)
	faucet.SetOption(SockSetCurveServerkey(serverCert.PublicText()))
	clientCert.Apply(faucet)
	err = faucet.Connect("inproc://curveFrontend")
	require.NoError(t, err)
	defer faucet.Destroy()

	sink := NewSock(Pull)
	sink.SetOption(SockSetCurveServerkey(serverCert.PublicText()))
	clientCert.Apply(sink)
	err = sink.Connect("inproc://curveBackend")
	require.NoError(t, err)
	defer sink.Destroy()

	// send some messages and check they arrived
	err = faucet.SendFrame([]byte("Hello"), FlagNone)
	require.NoError(t, err)

	b, f, err := sink.RecvFrame()
	require.NoError(t, err)

	if want, have := false, f == FlagMore; want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if want, have := "Hello", string(b); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}
}

func ExampleProxy() {
	proxy := NewProxy()
	defer proxy.Destroy()

	// set front end address and socket type
	err := proxy.SetFrontend(Pull, "inproc://frontend")
	if err != nil {
		panic(err)
	}

	// set back end address and socket type
	err = proxy.SetBackend(Push, "inproc://backend")
	if err != nil {
		panic(err)
	}

	// set address for "tee"ing proxy traffic to
	err = proxy.SetCapture("inproc://capture")
	if err != nil {
		panic(err)
	}

	// we can pause the proxy
	err = proxy.Pause()
	if err != nil {
		panic(err)
	}

	// and we can resume it
	err = proxy.Resume()
	if err != nil {
		panic(err)
	}

	proxy.Destroy()
}

func benchmarkProxySendFrame(size int, b *testing.B) {
	proxy := NewProxy()
	defer proxy.Destroy()

	err := proxy.SetFrontend(Pull, fmt.Sprintf("inproc://benchProxyFront%d", size))
	require.NoError(b, err)

	err = proxy.SetBackend(Push, fmt.Sprintf("inproc://benchProxyBack%d", size))
	require.NoError(b, err)

	pullSock := NewSock(Pull)
	defer pullSock.Destroy()

	err = pullSock.Connect(fmt.Sprintf("inproc://benchProxyBack%d", size))
	require.NoError(b, err)

	go func() {
		pushSock := NewSock(Push)
		defer pushSock.Destroy()
		err := pushSock.Connect(fmt.Sprintf("inproc://benchProxyFront%d", size))
		require.NoError(b, err)

		payload := make([]byte, size)
		for i := 0; i < b.N; i++ {
			err = pushSock.SendFrame(payload, FlagNone)
			require.NoError(b, err)
		}
	}()

	for i := 0; i < b.N; i++ {
		msg, _, err := pullSock.RecvFrame()
		if err != nil {
			panic(err)
		}
		if len(msg) != size {
			panic("msg too small")
		}
		b.SetBytes(int64(size))
	}
}

func BenchmarkProxySendFrame1k(b *testing.B)  { benchmarkProxySendFrame(1024, b) }
func BenchmarkProxySendFrame4k(b *testing.B)  { benchmarkProxySendFrame(4096, b) }
func BenchmarkProxySendFrame16k(b *testing.B) { benchmarkProxySendFrame(16384, b) }
