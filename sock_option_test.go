//go:generate gsl sockopts.xml
package goczmq
/*  =========================================================================
    zsock_option - get/set 0MQ socket options

            ****************************************************
            *   GENERATED SOURCE CODE, DO NOT EDIT!!           *
            ****************************************************

    Copyright (c) the Contributors as noted in the AUTHORS file.
    This file is part of goczmq, the high-level go binding for CZMQ:
    http://github.com/zeromq/goczmq

    This Source Code Form is subject to the terms of the Mozilla Public
    License, v. 2.0. If a copy of the MPL was not distributed with this
    file, You can obtain one at http://mozilla.org/MPL/2.0/.
    =========================================================================
*/

import (
	"testing"
)
func TestRouterNotify(t *testing.T) {
	sock := NewSock(Router)
	testval := 1
	sock.SetOption(SockSetRouterNotify(testval))
	sock.Destroy()
}

func TestHeartbeatIvl(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 2000
	sock.SetOption(SockSetHeartbeatIvl(testval))
	val := HeartbeatIvl(sock)
	if val != testval && val != 0 {
		t.Errorf("HeartbeatIvl returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestHeartbeatTtl(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 4000
	sock.SetOption(SockSetHeartbeatTtl(testval))
	val := HeartbeatTtl(sock)
	if val != testval && val != 0 {
		t.Errorf("HeartbeatTtl returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestHeartbeatTimeout(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 6000
	sock.SetOption(SockSetHeartbeatTimeout(testval))
	val := HeartbeatTimeout(sock)
	if val != testval && val != 0 {
		t.Errorf("HeartbeatTimeout returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestUseFd(t *testing.T) {
	sock := NewSock(Req)
	testval := 3
	sock.SetOption(SockSetUseFd(testval))
	val := UseFd(sock)
	if val != testval && val != 0 {
		t.Errorf("UseFd returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestXPubManual(t *testing.T) {
	sock := NewSock(XPub)
	testval := 1
	sock.SetOption(SockSetXPubManual(testval))
	sock.Destroy()
}

func TestXPubWelcomeMsg(t *testing.T) {
	sock := NewSock(XPub)
	testval := "welcome"
	sock.SetOption(SockSetXPubWelcomeMsg(testval))
	sock.Destroy()
}

func TestStreamNotify(t *testing.T) {
	sock := NewSock(Stream)
	testval := 1
	sock.SetOption(SockSetStreamNotify(testval))
	sock.Destroy()
}

func TestInvertMatching(t *testing.T) {
	sock := NewSock(XPub)
	testval := 1
	sock.SetOption(SockSetInvertMatching(testval))
	val := InvertMatching(sock)
	if val != testval && val != 0 {
		t.Errorf("InvertMatching returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestXPubVerboser(t *testing.T) {
	sock := NewSock(XPub)
	testval := 1
	sock.SetOption(SockSetXPubVerboser(testval))
	sock.Destroy()
}

func TestConnectTimeout(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 200
	sock.SetOption(SockSetConnectTimeout(testval))
	val := ConnectTimeout(sock)
	if val != testval && val != 0 {
		t.Errorf("ConnectTimeout returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestTcpMaxrt(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 200
	sock.SetOption(SockSetTcpMaxrt(testval))
	val := TcpMaxrt(sock)
	if val != testval && val != 0 {
		t.Errorf("TcpMaxrt returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestMulticastMaxtpdu(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 1400
	sock.SetOption(SockSetMulticastMaxtpdu(testval))
	val := MulticastMaxtpdu(sock)
	if val != testval && val != 0 {
		t.Errorf("MulticastMaxtpdu returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestConnectRid(t *testing.T) {
	sock := NewSock(Router)
	testval := "ABCD"
	sock.SetOption(SockSetConnectRid(testval))
	sock.Destroy()
}

func TestHandshakeIvl(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 200
	sock.SetOption(SockSetHandshakeIvl(testval))
	val := HandshakeIvl(sock)
	if val != testval && val != 0 {
		t.Errorf("HandshakeIvl returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestSocksProxy(t *testing.T) {
	sock := NewSock(Dealer)
	testval := "127.0.0.1"
	sock.SetOption(SockSetSocksProxy(testval))
	val := SocksProxy(sock)
	if val != testval && val != "" {
		t.Errorf("SocksProxy returned %s should be %s", val, testval)
	}
	sock.Destroy()
}

func TestXPubNodrop(t *testing.T) {
	sock := NewSock(XPub)
	testval := 1
	sock.SetOption(SockSetXPubNodrop(testval))
	sock.Destroy()
}

func TestTos(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 1
	sock.SetOption(SockSetTos(testval))
	val := Tos(sock)
	if val != testval && val != 0 {
		t.Errorf("Tos returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestRouterHandover(t *testing.T) {
	sock := NewSock(Router)
	testval := 1
	sock.SetOption(SockSetRouterHandover(testval))
	sock.Destroy()
}

func TestRouterMandatory(t *testing.T) {
	sock := NewSock(Router)
	testval := 1
	sock.SetOption(SockSetRouterMandatory(testval))
	sock.Destroy()
}

func TestProbeRouter(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 1
	sock.SetOption(SockSetProbeRouter(testval))
	sock.Destroy()
}

func TestReqRelaxed(t *testing.T) {
	sock := NewSock(Req)
	testval := 1
	sock.SetOption(SockSetReqRelaxed(testval))
	sock.Destroy()
}

func TestReqCorrelate(t *testing.T) {
	sock := NewSock(Req)
	testval := 1
	sock.SetOption(SockSetReqCorrelate(testval))
	sock.Destroy()
}

func TestConflate(t *testing.T) {
	sock := NewSock(Push)
	testval := 1
	sock.SetOption(SockSetConflate(testval))
	sock.Destroy()
}

func TestZapDomain(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetOption(SockSetZapDomain(testval))
	val := ZapDomain(sock)
	if val != testval && val != "" {
		t.Errorf("ZapDomain returned %s should be %s", val, testval)
	}
	sock.Destroy()
}

func TestPlainServer(t *testing.T) {
	sock := NewSock(Pub)
	testval := 1
	sock.SetOption(SockSetPlainServer(testval))
	val := PlainServer(sock)
	if val != testval && val != 0 {
		t.Errorf("PlainServer returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestPlainUsername(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetOption(SockSetPlainUsername(testval))
	val := PlainUsername(sock)
	if val != testval && val != "" {
		t.Errorf("PlainUsername returned %s should be %s", val, testval)
	}
	sock.Destroy()
}

func TestPlainPassword(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetOption(SockSetPlainPassword(testval))
	val := PlainPassword(sock)
	if val != testval && val != "" {
		t.Errorf("PlainPassword returned %s should be %s", val, testval)
	}
	sock.Destroy()
}

func TestIpv6(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetIpv6(testval))
	val := Ipv6(sock)
	if val != testval && val != 0 {
		t.Errorf("Ipv6 returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestImmediate(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 1
	sock.SetOption(SockSetImmediate(testval))
	val := Immediate(sock)
	if val != testval && val != 0 {
		t.Errorf("Immediate returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestRouterRaw(t *testing.T) {
	sock := NewSock(Router)
	testval := 1
	sock.SetOption(SockSetRouterRaw(testval))
	sock.Destroy()
}

func TestIpv4only(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetIpv4only(testval))
	val := Ipv4only(sock)
	if val != testval && val != 0 {
		t.Errorf("Ipv4only returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestDelayAttachOnConnect(t *testing.T) {
	sock := NewSock(Pub)
	testval := 1
	sock.SetOption(SockSetDelayAttachOnConnect(testval))
	sock.Destroy()
}

func TestSndhwm(t *testing.T) {
	sock := NewSock(Pub)
	testval := 1
	sock.SetOption(SockSetSndhwm(testval))
	val := Sndhwm(sock)
	if val != testval && val != 0 {
		t.Errorf("Sndhwm returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestRcvhwm(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetRcvhwm(testval))
	val := Rcvhwm(sock)
	if val != testval && val != 0 {
		t.Errorf("Rcvhwm returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestAffinity(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetAffinity(testval))
	val := Affinity(sock)
	if val != testval && val != 0 {
		t.Errorf("Affinity returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestSubscribe(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetOption(SockSetSubscribe(testval))
	sock.Destroy()
}

func TestUnsubscribe(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetOption(SockSetUnsubscribe(testval))
	sock.Destroy()
}

func TestIdentity(t *testing.T) {
	sock := NewSock(Dealer)
	testval := "test"
	sock.SetOption(SockSetIdentity(testval))
	val := Identity(sock)
	if val != testval && val != "" {
		t.Errorf("Identity returned %s should be %s", val, testval)
	}
	sock.Destroy()
}

func TestRate(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetRate(testval))
	val := Rate(sock)
	if val != testval && val != 0 {
		t.Errorf("Rate returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestRecoveryIvl(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetRecoveryIvl(testval))
	val := RecoveryIvl(sock)
	if val != testval && val != 0 {
		t.Errorf("RecoveryIvl returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestSndbuf(t *testing.T) {
	sock := NewSock(Pub)
	testval := 1
	sock.SetOption(SockSetSndbuf(testval))
	val := Sndbuf(sock)
	if val != testval && val != 0 {
		t.Errorf("Sndbuf returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestRcvbuf(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetRcvbuf(testval))
	val := Rcvbuf(sock)
	if val != testval && val != 0 {
		t.Errorf("Rcvbuf returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestLinger(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetLinger(testval))
	val := Linger(sock)
	if val != testval && val != 0 {
		t.Errorf("Linger returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestReconnectIvl(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetReconnectIvl(testval))
	val := ReconnectIvl(sock)
	if val != testval && val != 0 {
		t.Errorf("ReconnectIvl returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestReconnectIvlMax(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetReconnectIvlMax(testval))
	val := ReconnectIvlMax(sock)
	if val != testval && val != 0 {
		t.Errorf("ReconnectIvlMax returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestBacklog(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetBacklog(testval))
	val := Backlog(sock)
	if val != testval && val != 0 {
		t.Errorf("Backlog returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestMaxmsgsize(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetMaxmsgsize(testval))
	val := Maxmsgsize(sock)
	if val != testval && val != 0 {
		t.Errorf("Maxmsgsize returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestMulticastHops(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetMulticastHops(testval))
	val := MulticastHops(sock)
	if val != testval && val != 0 {
		t.Errorf("MulticastHops returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestRcvtimeo(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetRcvtimeo(testval))
	val := Rcvtimeo(sock)
	if val != testval && val != 0 {
		t.Errorf("Rcvtimeo returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestSndtimeo(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetSndtimeo(testval))
	val := Sndtimeo(sock)
	if val != testval && val != 0 {
		t.Errorf("Sndtimeo returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestXPubVerbose(t *testing.T) {
	sock := NewSock(XPub)
	testval := 1
	sock.SetOption(SockSetXPubVerbose(testval))
	sock.Destroy()
}

func TestTcpKeepalive(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetTcpKeepalive(testval))
	val := TcpKeepalive(sock)
	if val != testval && val != 0 {
		t.Errorf("TcpKeepalive returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestTcpKeepaliveIdle(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetTcpKeepaliveIdle(testval))
	val := TcpKeepaliveIdle(sock)
	if val != testval && val != 0 {
		t.Errorf("TcpKeepaliveIdle returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestTcpKeepaliveCnt(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetTcpKeepaliveCnt(testval))
	val := TcpKeepaliveCnt(sock)
	if val != testval && val != 0 {
		t.Errorf("TcpKeepaliveCnt returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestTcpKeepaliveIntvl(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetOption(SockSetTcpKeepaliveIntvl(testval))
	val := TcpKeepaliveIntvl(sock)
	if val != testval && val != 0 {
		t.Errorf("TcpKeepaliveIntvl returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestTcpAcceptFilter(t *testing.T) {
	sock := NewSock(Sub)
	testval := "127.0.0.1"
	sock.SetOption(SockSetTcpAcceptFilter(testval))
	val := TcpAcceptFilter(sock)
	if val != testval && val != "" {
		t.Errorf("TcpAcceptFilter returned %s should be %s", val, testval)
	}
	sock.Destroy()
}

