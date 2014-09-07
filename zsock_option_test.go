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
	sock := NewZsock(DEALER)
	expected := 1
	sock.SetTos(expected)
	val := sock.Tos()
	if val != expected && val != 0 {
		t.Errorf("Tos returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestZapDomain(t *testing.T) {
	sock := NewZsock(SUB)
	expected := "test"
	sock.SetZapDomain(expected)
	val := sock.ZapDomain()
	if val != expected && val != "" {
		t.Errorf("ZapDomain returned %s should be %s", val, expected)
	}
        sock.Destroy()
}
func TestPlainServer(t *testing.T) {
	sock := NewZsock(PUB)
	expected := 1
	sock.SetPlainServer(expected)
	val := sock.PlainServer()
	if val != expected && val != 0 {
		t.Errorf("PlainServer returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestPlainUsername(t *testing.T) {
	sock := NewZsock(SUB)
	expected := "test"
	sock.SetPlainUsername(expected)
	val := sock.PlainUsername()
	if val != expected && val != "" {
		t.Errorf("PlainUsername returned %s should be %s", val, expected)
	}
        sock.Destroy()
}
func TestPlainPassword(t *testing.T) {
	sock := NewZsock(SUB)
	expected := "test"
	sock.SetPlainPassword(expected)
	val := sock.PlainPassword()
	if val != expected && val != "" {
		t.Errorf("PlainPassword returned %s should be %s", val, expected)
	}
        sock.Destroy()
}
func TestIpv6(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetIpv6(expected)
	val := sock.Ipv6()
	if val != expected && val != 0 {
		t.Errorf("Ipv6 returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestImmediate(t *testing.T) {
	sock := NewZsock(DEALER)
	expected := 1
	sock.SetImmediate(expected)
	val := sock.Immediate()
	if val != expected && val != 0 {
		t.Errorf("Immediate returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestIpv4only(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetIpv4only(expected)
	val := sock.Ipv4only()
	if val != expected && val != 0 {
		t.Errorf("Ipv4only returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestSndhwm(t *testing.T) {
	sock := NewZsock(PUB)
	expected := 1
	sock.SetSndhwm(expected)
	val := sock.Sndhwm()
	if val != expected && val != 0 {
		t.Errorf("Sndhwm returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestRcvhwm(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetRcvhwm(expected)
	val := sock.Rcvhwm()
	if val != expected && val != 0 {
		t.Errorf("Rcvhwm returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestIdentity(t *testing.T) {
	sock := NewZsock(DEALER)
	expected := "test"
	sock.SetIdentity(expected)
	val := sock.Identity()
	if val != expected && val != "" {
		t.Errorf("Identity returned %s should be %s", val, expected)
	}
        sock.Destroy()
}
func TestRate(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetRate(expected)
	val := sock.Rate()
	if val != expected && val != 0 {
		t.Errorf("Rate returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestRecoveryIvl(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetRecoveryIvl(expected)
	val := sock.RecoveryIvl()
	if val != expected && val != 0 {
		t.Errorf("RecoveryIvl returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestSndbuf(t *testing.T) {
	sock := NewZsock(PUB)
	expected := 1
	sock.SetSndbuf(expected)
	val := sock.Sndbuf()
	if val != expected && val != 0 {
		t.Errorf("Sndbuf returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestRcvbuf(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetRcvbuf(expected)
	val := sock.Rcvbuf()
	if val != expected && val != 0 {
		t.Errorf("Rcvbuf returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestLinger(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetLinger(expected)
	val := sock.Linger()
	if val != expected && val != 0 {
		t.Errorf("Linger returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestReconnectIvl(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetReconnectIvl(expected)
	val := sock.ReconnectIvl()
	if val != expected && val != 0 {
		t.Errorf("ReconnectIvl returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestReconnectIvlMax(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetReconnectIvlMax(expected)
	val := sock.ReconnectIvlMax()
	if val != expected && val != 0 {
		t.Errorf("ReconnectIvlMax returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestBacklog(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetBacklog(expected)
	val := sock.Backlog()
	if val != expected && val != 0 {
		t.Errorf("Backlog returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestMulticastHops(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetMulticastHops(expected)
	val := sock.MulticastHops()
	if val != expected && val != 0 {
		t.Errorf("MulticastHops returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestRcvtimeo(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetRcvtimeo(expected)
	val := sock.Rcvtimeo()
	if val != expected && val != 0 {
		t.Errorf("Rcvtimeo returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestSndtimeo(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetSndtimeo(expected)
	val := sock.Sndtimeo()
	if val != expected && val != 0 {
		t.Errorf("Sndtimeo returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestTcpKeepalive(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetTcpKeepalive(expected)
	val := sock.TcpKeepalive()
	if val != expected && val != 0 {
		t.Errorf("TcpKeepalive returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestTcpKeepaliveIdle(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetTcpKeepaliveIdle(expected)
	val := sock.TcpKeepaliveIdle()
	if val != expected && val != 0 {
		t.Errorf("TcpKeepaliveIdle returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestTcpKeepaliveCnt(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetTcpKeepaliveCnt(expected)
	val := sock.TcpKeepaliveCnt()
	if val != expected && val != 0 {
		t.Errorf("TcpKeepaliveCnt returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestTcpKeepaliveIntvl(t *testing.T) {
	sock := NewZsock(SUB)
	expected := 1
	sock.SetTcpKeepaliveIntvl(expected)
	val := sock.TcpKeepaliveIntvl()
	if val != expected && val != 0 {
		t.Errorf("TcpKeepaliveIntvl returned %d, should be %d", val, expected)
	}
	sock.Destroy()
}

func TestTcpAcceptFilter(t *testing.T) {
	sock := NewZsock(SUB)
	expected := "127.0.0.1"
	sock.SetTcpAcceptFilter(expected)
	val := sock.TcpAcceptFilter()
	if val != expected && val != "" {
		t.Errorf("TcpAcceptFilter returned %s should be %s", val, expected)
	}
        sock.Destroy()
}
