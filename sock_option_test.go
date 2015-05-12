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
func TestTos(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 1
	sock.SetTos(testval)
	val := sock.Tos()
	if val != testval && val != 0 {
		t.Errorf("Tos returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestRouterHandover(t *testing.T) {
	sock := NewSock(Router)
	testval := 1
	sock.SetRouterHandover(testval)
	sock.Destroy()
}

func TestRouterMandatory(t *testing.T) {
	sock := NewSock(Router)
	testval := 1
	sock.SetRouterMandatory(testval)
	sock.Destroy()
}

func TestProbeRouter(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 1
	sock.SetProbeRouter(testval)
	sock.Destroy()
}

func TestReqRelaxed(t *testing.T) {
	sock := NewSock(Req)
	testval := 1
	sock.SetReqRelaxed(testval)
	sock.Destroy()
}

func TestReqCorrelate(t *testing.T) {
	sock := NewSock(Req)
	testval := 1
	sock.SetReqCorrelate(testval)
	sock.Destroy()
}

func TestConflate(t *testing.T) {
	sock := NewSock(Push)
	testval := 1
	sock.SetConflate(testval)
	sock.Destroy()
}

func TestZapDomain(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetZapDomain(testval)
	val := sock.ZapDomain()
	if val != testval && val != "" {
		t.Errorf("ZapDomain returned %s should be %s", val, testval)
	}
	sock.Destroy()
}

func TestPlainServer(t *testing.T) {
	sock := NewSock(Pub)
	testval := 1
	sock.SetPlainServer(testval)
	val := sock.PlainServer()
	if val != testval && val != 0 {
		t.Errorf("PlainServer returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestPlainUsername(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetPlainUsername(testval)
	val := sock.PlainUsername()
	if val != testval && val != "" {
		t.Errorf("PlainUsername returned %s should be %s", val, testval)
	}
	sock.Destroy()
}

func TestPlainPassword(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetPlainPassword(testval)
	val := sock.PlainPassword()
	if val != testval && val != "" {
		t.Errorf("PlainPassword returned %s should be %s", val, testval)
	}
	sock.Destroy()
}

func TestIpv6(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetIpv6(testval)
	val := sock.Ipv6()
	if val != testval && val != 0 {
		t.Errorf("Ipv6 returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestImmediate(t *testing.T) {
	sock := NewSock(Dealer)
	testval := 1
	sock.SetImmediate(testval)
	val := sock.Immediate()
	if val != testval && val != 0 {
		t.Errorf("Immediate returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestRouterRaw(t *testing.T) {
	sock := NewSock(Router)
	testval := 1
	sock.SetRouterRaw(testval)
	sock.Destroy()
}

func TestIpv4only(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetIpv4only(testval)
	val := sock.Ipv4only()
	if val != testval && val != 0 {
		t.Errorf("Ipv4only returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestDelayAttachOnConnect(t *testing.T) {
	sock := NewSock(Pub)
	testval := 1
	sock.SetDelayAttachOnConnect(testval)
	sock.Destroy()
}

func TestSndhwm(t *testing.T) {
	sock := NewSock(Pub)
	testval := 1
	sock.SetSndhwm(testval)
	val := sock.Sndhwm()
	if val != testval && val != 0 {
		t.Errorf("Sndhwm returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestRcvhwm(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetRcvhwm(testval)
	val := sock.Rcvhwm()
	if val != testval && val != 0 {
		t.Errorf("Rcvhwm returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestAffinity(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetAffinity(testval)
	val := sock.Affinity()
	if val != testval && val != 0 {
		t.Errorf("Affinity returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestSubscribe(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetSubscribe(testval)
	sock.Destroy()
}

func TestUnsubscribe(t *testing.T) {
	sock := NewSock(Sub)
	testval := "test"
	sock.SetUnsubscribe(testval)
	sock.Destroy()
}

func TestIdentity(t *testing.T) {
	sock := NewSock(Dealer)
	testval := "test"
	sock.SetIdentity(testval)
	val := sock.Identity()
	if val != testval && val != "" {
		t.Errorf("Identity returned %s should be %s", val, testval)
	}
	sock.Destroy()
}

func TestRate(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetRate(testval)
	val := sock.Rate()
	if val != testval && val != 0 {
		t.Errorf("Rate returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestRecoveryIvl(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetRecoveryIvl(testval)
	val := sock.RecoveryIvl()
	if val != testval && val != 0 {
		t.Errorf("RecoveryIvl returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestSndbuf(t *testing.T) {
	sock := NewSock(Pub)
	testval := 1
	sock.SetSndbuf(testval)
	val := sock.Sndbuf()
	if val != testval && val != 0 {
		t.Errorf("Sndbuf returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestRcvbuf(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetRcvbuf(testval)
	val := sock.Rcvbuf()
	if val != testval && val != 0 {
		t.Errorf("Rcvbuf returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestLinger(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetLinger(testval)
	val := sock.Linger()
	if val != testval && val != 0 {
		t.Errorf("Linger returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestReconnectIvl(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetReconnectIvl(testval)
	val := sock.ReconnectIvl()
	if val != testval && val != 0 {
		t.Errorf("ReconnectIvl returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestReconnectIvlMax(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetReconnectIvlMax(testval)
	val := sock.ReconnectIvlMax()
	if val != testval && val != 0 {
		t.Errorf("ReconnectIvlMax returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestBacklog(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetBacklog(testval)
	val := sock.Backlog()
	if val != testval && val != 0 {
		t.Errorf("Backlog returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestMaxmsgsize(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetMaxmsgsize(testval)
	val := sock.Maxmsgsize()
	if val != testval && val != 0 {
		t.Errorf("Maxmsgsize returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestMulticastHops(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetMulticastHops(testval)
	val := sock.MulticastHops()
	if val != testval && val != 0 {
		t.Errorf("MulticastHops returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestRcvtimeo(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetRcvtimeo(testval)
	val := sock.Rcvtimeo()
	if val != testval && val != 0 {
		t.Errorf("Rcvtimeo returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestSndtimeo(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetSndtimeo(testval)
	val := sock.Sndtimeo()
	if val != testval && val != 0 {
		t.Errorf("Sndtimeo returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestXpubVerbose(t *testing.T) {
	sock := NewSock(XPub)
	testval := 1
	sock.SetXpubVerbose(testval)
	sock.Destroy()
}

func TestTcpKeepalive(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetTcpKeepalive(testval)
	val := sock.TcpKeepalive()
	if val != testval && val != 0 {
		t.Errorf("TcpKeepalive returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestTcpKeepaliveIdle(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetTcpKeepaliveIdle(testval)
	val := sock.TcpKeepaliveIdle()
	if val != testval && val != 0 {
		t.Errorf("TcpKeepaliveIdle returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestTcpKeepaliveCnt(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetTcpKeepaliveCnt(testval)
	val := sock.TcpKeepaliveCnt()
	if val != testval && val != 0 {
		t.Errorf("TcpKeepaliveCnt returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestTcpKeepaliveIntvl(t *testing.T) {
	sock := NewSock(Sub)
	testval := 1
	sock.SetTcpKeepaliveIntvl(testval)
	val := sock.TcpKeepaliveIntvl()
	if val != testval && val != 0 {
		t.Errorf("TcpKeepaliveIntvl returned %d, should be %d", val, testval)
	}
	sock.Destroy()
}

func TestTcpAcceptFilter(t *testing.T) {
	sock := NewSock(Sub)
	testval := "127.0.0.1"
	sock.SetTcpAcceptFilter(testval)
	val := sock.TcpAcceptFilter()
	if val != testval && val != "" {
		t.Errorf("TcpAcceptFilter returned %s should be %s", val, testval)
	}
	sock.Destroy()
}

