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
func TestDeprecatedHeartbeatIvl(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 2000
	sock.SetHeartbeatIvl(testval)
	val := sock.HeartbeatIvl()
	if val != testval && val != 0 {
		t.Errorf("HeartbeatIvl returned %d, should be %d", val, testval)
	}
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

func TestDeprecatedHeartbeatTtl(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 4000
	sock.SetHeartbeatTtl(testval)
	val := sock.HeartbeatTtl()
	if val != testval && val != 0 {
		t.Errorf("HeartbeatTtl returned %d, should be %d", val, testval)
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

func TestDeprecatedHeartbeatTimeout(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 6000
	sock.SetHeartbeatTimeout(testval)
	val := sock.HeartbeatTimeout()
	if val != testval && val != 0 {
		t.Errorf("HeartbeatTimeout returned %d, should be %d", val, testval)
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

func TestDeprecatedUseFd(t *testing.T) {
	sock := NewSock(Req)
	testval := 3
	sock.SetUseFd(testval)
	val := sock.UseFd()
	if val != testval && val != 0 {
		t.Errorf("UseFd returned %d, should be %d", val, testval)
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

func TestDeprecatedXPubManual(t *testing.T) {
	sock := NewSock(XPub)
	testval := 1
	sock.SetXPubManual(testval)
	sock.Destroy()
}

func TestXPubManual(t *testing.T) {
	sock := NewSock(XPub)
	testval := 1
	sock.SetOption(SockSetXPubManual(testval))
	sock.Destroy()
}

func TestDeprecatedXPubWelcomeMsg(t *testing.T) {
	sock := NewSock(XPub)
	testval := "welcome"
	sock.SetXPubWelcomeMsg(testval)
	sock.Destroy()
}

func TestXPubWelcomeMsg(t *testing.T) {
	sock := NewSock(XPub)
	testval := "welcome"
	sock.SetOption(SockSetXPubWelcomeMsg(testval))
	sock.Destroy()
}

func TestDeprecatedStreamNotify(t *testing.T) {
	sock := NewSock(Stream)
	testval := 1
	sock.SetStreamNotify(testval)
	sock.Destroy()
}

func TestStreamNotify(t *testing.T) {
	sock := NewSock(Stream)
	testval := 1
	sock.SetOption(SockSetStreamNotify(testval))
	sock.Destroy()
}

func TestDeprecatedInvertMatching(t *testing.T) {
	sock := NewSock(XPub)
	testval := 1
	sock.SetInvertMatching(testval)
	val := sock.InvertMatching()
	if val != testval && val != 0 {
		t.Errorf("InvertMatching returned %d, should be %d", val, testval)
	}
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

func TestDeprecatedXPubVerboser(t *testing.T) {
	sock := NewSock(XPub)
	testval := 1
	sock.SetXPubVerboser(testval)
	sock.Destroy()
}

func TestXPubVerboser(t *testing.T) {
	sock := NewSock(XPub)
	testval := 1
	sock.SetOption(SockSetXPubVerboser(testval))
	sock.Destroy()
}

func TestDeprecatedConnectTimeout(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 200
	sock.SetConnectTimeout(testval)
	val := sock.ConnectTimeout()
	if val != testval && val != 0 {
		t.Errorf("ConnectTimeout returned %d, should be %d", val, testval)
	}
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

func TestDeprecatedTcpMaxrt(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 200
	sock.SetTcpMaxrt(testval)
	val := sock.TcpMaxrt()
	if val != testval && val != 0 {
		t.Errorf("TcpMaxrt returned %d, should be %d", val, testval)
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

func TestDeprecatedMulticastMaxtpdu(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 1400
	sock.SetMulticastMaxtpdu(testval)
	val := sock.MulticastMaxtpdu()
	if val != testval && val != 0 {
		t.Errorf("MulticastMaxtpdu returned %d, should be %d", val, testval)
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

func TestDeprecatedConnectRid(t *testing.T) {
	sock := NewSock(Router)
	testval := "ABCD"
	sock.SetConnectRid(testval)
	sock.Destroy()
}

func TestConnectRid(t *testing.T) {
	sock := NewSock(Router)
	testval := "ABCD"
	sock.SetOption(SockSetConnectRid(testval))
	sock.Destroy()
}

func TestDeprecatedHandshakeIvl(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 200
	sock.SetHandshakeIvl(testval)
	val := sock.HandshakeIvl()
	if val != testval && val != 0 {
		t.Errorf("HandshakeIvl returned %d, should be %d", val, testval)
	}
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

func TestDeprecatedSocksProxy(t *testing.T) {
	sock := NewSock(Dealer)
	testval := "127.0.0.1"
	sock.SetSocksProxy(testval)
	val := sock.SocksProxy()
	if val != testval && val != "" {
		t.Errorf("SocksProxy returned %s should be %s", val, testval)
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

func TestDeprecatedXPubNodrop(t *testing.T) {
	sock := NewSock(XPub)
	testval := 1
	sock.SetXPubNodrop(testval)
	sock.Destroy()
}

func TestXPubNodrop(t *testing.T) {
	sock := NewSock(XPub)
	testval := 1
	sock.SetOption(SockSetXPubNodrop(testval))
	sock.Destroy()
}

func TestDeprecatedTos(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 1
	sock.SetTos(testval)
	val := sock.Tos()
	if val != testval && val != 0 {
		t.Errorf("Tos returned %d, should be %d", val, testval)
	}
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

func TestDeprecatedRouterHandover(t *testing.T) {
	sock := NewSock(Router)
	testval := 1
	sock.SetRouterHandover(testval)
	sock.Destroy()
}

func TestRouterHandover(t *testing.T) {
	sock := NewSock(Router)
	testval := 1
	sock.SetOption(SockSetRouterHandover(testval))
	sock.Destroy()
}

func TestDeprecatedRouterMandatory(t *testing.T) {
	sock := NewSock(Router)
	testval := 1
	sock.SetRouterMandatory(testval)
	sock.Destroy()
}

func TestRouterMandatory(t *testing.T) {
	sock := NewSock(Router)
	testval := 1
	sock.SetOption(SockSetRouterMandatory(testval))
	sock.Destroy()
}

func TestDeprecatedProbeRouter(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 1
	sock.SetProbeRouter(testval)
	sock.Destroy()
}

func TestProbeRouter(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 1
	sock.SetOption(SockSetProbeRouter(testval))
	sock.Destroy()
}

func TestDeprecatedReqRelaxed(t *testing.T) {
	sock := NewSock(Req)
	testval := 1
	sock.SetReqRelaxed(testval)
	sock.Destroy()
}

func TestReqRelaxed(t *testing.T) {
	sock := NewSock(Req)
	testval := 1
	sock.SetOption(SockSetReqRelaxed(testval))
	sock.Destroy()
}

func TestDeprecatedReqCorrelate(t *testing.T) {
	sock := NewSock(Req)
	testval := 1
	sock.SetReqCorrelate(testval)
	sock.Destroy()
}

func TestReqCorrelate(t *testing.T) {
	sock := NewSock(Req)
	testval := 1
	sock.SetOption(SockSetReqCorrelate(testval))
	sock.Destroy()
}

func TestDeprecatedConflate(t *testing.T) {
	sock := NewSock(Push)
	testval := 1
	sock.SetConflate(testval)
	sock.Destroy()
}

func TestConflate(t *testing.T) {
	sock := NewSock(Push)
	testval := 1
	sock.SetOption(SockSetConflate(testval))
	sock.Destroy()
}

func TestDeprecatedZapDomain(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetZapDomain(testval)
	val := sock.ZapDomain()
	if val != testval && val != "" {
		t.Errorf("ZapDomain returned %s should be %s", val, testval)
	}
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

func TestDeprecatedPlainServer(t *testing.T) {
	sock := NewSock(Pub)
	testval := 1
	sock.SetPlainServer(testval)
	val := sock.PlainServer()
	if val != testval && val != 0 {
		t.Errorf("PlainServer returned %d, should be %d", val, testval)
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

func TestDeprecatedPlainUsername(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetPlainUsername(testval)
	val := sock.PlainUsername()
	if val != testval && val != "" {
		t.Errorf("PlainUsername returned %s should be %s", val, testval)
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

func TestDeprecatedPlainPassword(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetPlainPassword(testval)
	val := sock.PlainPassword()
	if val != testval && val != "" {
		t.Errorf("PlainPassword returned %s should be %s", val, testval)
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

func TestDeprecatedIpv6(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetIpv6(testval)
	val := sock.Ipv6()
	if val != testval && val != 0 {
		t.Errorf("Ipv6 returned %d, should be %d", val, testval)
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

func TestDeprecatedImmediate(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 1
	sock.SetImmediate(testval)
	val := sock.Immediate()
	if val != testval && val != 0 {
		t.Errorf("Immediate returned %d, should be %d", val, testval)
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

func TestDeprecatedRouterRaw(t *testing.T) {
	sock := NewSock(Router)
	testval := 1
	sock.SetRouterRaw(testval)
	sock.Destroy()
}

func TestRouterRaw(t *testing.T) {
	sock := NewSock(Router)
	testval := 1
	sock.SetOption(SockSetRouterRaw(testval))
	sock.Destroy()
}

func TestDeprecatedIpv4only(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetIpv4only(testval)
	val := sock.Ipv4only()
	if val != testval && val != 0 {
		t.Errorf("Ipv4only returned %d, should be %d", val, testval)
	}
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

func TestDeprecatedDelayAttachOnConnect(t *testing.T) {
	sock := NewSock(Pub)
	testval := 1
	sock.SetDelayAttachOnConnect(testval)
	sock.Destroy()
}

func TestDelayAttachOnConnect(t *testing.T) {
	sock := NewSock(Pub)
	testval := 1
	sock.SetOption(SockSetDelayAttachOnConnect(testval))
	sock.Destroy()
}

func TestDeprecatedSndhwm(t *testing.T) {
	sock := NewSock(Pub)
	testval := 1
	sock.SetSndhwm(testval)
	val := sock.Sndhwm()
	if val != testval && val != 0 {
		t.Errorf("Sndhwm returned %d, should be %d", val, testval)
	}
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

func TestDeprecatedRcvhwm(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetRcvhwm(testval)
	val := sock.Rcvhwm()
	if val != testval && val != 0 {
		t.Errorf("Rcvhwm returned %d, should be %d", val, testval)
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

func TestDeprecatedAffinity(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetAffinity(testval)
	val := sock.Affinity()
	if val != testval && val != 0 {
		t.Errorf("Affinity returned %d, should be %d", val, testval)
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

func TestDeprecatedSubscribe(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetSubscribe(testval)
	sock.Destroy()
}

func TestSubscribe(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetOption(SockSetSubscribe(testval))
	sock.Destroy()
}

func TestDeprecatedUnsubscribe(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetUnsubscribe(testval)
	sock.Destroy()
}

func TestUnsubscribe(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetOption(SockSetUnsubscribe(testval))
	sock.Destroy()
}

func TestDeprecatedIdentity(t *testing.T) {
	sock := NewSock(Dealer)
	testval := "test"
	sock.SetIdentity(testval)
	val := sock.Identity()
	if val != testval && val != "" {
		t.Errorf("Identity returned %s should be %s", val, testval)
	}
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

func TestDeprecatedRate(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetRate(testval)
	val := sock.Rate()
	if val != testval && val != 0 {
		t.Errorf("Rate returned %d, should be %d", val, testval)
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

func TestDeprecatedRecoveryIvl(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetRecoveryIvl(testval)
	val := sock.RecoveryIvl()
	if val != testval && val != 0 {
		t.Errorf("RecoveryIvl returned %d, should be %d", val, testval)
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

func TestDeprecatedSndbuf(t *testing.T) {
	sock := NewSock(Pub)
	testval := 1
	sock.SetSndbuf(testval)
	val := sock.Sndbuf()
	if val != testval && val != 0 {
		t.Errorf("Sndbuf returned %d, should be %d", val, testval)
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

func TestDeprecatedRcvbuf(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetRcvbuf(testval)
	val := sock.Rcvbuf()
	if val != testval && val != 0 {
		t.Errorf("Rcvbuf returned %d, should be %d", val, testval)
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

func TestDeprecatedLinger(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetLinger(testval)
	val := sock.Linger()
	if val != testval && val != 0 {
		t.Errorf("Linger returned %d, should be %d", val, testval)
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

func TestDeprecatedReconnectIvl(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetReconnectIvl(testval)
	val := sock.ReconnectIvl()
	if val != testval && val != 0 {
		t.Errorf("ReconnectIvl returned %d, should be %d", val, testval)
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

func TestDeprecatedReconnectIvlMax(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetReconnectIvlMax(testval)
	val := sock.ReconnectIvlMax()
	if val != testval && val != 0 {
		t.Errorf("ReconnectIvlMax returned %d, should be %d", val, testval)
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

func TestDeprecatedBacklog(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetBacklog(testval)
	val := sock.Backlog()
	if val != testval && val != 0 {
		t.Errorf("Backlog returned %d, should be %d", val, testval)
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

func TestDeprecatedMaxmsgsize(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetMaxmsgsize(testval)
	val := sock.Maxmsgsize()
	if val != testval && val != 0 {
		t.Errorf("Maxmsgsize returned %d, should be %d", val, testval)
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

func TestDeprecatedMulticastHops(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetMulticastHops(testval)
	val := sock.MulticastHops()
	if val != testval && val != 0 {
		t.Errorf("MulticastHops returned %d, should be %d", val, testval)
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

func TestDeprecatedRcvtimeo(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetRcvtimeo(testval)
	val := sock.Rcvtimeo()
	if val != testval && val != 0 {
		t.Errorf("Rcvtimeo returned %d, should be %d", val, testval)
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

func TestDeprecatedSndtimeo(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetSndtimeo(testval)
	val := sock.Sndtimeo()
	if val != testval && val != 0 {
		t.Errorf("Sndtimeo returned %d, should be %d", val, testval)
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

func TestDeprecatedXPubVerbose(t *testing.T) {
	sock := NewSock(XPub)
	testval := 1
	sock.SetXPubVerbose(testval)
	sock.Destroy()
}

func TestXPubVerbose(t *testing.T) {
	sock := NewSock(XPub)
	testval := 1
	sock.SetOption(SockSetXPubVerbose(testval))
	sock.Destroy()
}

func TestDeprecatedTcpKeepalive(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetTcpKeepalive(testval)
	val := sock.TcpKeepalive()
	if val != testval && val != 0 {
		t.Errorf("TcpKeepalive returned %d, should be %d", val, testval)
	}
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

func TestDeprecatedTcpKeepaliveIdle(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetTcpKeepaliveIdle(testval)
	val := sock.TcpKeepaliveIdle()
	if val != testval && val != 0 {
		t.Errorf("TcpKeepaliveIdle returned %d, should be %d", val, testval)
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

func TestDeprecatedTcpKeepaliveCnt(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetTcpKeepaliveCnt(testval)
	val := sock.TcpKeepaliveCnt()
	if val != testval && val != 0 {
		t.Errorf("TcpKeepaliveCnt returned %d, should be %d", val, testval)
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

func TestDeprecatedTcpKeepaliveIntvl(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetTcpKeepaliveIntvl(testval)
	val := sock.TcpKeepaliveIntvl()
	if val != testval && val != 0 {
		t.Errorf("TcpKeepaliveIntvl returned %d, should be %d", val, testval)
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

func TestDeprecatedTcpAcceptFilter(t *testing.T) {
	sock := NewSock(Sub)
	testval := "127.0.0.1"
	sock.SetTcpAcceptFilter(testval)
	val := sock.TcpAcceptFilter()
	if val != testval && val != "" {
		t.Errorf("TcpAcceptFilter returned %s should be %s", val, testval)
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

