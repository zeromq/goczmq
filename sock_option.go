//go:generate gsl sockopts.xml
package goczmq

/*  =========================================================================
    zsock_option - get/set 0MQ socket options

            ****************************************************
            *   GENERATED SOURCE CODE, DO NOT EDIT!!           *
            *   TO CHANGE THIS, EDIT sockopts.gsl              *
            *   AND RUN gsl -q sockopts.xml                    *
            ****************************************************

    Copyright (c) the Contributors as noted in the AUTHORS file.
    This file is part of goczmq, the high-level go binding for CZMQ:
    http://github.com/zeromq/goczmq

    This Source Code Form is subject to the terms of the Mozilla Public
    License, v. 2.0. If a copy of the MPL was not distributed with this
    file, You can obtain one at http://mozilla.org/MPL/2.0/.
    =========================================================================
*/

/*
#include "czmq.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"unsafe"
)

// SetHeartbeatIvl sets the heartbeat_ivl option for the socket
func (s *Sock) SetHeartbeatIvl(val int) {
	C.zsock_set_heartbeat_ivl(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetHeartbeatIvl sets the heartbeat_ivl option for the socket
func SockSetHeartbeatIvl(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_heartbeat_ivl(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// HeartbeatIvl returns the current value of the socket's heartbeat_ivl option
func (s *Sock) HeartbeatIvl() int {
	val := C.zsock_heartbeat_ivl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// HeartbeatIvl returns the current value of the socket's heartbeat_ivl option
func HeartbeatIvl(s *Sock) int {
	val := C.zsock_heartbeat_ivl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetHeartbeatTtl sets the heartbeat_ttl option for the socket
func (s *Sock) SetHeartbeatTtl(val int) {
	C.zsock_set_heartbeat_ttl(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetHeartbeatTtl sets the heartbeat_ttl option for the socket
func SockSetHeartbeatTtl(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_heartbeat_ttl(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// HeartbeatTtl returns the current value of the socket's heartbeat_ttl option
func (s *Sock) HeartbeatTtl() int {
	val := C.zsock_heartbeat_ttl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// HeartbeatTtl returns the current value of the socket's heartbeat_ttl option
func HeartbeatTtl(s *Sock) int {
	val := C.zsock_heartbeat_ttl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetHeartbeatTimeout sets the heartbeat_timeout option for the socket
func (s *Sock) SetHeartbeatTimeout(val int) {
	C.zsock_set_heartbeat_timeout(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetHeartbeatTimeout sets the heartbeat_timeout option for the socket
func SockSetHeartbeatTimeout(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_heartbeat_timeout(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// HeartbeatTimeout returns the current value of the socket's heartbeat_timeout option
func (s *Sock) HeartbeatTimeout() int {
	val := C.zsock_heartbeat_timeout(unsafe.Pointer(s.zsockT))
	return int(val)
}

// HeartbeatTimeout returns the current value of the socket's heartbeat_timeout option
func HeartbeatTimeout(s *Sock) int {
	val := C.zsock_heartbeat_timeout(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetUseFd sets the use_fd option for the socket
func (s *Sock) SetUseFd(val int) {
	C.zsock_set_use_fd(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetUseFd sets the use_fd option for the socket
func SockSetUseFd(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_use_fd(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// UseFd returns the current value of the socket's use_fd option
func (s *Sock) UseFd() int {
	val := C.zsock_use_fd(unsafe.Pointer(s.zsockT))
	return int(val)
}

// UseFd returns the current value of the socket's use_fd option
func UseFd(s *Sock) int {
	val := C.zsock_use_fd(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetXPubManual sets the xpub_manual option for the socket
func (s *Sock) SetXPubManual(val int) {
	C.zsock_set_xpub_manual(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetXPubManual sets the xpub_manual option for the socket
func SockSetXPubManual(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_xpub_manual(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SetXPubWelcomeMsg sets the xpub_welcome_msg option for the socket
func (s *Sock) SetXPubWelcomeMsg(val string) {
    cVal := C.CString(val)
    defer C.free(unsafe.Pointer(cVal))

	C.zsock_set_xpub_welcome_msg(unsafe.Pointer(s.zsockT), cVal)
}

// SockSetXPubWelcomeMsg sets the xpub_welcome_msg option for the socket
func SockSetXPubWelcomeMsg(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_xpub_welcome_msg(unsafe.Pointer(s.zsockT), cV)
	}
}

// SetStreamNotify sets the stream_notify option for the socket
func (s *Sock) SetStreamNotify(val int) {
	C.zsock_set_stream_notify(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetStreamNotify sets the stream_notify option for the socket
func SockSetStreamNotify(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_stream_notify(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SetInvertMatching sets the invert_matching option for the socket
func (s *Sock) SetInvertMatching(val int) {
	C.zsock_set_invert_matching(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetInvertMatching sets the invert_matching option for the socket
func SockSetInvertMatching(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_invert_matching(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// InvertMatching returns the current value of the socket's invert_matching option
func (s *Sock) InvertMatching() int {
	val := C.zsock_invert_matching(unsafe.Pointer(s.zsockT))
	return int(val)
}

// InvertMatching returns the current value of the socket's invert_matching option
func InvertMatching(s *Sock) int {
	val := C.zsock_invert_matching(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetXPubVerboser sets the xpub_verboser option for the socket
func (s *Sock) SetXPubVerboser(val int) {
	C.zsock_set_xpub_verboser(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetXPubVerboser sets the xpub_verboser option for the socket
func SockSetXPubVerboser(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_xpub_verboser(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SetConnectTimeout sets the connect_timeout option for the socket
func (s *Sock) SetConnectTimeout(val int) {
	C.zsock_set_connect_timeout(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetConnectTimeout sets the connect_timeout option for the socket
func SockSetConnectTimeout(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_connect_timeout(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// ConnectTimeout returns the current value of the socket's connect_timeout option
func (s *Sock) ConnectTimeout() int {
	val := C.zsock_connect_timeout(unsafe.Pointer(s.zsockT))
	return int(val)
}

// ConnectTimeout returns the current value of the socket's connect_timeout option
func ConnectTimeout(s *Sock) int {
	val := C.zsock_connect_timeout(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetTcpMaxrt sets the tcp_maxrt option for the socket
func (s *Sock) SetTcpMaxrt(val int) {
	C.zsock_set_tcp_maxrt(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetTcpMaxrt sets the tcp_maxrt option for the socket
func SockSetTcpMaxrt(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_tcp_maxrt(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// TcpMaxrt returns the current value of the socket's tcp_maxrt option
func (s *Sock) TcpMaxrt() int {
	val := C.zsock_tcp_maxrt(unsafe.Pointer(s.zsockT))
	return int(val)
}

// TcpMaxrt returns the current value of the socket's tcp_maxrt option
func TcpMaxrt(s *Sock) int {
	val := C.zsock_tcp_maxrt(unsafe.Pointer(s.zsockT))
	return int(val)
}

// ThreadSafe returns the current value of the socket's thread_safe option
func (s *Sock) ThreadSafe() int {
	val := C.zsock_thread_safe(unsafe.Pointer(s.zsockT))
	return int(val)
}

// ThreadSafe returns the current value of the socket's thread_safe option
func ThreadSafe(s *Sock) int {
	val := C.zsock_thread_safe(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetMulticastMaxtpdu sets the multicast_maxtpdu option for the socket
func (s *Sock) SetMulticastMaxtpdu(val int) {
	C.zsock_set_multicast_maxtpdu(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetMulticastMaxtpdu sets the multicast_maxtpdu option for the socket
func SockSetMulticastMaxtpdu(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_multicast_maxtpdu(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// MulticastMaxtpdu returns the current value of the socket's multicast_maxtpdu option
func (s *Sock) MulticastMaxtpdu() int {
	val := C.zsock_multicast_maxtpdu(unsafe.Pointer(s.zsockT))
	return int(val)
}

// MulticastMaxtpdu returns the current value of the socket's multicast_maxtpdu option
func MulticastMaxtpdu(s *Sock) int {
	val := C.zsock_multicast_maxtpdu(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetVmciBufferSize sets the vmci_buffer_size option for the socket
func (s *Sock) SetVmciBufferSize(val int) {
	C.zsock_set_vmci_buffer_size(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetVmciBufferSize sets the vmci_buffer_size option for the socket
func SockSetVmciBufferSize(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_vmci_buffer_size(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// VmciBufferSize returns the current value of the socket's vmci_buffer_size option
func (s *Sock) VmciBufferSize() int {
	val := C.zsock_vmci_buffer_size(unsafe.Pointer(s.zsockT))
	return int(val)
}

// VmciBufferSize returns the current value of the socket's vmci_buffer_size option
func VmciBufferSize(s *Sock) int {
	val := C.zsock_vmci_buffer_size(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetVmciBufferMinSize sets the vmci_buffer_min_size option for the socket
func (s *Sock) SetVmciBufferMinSize(val int) {
	C.zsock_set_vmci_buffer_min_size(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetVmciBufferMinSize sets the vmci_buffer_min_size option for the socket
func SockSetVmciBufferMinSize(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_vmci_buffer_min_size(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// VmciBufferMinSize returns the current value of the socket's vmci_buffer_min_size option
func (s *Sock) VmciBufferMinSize() int {
	val := C.zsock_vmci_buffer_min_size(unsafe.Pointer(s.zsockT))
	return int(val)
}

// VmciBufferMinSize returns the current value of the socket's vmci_buffer_min_size option
func VmciBufferMinSize(s *Sock) int {
	val := C.zsock_vmci_buffer_min_size(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetVmciBufferMaxSize sets the vmci_buffer_max_size option for the socket
func (s *Sock) SetVmciBufferMaxSize(val int) {
	C.zsock_set_vmci_buffer_max_size(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetVmciBufferMaxSize sets the vmci_buffer_max_size option for the socket
func SockSetVmciBufferMaxSize(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_vmci_buffer_max_size(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// VmciBufferMaxSize returns the current value of the socket's vmci_buffer_max_size option
func (s *Sock) VmciBufferMaxSize() int {
	val := C.zsock_vmci_buffer_max_size(unsafe.Pointer(s.zsockT))
	return int(val)
}

// VmciBufferMaxSize returns the current value of the socket's vmci_buffer_max_size option
func VmciBufferMaxSize(s *Sock) int {
	val := C.zsock_vmci_buffer_max_size(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetVmciConnectTimeout sets the vmci_connect_timeout option for the socket
func (s *Sock) SetVmciConnectTimeout(val int) {
	C.zsock_set_vmci_connect_timeout(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetVmciConnectTimeout sets the vmci_connect_timeout option for the socket
func SockSetVmciConnectTimeout(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_vmci_connect_timeout(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// VmciConnectTimeout returns the current value of the socket's vmci_connect_timeout option
func (s *Sock) VmciConnectTimeout() int {
	val := C.zsock_vmci_connect_timeout(unsafe.Pointer(s.zsockT))
	return int(val)
}

// VmciConnectTimeout returns the current value of the socket's vmci_connect_timeout option
func VmciConnectTimeout(s *Sock) int {
	val := C.zsock_vmci_connect_timeout(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetConnectRid sets the connect_rid option for the socket
func (s *Sock) SetConnectRid(val string) {
    cVal := C.CString(val)
    defer C.free(unsafe.Pointer(cVal))

	C.zsock_set_connect_rid(unsafe.Pointer(s.zsockT), cVal)
}

// SockSetConnectRid sets the connect_rid option for the socket
func SockSetConnectRid(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_connect_rid(unsafe.Pointer(s.zsockT), cV)
	}
}

// SetHandshakeIvl sets the handshake_ivl option for the socket
func (s *Sock) SetHandshakeIvl(val int) {
	C.zsock_set_handshake_ivl(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetHandshakeIvl sets the handshake_ivl option for the socket
func SockSetHandshakeIvl(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_handshake_ivl(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// HandshakeIvl returns the current value of the socket's handshake_ivl option
func (s *Sock) HandshakeIvl() int {
	val := C.zsock_handshake_ivl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// HandshakeIvl returns the current value of the socket's handshake_ivl option
func HandshakeIvl(s *Sock) int {
	val := C.zsock_handshake_ivl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetSocksProxy sets the socks_proxy option for the socket
func (s *Sock) SetSocksProxy(val string) {
    cVal := C.CString(val)
    defer C.free(unsafe.Pointer(cVal))

	C.zsock_set_socks_proxy(unsafe.Pointer(s.zsockT), cVal)
}

// SockSetSocksProxy sets the socks_proxy option for the socket
func SockSetSocksProxy(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_socks_proxy(unsafe.Pointer(s.zsockT), cV)
	}
}

// SocksProxy returns the current value of the socket's socks_proxy option
func (s *Sock) SocksProxy() string {
	val := C.zsock_socks_proxy(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SocksProxy returns the current value of the socket's socks_proxy option
func SocksProxy(s *Sock) string {
	val := C.zsock_socks_proxy(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetXPubNodrop sets the xpub_nodrop option for the socket
func (s *Sock) SetXPubNodrop(val int) {
	C.zsock_set_xpub_nodrop(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetXPubNodrop sets the xpub_nodrop option for the socket
func SockSetXPubNodrop(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_xpub_nodrop(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SetTos sets the tos option for the socket
func (s *Sock) SetTos(val int) {
	C.zsock_set_tos(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetTos sets the tos option for the socket
func SockSetTos(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_tos(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Tos returns the current value of the socket's tos option
func (s *Sock) Tos() int {
	val := C.zsock_tos(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Tos returns the current value of the socket's tos option
func Tos(s *Sock) int {
	val := C.zsock_tos(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetRouterHandover sets the router_handover option for the socket
func (s *Sock) SetRouterHandover(val int) {
	C.zsock_set_router_handover(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetRouterHandover sets the router_handover option for the socket
func SockSetRouterHandover(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_router_handover(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SetRouterMandatory sets the router_mandatory option for the socket
func (s *Sock) SetRouterMandatory(val int) {
	C.zsock_set_router_mandatory(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetRouterMandatory sets the router_mandatory option for the socket
func SockSetRouterMandatory(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_router_mandatory(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SetProbeRouter sets the probe_router option for the socket
func (s *Sock) SetProbeRouter(val int) {
	C.zsock_set_probe_router(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetProbeRouter sets the probe_router option for the socket
func SockSetProbeRouter(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_probe_router(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SetReqRelaxed sets the req_relaxed option for the socket
func (s *Sock) SetReqRelaxed(val int) {
	C.zsock_set_req_relaxed(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetReqRelaxed sets the req_relaxed option for the socket
func SockSetReqRelaxed(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_req_relaxed(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SetReqCorrelate sets the req_correlate option for the socket
func (s *Sock) SetReqCorrelate(val int) {
	C.zsock_set_req_correlate(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetReqCorrelate sets the req_correlate option for the socket
func SockSetReqCorrelate(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_req_correlate(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SetConflate sets the conflate option for the socket
func (s *Sock) SetConflate(val int) {
	C.zsock_set_conflate(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetConflate sets the conflate option for the socket
func SockSetConflate(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_conflate(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SetZapDomain sets the zap_domain option for the socket
func (s *Sock) SetZapDomain(val string) {
    cVal := C.CString(val)
    defer C.free(unsafe.Pointer(cVal))

	C.zsock_set_zap_domain(unsafe.Pointer(s.zsockT), cVal)
}

// SockSetZapDomain sets the zap_domain option for the socket
func SockSetZapDomain(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_zap_domain(unsafe.Pointer(s.zsockT), cV)
	}
}

// ZapDomain returns the current value of the socket's zap_domain option
func (s *Sock) ZapDomain() string {
	val := C.zsock_zap_domain(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// ZapDomain returns the current value of the socket's zap_domain option
func ZapDomain(s *Sock) string {
	val := C.zsock_zap_domain(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// Mechanism returns the current value of the socket's mechanism option
func (s *Sock) Mechanism() int {
	val := C.zsock_mechanism(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Mechanism returns the current value of the socket's mechanism option
func Mechanism(s *Sock) int {
	val := C.zsock_mechanism(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetPlainServer sets the plain_server option for the socket
func (s *Sock) SetPlainServer(val int) {
	C.zsock_set_plain_server(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetPlainServer sets the plain_server option for the socket
func SockSetPlainServer(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_plain_server(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// PlainServer returns the current value of the socket's plain_server option
func (s *Sock) PlainServer() int {
	val := C.zsock_plain_server(unsafe.Pointer(s.zsockT))
	return int(val)
}

// PlainServer returns the current value of the socket's plain_server option
func PlainServer(s *Sock) int {
	val := C.zsock_plain_server(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetPlainUsername sets the plain_username option for the socket
func (s *Sock) SetPlainUsername(val string) {
    cVal := C.CString(val)
    defer C.free(unsafe.Pointer(cVal))

	C.zsock_set_plain_username(unsafe.Pointer(s.zsockT), cVal)
}

// SockSetPlainUsername sets the plain_username option for the socket
func SockSetPlainUsername(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_plain_username(unsafe.Pointer(s.zsockT), cV)
	}
}

// PlainUsername returns the current value of the socket's plain_username option
func (s *Sock) PlainUsername() string {
	val := C.zsock_plain_username(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// PlainUsername returns the current value of the socket's plain_username option
func PlainUsername(s *Sock) string {
	val := C.zsock_plain_username(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetPlainPassword sets the plain_password option for the socket
func (s *Sock) SetPlainPassword(val string) {
    cVal := C.CString(val)
    defer C.free(unsafe.Pointer(cVal))

	C.zsock_set_plain_password(unsafe.Pointer(s.zsockT), cVal)
}

// SockSetPlainPassword sets the plain_password option for the socket
func SockSetPlainPassword(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_plain_password(unsafe.Pointer(s.zsockT), cV)
	}
}

// PlainPassword returns the current value of the socket's plain_password option
func (s *Sock) PlainPassword() string {
	val := C.zsock_plain_password(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// PlainPassword returns the current value of the socket's plain_password option
func PlainPassword(s *Sock) string {
	val := C.zsock_plain_password(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetCurveServer sets the curve_server option for the socket
func (s *Sock) SetCurveServer(val int) {
	C.zsock_set_curve_server(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetCurveServer sets the curve_server option for the socket
func SockSetCurveServer(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_curve_server(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// CurveServer returns the current value of the socket's curve_server option
func (s *Sock) CurveServer() int {
	val := C.zsock_curve_server(unsafe.Pointer(s.zsockT))
	return int(val)
}

// CurveServer returns the current value of the socket's curve_server option
func CurveServer(s *Sock) int {
	val := C.zsock_curve_server(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetCurvePublickey sets the curve_publickey option for the socket
func (s *Sock) SetCurvePublickey(val string) {
    cVal := C.CString(val)
    defer C.free(unsafe.Pointer(cVal))

	C.zsock_set_curve_publickey(unsafe.Pointer(s.zsockT), cVal)
}

// SockSetCurvePublickey sets the curve_publickey option for the socket
func SockSetCurvePublickey(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_curve_publickey(unsafe.Pointer(s.zsockT), cV)
	}
}

// CurvePublickey returns the current value of the socket's curve_publickey option
func (s *Sock) CurvePublickey() string {
	val := C.zsock_curve_publickey(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// CurvePublickey returns the current value of the socket's curve_publickey option
func CurvePublickey(s *Sock) string {
	val := C.zsock_curve_publickey(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetCurveSecretkey sets the curve_secretkey option for the socket
func (s *Sock) SetCurveSecretkey(val string) {
    cVal := C.CString(val)
    defer C.free(unsafe.Pointer(cVal))

	C.zsock_set_curve_secretkey(unsafe.Pointer(s.zsockT), cVal)
}

// SockSetCurveSecretkey sets the curve_secretkey option for the socket
func SockSetCurveSecretkey(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_curve_secretkey(unsafe.Pointer(s.zsockT), cV)
	}
}

// CurveSecretkey returns the current value of the socket's curve_secretkey option
func (s *Sock) CurveSecretkey() string {
	val := C.zsock_curve_secretkey(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// CurveSecretkey returns the current value of the socket's curve_secretkey option
func CurveSecretkey(s *Sock) string {
	val := C.zsock_curve_secretkey(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetCurveServerkey sets the curve_serverkey option for the socket
func (s *Sock) SetCurveServerkey(val string) {
    cVal := C.CString(val)
    defer C.free(unsafe.Pointer(cVal))

	C.zsock_set_curve_serverkey(unsafe.Pointer(s.zsockT), cVal)
}

// SockSetCurveServerkey sets the curve_serverkey option for the socket
func SockSetCurveServerkey(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_curve_serverkey(unsafe.Pointer(s.zsockT), cV)
	}
}

// CurveServerkey returns the current value of the socket's curve_serverkey option
func (s *Sock) CurveServerkey() string {
	val := C.zsock_curve_serverkey(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// CurveServerkey returns the current value of the socket's curve_serverkey option
func CurveServerkey(s *Sock) string {
	val := C.zsock_curve_serverkey(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetGssapiServer sets the gssapi_server option for the socket
func (s *Sock) SetGssapiServer(val int) {
	C.zsock_set_gssapi_server(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetGssapiServer sets the gssapi_server option for the socket
func SockSetGssapiServer(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_gssapi_server(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// GssapiServer returns the current value of the socket's gssapi_server option
func (s *Sock) GssapiServer() int {
	val := C.zsock_gssapi_server(unsafe.Pointer(s.zsockT))
	return int(val)
}

// GssapiServer returns the current value of the socket's gssapi_server option
func GssapiServer(s *Sock) int {
	val := C.zsock_gssapi_server(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetGssapiPlaintext sets the gssapi_plaintext option for the socket
func (s *Sock) SetGssapiPlaintext(val int) {
	C.zsock_set_gssapi_plaintext(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetGssapiPlaintext sets the gssapi_plaintext option for the socket
func SockSetGssapiPlaintext(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_gssapi_plaintext(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// GssapiPlaintext returns the current value of the socket's gssapi_plaintext option
func (s *Sock) GssapiPlaintext() int {
	val := C.zsock_gssapi_plaintext(unsafe.Pointer(s.zsockT))
	return int(val)
}

// GssapiPlaintext returns the current value of the socket's gssapi_plaintext option
func GssapiPlaintext(s *Sock) int {
	val := C.zsock_gssapi_plaintext(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetGssapiPrincipal sets the gssapi_principal option for the socket
func (s *Sock) SetGssapiPrincipal(val string) {
    cVal := C.CString(val)
    defer C.free(unsafe.Pointer(cVal))

	C.zsock_set_gssapi_principal(unsafe.Pointer(s.zsockT), cVal)
}

// SockSetGssapiPrincipal sets the gssapi_principal option for the socket
func SockSetGssapiPrincipal(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_gssapi_principal(unsafe.Pointer(s.zsockT), cV)
	}
}

// GssapiPrincipal returns the current value of the socket's gssapi_principal option
func (s *Sock) GssapiPrincipal() string {
	val := C.zsock_gssapi_principal(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// GssapiPrincipal returns the current value of the socket's gssapi_principal option
func GssapiPrincipal(s *Sock) string {
	val := C.zsock_gssapi_principal(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetGssapiServicePrincipal sets the gssapi_service_principal option for the socket
func (s *Sock) SetGssapiServicePrincipal(val string) {
    cVal := C.CString(val)
    defer C.free(unsafe.Pointer(cVal))

	C.zsock_set_gssapi_service_principal(unsafe.Pointer(s.zsockT), cVal)
}

// SockSetGssapiServicePrincipal sets the gssapi_service_principal option for the socket
func SockSetGssapiServicePrincipal(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_gssapi_service_principal(unsafe.Pointer(s.zsockT), cV)
	}
}

// GssapiServicePrincipal returns the current value of the socket's gssapi_service_principal option
func (s *Sock) GssapiServicePrincipal() string {
	val := C.zsock_gssapi_service_principal(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// GssapiServicePrincipal returns the current value of the socket's gssapi_service_principal option
func GssapiServicePrincipal(s *Sock) string {
	val := C.zsock_gssapi_service_principal(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetIpv6 sets the ipv6 option for the socket
func (s *Sock) SetIpv6(val int) {
	C.zsock_set_ipv6(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetIpv6 sets the ipv6 option for the socket
func SockSetIpv6(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_ipv6(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Ipv6 returns the current value of the socket's ipv6 option
func (s *Sock) Ipv6() int {
	val := C.zsock_ipv6(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Ipv6 returns the current value of the socket's ipv6 option
func Ipv6(s *Sock) int {
	val := C.zsock_ipv6(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetImmediate sets the immediate option for the socket
func (s *Sock) SetImmediate(val int) {
	C.zsock_set_immediate(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetImmediate sets the immediate option for the socket
func SockSetImmediate(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_immediate(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Immediate returns the current value of the socket's immediate option
func (s *Sock) Immediate() int {
	val := C.zsock_immediate(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Immediate returns the current value of the socket's immediate option
func Immediate(s *Sock) int {
	val := C.zsock_immediate(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetRouterRaw sets the router_raw option for the socket
func (s *Sock) SetRouterRaw(val int) {
	C.zsock_set_router_raw(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetRouterRaw sets the router_raw option for the socket
func SockSetRouterRaw(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_router_raw(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SetIpv4only sets the ipv4only option for the socket
func (s *Sock) SetIpv4only(val int) {
	C.zsock_set_ipv4only(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetIpv4only sets the ipv4only option for the socket
func SockSetIpv4only(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_ipv4only(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Ipv4only returns the current value of the socket's ipv4only option
func (s *Sock) Ipv4only() int {
	val := C.zsock_ipv4only(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Ipv4only returns the current value of the socket's ipv4only option
func Ipv4only(s *Sock) int {
	val := C.zsock_ipv4only(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetDelayAttachOnConnect sets the delay_attach_on_connect option for the socket
func (s *Sock) SetDelayAttachOnConnect(val int) {
	C.zsock_set_delay_attach_on_connect(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetDelayAttachOnConnect sets the delay_attach_on_connect option for the socket
func SockSetDelayAttachOnConnect(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_delay_attach_on_connect(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Type returns the current value of the socket's type option
func (s *Sock) Type() int {
	val := C.zsock_type(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Type returns the current value of the socket's type option
func Type(s *Sock) int {
	val := C.zsock_type(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetSndhwm sets the sndhwm option for the socket
func (s *Sock) SetSndhwm(val int) {
	C.zsock_set_sndhwm(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetSndhwm sets the sndhwm option for the socket
func SockSetSndhwm(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_sndhwm(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Sndhwm returns the current value of the socket's sndhwm option
func (s *Sock) Sndhwm() int {
	val := C.zsock_sndhwm(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Sndhwm returns the current value of the socket's sndhwm option
func Sndhwm(s *Sock) int {
	val := C.zsock_sndhwm(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetRcvhwm sets the rcvhwm option for the socket
func (s *Sock) SetRcvhwm(val int) {
	C.zsock_set_rcvhwm(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetRcvhwm sets the rcvhwm option for the socket
func SockSetRcvhwm(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_rcvhwm(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Rcvhwm returns the current value of the socket's rcvhwm option
func (s *Sock) Rcvhwm() int {
	val := C.zsock_rcvhwm(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Rcvhwm returns the current value of the socket's rcvhwm option
func Rcvhwm(s *Sock) int {
	val := C.zsock_rcvhwm(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetAffinity sets the affinity option for the socket
func (s *Sock) SetAffinity(val int) {
	C.zsock_set_affinity(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetAffinity sets the affinity option for the socket
func SockSetAffinity(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_affinity(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Affinity returns the current value of the socket's affinity option
func (s *Sock) Affinity() int {
	val := C.zsock_affinity(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Affinity returns the current value of the socket's affinity option
func Affinity(s *Sock) int {
	val := C.zsock_affinity(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetSubscribe sets the subscribe option for the socket
func (s *Sock) SetSubscribe(val string) {
    cVal := C.CString(val)
    defer C.free(unsafe.Pointer(cVal))

	C.zsock_set_subscribe(unsafe.Pointer(s.zsockT), cVal)
}

// SockSetSubscribe sets the subscribe option for the socket
func SockSetSubscribe(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_subscribe(unsafe.Pointer(s.zsockT), cV)
	}
}

// SetUnsubscribe sets the unsubscribe option for the socket
func (s *Sock) SetUnsubscribe(val string) {
    cVal := C.CString(val)
    defer C.free(unsafe.Pointer(cVal))

	C.zsock_set_unsubscribe(unsafe.Pointer(s.zsockT), cVal)
}

// SockSetUnsubscribe sets the unsubscribe option for the socket
func SockSetUnsubscribe(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_unsubscribe(unsafe.Pointer(s.zsockT), cV)
	}
}

// SetIdentity sets the identity option for the socket
func (s *Sock) SetIdentity(val string) {
    cVal := C.CString(val)
    defer C.free(unsafe.Pointer(cVal))

	C.zsock_set_identity(unsafe.Pointer(s.zsockT), cVal)
}

// SockSetIdentity sets the identity option for the socket
func SockSetIdentity(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_identity(unsafe.Pointer(s.zsockT), cV)
	}
}

// Identity returns the current value of the socket's identity option
func (s *Sock) Identity() string {
	val := C.zsock_identity(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// Identity returns the current value of the socket's identity option
func Identity(s *Sock) string {
	val := C.zsock_identity(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetRate sets the rate option for the socket
func (s *Sock) SetRate(val int) {
	C.zsock_set_rate(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetRate sets the rate option for the socket
func SockSetRate(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_rate(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Rate returns the current value of the socket's rate option
func (s *Sock) Rate() int {
	val := C.zsock_rate(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Rate returns the current value of the socket's rate option
func Rate(s *Sock) int {
	val := C.zsock_rate(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetRecoveryIvl sets the recovery_ivl option for the socket
func (s *Sock) SetRecoveryIvl(val int) {
	C.zsock_set_recovery_ivl(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetRecoveryIvl sets the recovery_ivl option for the socket
func SockSetRecoveryIvl(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_recovery_ivl(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// RecoveryIvl returns the current value of the socket's recovery_ivl option
func (s *Sock) RecoveryIvl() int {
	val := C.zsock_recovery_ivl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// RecoveryIvl returns the current value of the socket's recovery_ivl option
func RecoveryIvl(s *Sock) int {
	val := C.zsock_recovery_ivl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetSndbuf sets the sndbuf option for the socket
func (s *Sock) SetSndbuf(val int) {
	C.zsock_set_sndbuf(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetSndbuf sets the sndbuf option for the socket
func SockSetSndbuf(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_sndbuf(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Sndbuf returns the current value of the socket's sndbuf option
func (s *Sock) Sndbuf() int {
	val := C.zsock_sndbuf(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Sndbuf returns the current value of the socket's sndbuf option
func Sndbuf(s *Sock) int {
	val := C.zsock_sndbuf(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetRcvbuf sets the rcvbuf option for the socket
func (s *Sock) SetRcvbuf(val int) {
	C.zsock_set_rcvbuf(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetRcvbuf sets the rcvbuf option for the socket
func SockSetRcvbuf(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_rcvbuf(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Rcvbuf returns the current value of the socket's rcvbuf option
func (s *Sock) Rcvbuf() int {
	val := C.zsock_rcvbuf(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Rcvbuf returns the current value of the socket's rcvbuf option
func Rcvbuf(s *Sock) int {
	val := C.zsock_rcvbuf(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetLinger sets the linger option for the socket
func (s *Sock) SetLinger(val int) {
	C.zsock_set_linger(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetLinger sets the linger option for the socket
func SockSetLinger(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_linger(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Linger returns the current value of the socket's linger option
func (s *Sock) Linger() int {
	val := C.zsock_linger(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Linger returns the current value of the socket's linger option
func Linger(s *Sock) int {
	val := C.zsock_linger(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetReconnectIvl sets the reconnect_ivl option for the socket
func (s *Sock) SetReconnectIvl(val int) {
	C.zsock_set_reconnect_ivl(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetReconnectIvl sets the reconnect_ivl option for the socket
func SockSetReconnectIvl(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_reconnect_ivl(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// ReconnectIvl returns the current value of the socket's reconnect_ivl option
func (s *Sock) ReconnectIvl() int {
	val := C.zsock_reconnect_ivl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// ReconnectIvl returns the current value of the socket's reconnect_ivl option
func ReconnectIvl(s *Sock) int {
	val := C.zsock_reconnect_ivl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetReconnectIvlMax sets the reconnect_ivl_max option for the socket
func (s *Sock) SetReconnectIvlMax(val int) {
	C.zsock_set_reconnect_ivl_max(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetReconnectIvlMax sets the reconnect_ivl_max option for the socket
func SockSetReconnectIvlMax(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_reconnect_ivl_max(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// ReconnectIvlMax returns the current value of the socket's reconnect_ivl_max option
func (s *Sock) ReconnectIvlMax() int {
	val := C.zsock_reconnect_ivl_max(unsafe.Pointer(s.zsockT))
	return int(val)
}

// ReconnectIvlMax returns the current value of the socket's reconnect_ivl_max option
func ReconnectIvlMax(s *Sock) int {
	val := C.zsock_reconnect_ivl_max(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetBacklog sets the backlog option for the socket
func (s *Sock) SetBacklog(val int) {
	C.zsock_set_backlog(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetBacklog sets the backlog option for the socket
func SockSetBacklog(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_backlog(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Backlog returns the current value of the socket's backlog option
func (s *Sock) Backlog() int {
	val := C.zsock_backlog(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Backlog returns the current value of the socket's backlog option
func Backlog(s *Sock) int {
	val := C.zsock_backlog(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetMaxmsgsize sets the maxmsgsize option for the socket
func (s *Sock) SetMaxmsgsize(val int) {
	C.zsock_set_maxmsgsize(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetMaxmsgsize sets the maxmsgsize option for the socket
func SockSetMaxmsgsize(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_maxmsgsize(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Maxmsgsize returns the current value of the socket's maxmsgsize option
func (s *Sock) Maxmsgsize() int {
	val := C.zsock_maxmsgsize(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Maxmsgsize returns the current value of the socket's maxmsgsize option
func Maxmsgsize(s *Sock) int {
	val := C.zsock_maxmsgsize(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetMulticastHops sets the multicast_hops option for the socket
func (s *Sock) SetMulticastHops(val int) {
	C.zsock_set_multicast_hops(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetMulticastHops sets the multicast_hops option for the socket
func SockSetMulticastHops(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_multicast_hops(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// MulticastHops returns the current value of the socket's multicast_hops option
func (s *Sock) MulticastHops() int {
	val := C.zsock_multicast_hops(unsafe.Pointer(s.zsockT))
	return int(val)
}

// MulticastHops returns the current value of the socket's multicast_hops option
func MulticastHops(s *Sock) int {
	val := C.zsock_multicast_hops(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetRcvtimeo sets the rcvtimeo option for the socket
func (s *Sock) SetRcvtimeo(val int) {
	C.zsock_set_rcvtimeo(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetRcvtimeo sets the rcvtimeo option for the socket
func SockSetRcvtimeo(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_rcvtimeo(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Rcvtimeo returns the current value of the socket's rcvtimeo option
func (s *Sock) Rcvtimeo() int {
	val := C.zsock_rcvtimeo(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Rcvtimeo returns the current value of the socket's rcvtimeo option
func Rcvtimeo(s *Sock) int {
	val := C.zsock_rcvtimeo(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetSndtimeo sets the sndtimeo option for the socket
func (s *Sock) SetSndtimeo(val int) {
	C.zsock_set_sndtimeo(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetSndtimeo sets the sndtimeo option for the socket
func SockSetSndtimeo(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_sndtimeo(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Sndtimeo returns the current value of the socket's sndtimeo option
func (s *Sock) Sndtimeo() int {
	val := C.zsock_sndtimeo(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Sndtimeo returns the current value of the socket's sndtimeo option
func Sndtimeo(s *Sock) int {
	val := C.zsock_sndtimeo(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetXPubVerbose sets the xpub_verbose option for the socket
func (s *Sock) SetXPubVerbose(val int) {
	C.zsock_set_xpub_verbose(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetXPubVerbose sets the xpub_verbose option for the socket
func SockSetXPubVerbose(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_xpub_verbose(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SetTcpKeepalive sets the tcp_keepalive option for the socket
func (s *Sock) SetTcpKeepalive(val int) {
	C.zsock_set_tcp_keepalive(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetTcpKeepalive sets the tcp_keepalive option for the socket
func SockSetTcpKeepalive(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_tcp_keepalive(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// TcpKeepalive returns the current value of the socket's tcp_keepalive option
func (s *Sock) TcpKeepalive() int {
	val := C.zsock_tcp_keepalive(unsafe.Pointer(s.zsockT))
	return int(val)
}

// TcpKeepalive returns the current value of the socket's tcp_keepalive option
func TcpKeepalive(s *Sock) int {
	val := C.zsock_tcp_keepalive(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetTcpKeepaliveIdle sets the tcp_keepalive_idle option for the socket
func (s *Sock) SetTcpKeepaliveIdle(val int) {
	C.zsock_set_tcp_keepalive_idle(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetTcpKeepaliveIdle sets the tcp_keepalive_idle option for the socket
func SockSetTcpKeepaliveIdle(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_tcp_keepalive_idle(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// TcpKeepaliveIdle returns the current value of the socket's tcp_keepalive_idle option
func (s *Sock) TcpKeepaliveIdle() int {
	val := C.zsock_tcp_keepalive_idle(unsafe.Pointer(s.zsockT))
	return int(val)
}

// TcpKeepaliveIdle returns the current value of the socket's tcp_keepalive_idle option
func TcpKeepaliveIdle(s *Sock) int {
	val := C.zsock_tcp_keepalive_idle(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetTcpKeepaliveCnt sets the tcp_keepalive_cnt option for the socket
func (s *Sock) SetTcpKeepaliveCnt(val int) {
	C.zsock_set_tcp_keepalive_cnt(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetTcpKeepaliveCnt sets the tcp_keepalive_cnt option for the socket
func SockSetTcpKeepaliveCnt(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_tcp_keepalive_cnt(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// TcpKeepaliveCnt returns the current value of the socket's tcp_keepalive_cnt option
func (s *Sock) TcpKeepaliveCnt() int {
	val := C.zsock_tcp_keepalive_cnt(unsafe.Pointer(s.zsockT))
	return int(val)
}

// TcpKeepaliveCnt returns the current value of the socket's tcp_keepalive_cnt option
func TcpKeepaliveCnt(s *Sock) int {
	val := C.zsock_tcp_keepalive_cnt(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetTcpKeepaliveIntvl sets the tcp_keepalive_intvl option for the socket
func (s *Sock) SetTcpKeepaliveIntvl(val int) {
	C.zsock_set_tcp_keepalive_intvl(unsafe.Pointer(s.zsockT), C.int(val))
}

// SockSetTcpKeepaliveIntvl sets the tcp_keepalive_intvl option for the socket
func SockSetTcpKeepaliveIntvl(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_tcp_keepalive_intvl(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// TcpKeepaliveIntvl returns the current value of the socket's tcp_keepalive_intvl option
func (s *Sock) TcpKeepaliveIntvl() int {
	val := C.zsock_tcp_keepalive_intvl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// TcpKeepaliveIntvl returns the current value of the socket's tcp_keepalive_intvl option
func TcpKeepaliveIntvl(s *Sock) int {
	val := C.zsock_tcp_keepalive_intvl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetTcpAcceptFilter sets the tcp_accept_filter option for the socket
func (s *Sock) SetTcpAcceptFilter(val string) {
    cVal := C.CString(val)
    defer C.free(unsafe.Pointer(cVal))

	C.zsock_set_tcp_accept_filter(unsafe.Pointer(s.zsockT), cVal)
}

// SockSetTcpAcceptFilter sets the tcp_accept_filter option for the socket
func SockSetTcpAcceptFilter(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_tcp_accept_filter(unsafe.Pointer(s.zsockT), cV)
	}
}

// TcpAcceptFilter returns the current value of the socket's tcp_accept_filter option
func (s *Sock) TcpAcceptFilter() string {
	val := C.zsock_tcp_accept_filter(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// TcpAcceptFilter returns the current value of the socket's tcp_accept_filter option
func TcpAcceptFilter(s *Sock) string {
	val := C.zsock_tcp_accept_filter(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// Rcvmore returns the current value of the socket's rcvmore option
func (s *Sock) Rcvmore() int {
	val := C.zsock_rcvmore(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Rcvmore returns the current value of the socket's rcvmore option
func Rcvmore(s *Sock) int {
	val := C.zsock_rcvmore(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Fd returns the current value of the socket's fd option
func (s *Sock) Fd() int {
	val := C.zsock_fd(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Fd returns the current value of the socket's fd option
func Fd(s *Sock) int {
	val := C.zsock_fd(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Events returns the current value of the socket's events option
func (s *Sock) Events() int {
	val := C.zsock_events(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Events returns the current value of the socket's events option
func Events(s *Sock) int {
	val := C.zsock_events(unsafe.Pointer(s.zsockT))
	return int(val)
}

// LastEndpoint returns the current value of the socket's last_endpoint option
func (s *Sock) LastEndpoint() string {
	val := C.zsock_last_endpoint(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// LastEndpoint returns the current value of the socket's last_endpoint option
func LastEndpoint(s *Sock) string {
	val := C.zsock_last_endpoint(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

