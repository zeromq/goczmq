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

// SockSetHeartbeatIvl sets the heartbeat_ivl option for the socket
func SockSetHeartbeatIvl(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_heartbeat_ivl(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// HeartbeatIvl returns the current value of the socket's heartbeat_ivl option
func HeartbeatIvl(s *Sock) int {
	val := C.zsock_heartbeat_ivl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetHeartbeatTtl sets the heartbeat_ttl option for the socket
func SockSetHeartbeatTtl(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_heartbeat_ttl(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// HeartbeatTtl returns the current value of the socket's heartbeat_ttl option
func HeartbeatTtl(s *Sock) int {
	val := C.zsock_heartbeat_ttl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetHeartbeatTimeout sets the heartbeat_timeout option for the socket
func SockSetHeartbeatTimeout(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_heartbeat_timeout(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// HeartbeatTimeout returns the current value of the socket's heartbeat_timeout option
func HeartbeatTimeout(s *Sock) int {
	val := C.zsock_heartbeat_timeout(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetUseFd sets the use_fd option for the socket
func SockSetUseFd(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_use_fd(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// UseFd returns the current value of the socket's use_fd option
func UseFd(s *Sock) int {
	val := C.zsock_use_fd(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetXPubManual sets the xpub_manual option for the socket
func SockSetXPubManual(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_xpub_manual(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SockSetXPubWelcomeMsg sets the xpub_welcome_msg option for the socket
func SockSetXPubWelcomeMsg(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_xpub_welcome_msg(unsafe.Pointer(s.zsockT), cV)
	}
}

// SockSetStreamNotify sets the stream_notify option for the socket
func SockSetStreamNotify(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_stream_notify(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SockSetInvertMatching sets the invert_matching option for the socket
func SockSetInvertMatching(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_invert_matching(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// InvertMatching returns the current value of the socket's invert_matching option
func InvertMatching(s *Sock) int {
	val := C.zsock_invert_matching(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetXPubVerboser sets the xpub_verboser option for the socket
func SockSetXPubVerboser(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_xpub_verboser(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SockSetConnectTimeout sets the connect_timeout option for the socket
func SockSetConnectTimeout(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_connect_timeout(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// ConnectTimeout returns the current value of the socket's connect_timeout option
func ConnectTimeout(s *Sock) int {
	val := C.zsock_connect_timeout(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetTcpMaxrt sets the tcp_maxrt option for the socket
func SockSetTcpMaxrt(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_tcp_maxrt(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// TcpMaxrt returns the current value of the socket's tcp_maxrt option
func TcpMaxrt(s *Sock) int {
	val := C.zsock_tcp_maxrt(unsafe.Pointer(s.zsockT))
	return int(val)
}

// ThreadSafe returns the current value of the socket's thread_safe option
func ThreadSafe(s *Sock) int {
	val := C.zsock_thread_safe(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetMulticastMaxtpdu sets the multicast_maxtpdu option for the socket
func SockSetMulticastMaxtpdu(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_multicast_maxtpdu(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// MulticastMaxtpdu returns the current value of the socket's multicast_maxtpdu option
func MulticastMaxtpdu(s *Sock) int {
	val := C.zsock_multicast_maxtpdu(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetVmciBufferSize sets the vmci_buffer_size option for the socket
func SockSetVmciBufferSize(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_vmci_buffer_size(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// VmciBufferSize returns the current value of the socket's vmci_buffer_size option
func VmciBufferSize(s *Sock) int {
	val := C.zsock_vmci_buffer_size(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetVmciBufferMinSize sets the vmci_buffer_min_size option for the socket
func SockSetVmciBufferMinSize(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_vmci_buffer_min_size(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// VmciBufferMinSize returns the current value of the socket's vmci_buffer_min_size option
func VmciBufferMinSize(s *Sock) int {
	val := C.zsock_vmci_buffer_min_size(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetVmciBufferMaxSize sets the vmci_buffer_max_size option for the socket
func SockSetVmciBufferMaxSize(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_vmci_buffer_max_size(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// VmciBufferMaxSize returns the current value of the socket's vmci_buffer_max_size option
func VmciBufferMaxSize(s *Sock) int {
	val := C.zsock_vmci_buffer_max_size(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetVmciConnectTimeout sets the vmci_connect_timeout option for the socket
func SockSetVmciConnectTimeout(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_vmci_connect_timeout(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// VmciConnectTimeout returns the current value of the socket's vmci_connect_timeout option
func VmciConnectTimeout(s *Sock) int {
	val := C.zsock_vmci_connect_timeout(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetConnectRid sets the connect_rid option for the socket
func SockSetConnectRid(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_connect_rid(unsafe.Pointer(s.zsockT), cV)
	}
}

// SockSetHandshakeIvl sets the handshake_ivl option for the socket
func SockSetHandshakeIvl(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_handshake_ivl(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// HandshakeIvl returns the current value of the socket's handshake_ivl option
func HandshakeIvl(s *Sock) int {
	val := C.zsock_handshake_ivl(unsafe.Pointer(s.zsockT))
	return int(val)
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
func SocksProxy(s *Sock) string {
	val := C.zsock_socks_proxy(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SockSetXPubNodrop sets the xpub_nodrop option for the socket
func SockSetXPubNodrop(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_xpub_nodrop(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SockSetTos sets the tos option for the socket
func SockSetTos(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_tos(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Tos returns the current value of the socket's tos option
func Tos(s *Sock) int {
	val := C.zsock_tos(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetRouterHandover sets the router_handover option for the socket
func SockSetRouterHandover(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_router_handover(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SockSetRouterMandatory sets the router_mandatory option for the socket
func SockSetRouterMandatory(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_router_mandatory(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SockSetProbeRouter sets the probe_router option for the socket
func SockSetProbeRouter(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_probe_router(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SockSetReqRelaxed sets the req_relaxed option for the socket
func SockSetReqRelaxed(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_req_relaxed(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SockSetReqCorrelate sets the req_correlate option for the socket
func SockSetReqCorrelate(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_req_correlate(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SockSetConflate sets the conflate option for the socket
func SockSetConflate(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_conflate(unsafe.Pointer(s.zsockT), C.int(v))
	}
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
func ZapDomain(s *Sock) string {
	val := C.zsock_zap_domain(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// Mechanism returns the current value of the socket's mechanism option
func Mechanism(s *Sock) int {
	val := C.zsock_mechanism(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetPlainServer sets the plain_server option for the socket
func SockSetPlainServer(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_plain_server(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// PlainServer returns the current value of the socket's plain_server option
func PlainServer(s *Sock) int {
	val := C.zsock_plain_server(unsafe.Pointer(s.zsockT))
	return int(val)
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
func PlainUsername(s *Sock) string {
	val := C.zsock_plain_username(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
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
func PlainPassword(s *Sock) string {
	val := C.zsock_plain_password(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SockSetCurveServer sets the curve_server option for the socket
func SockSetCurveServer(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_curve_server(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// CurveServer returns the current value of the socket's curve_server option
func CurveServer(s *Sock) int {
	val := C.zsock_curve_server(unsafe.Pointer(s.zsockT))
	return int(val)
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
func CurvePublickey(s *Sock) string {
	val := C.zsock_curve_publickey(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
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
func CurveSecretkey(s *Sock) string {
	val := C.zsock_curve_secretkey(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
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
func CurveServerkey(s *Sock) string {
	val := C.zsock_curve_serverkey(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SockSetGssapiServer sets the gssapi_server option for the socket
func SockSetGssapiServer(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_gssapi_server(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// GssapiServer returns the current value of the socket's gssapi_server option
func GssapiServer(s *Sock) int {
	val := C.zsock_gssapi_server(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetGssapiPlaintext sets the gssapi_plaintext option for the socket
func SockSetGssapiPlaintext(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_gssapi_plaintext(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// GssapiPlaintext returns the current value of the socket's gssapi_plaintext option
func GssapiPlaintext(s *Sock) int {
	val := C.zsock_gssapi_plaintext(unsafe.Pointer(s.zsockT))
	return int(val)
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
func GssapiPrincipal(s *Sock) string {
	val := C.zsock_gssapi_principal(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
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
func GssapiServicePrincipal(s *Sock) string {
	val := C.zsock_gssapi_service_principal(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SockSetIpv6 sets the ipv6 option for the socket
func SockSetIpv6(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_ipv6(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Ipv6 returns the current value of the socket's ipv6 option
func Ipv6(s *Sock) int {
	val := C.zsock_ipv6(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetImmediate sets the immediate option for the socket
func SockSetImmediate(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_immediate(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Immediate returns the current value of the socket's immediate option
func Immediate(s *Sock) int {
	val := C.zsock_immediate(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetRouterRaw sets the router_raw option for the socket
func SockSetRouterRaw(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_router_raw(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SockSetIpv4only sets the ipv4only option for the socket
func SockSetIpv4only(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_ipv4only(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Ipv4only returns the current value of the socket's ipv4only option
func Ipv4only(s *Sock) int {
	val := C.zsock_ipv4only(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetDelayAttachOnConnect sets the delay_attach_on_connect option for the socket
func SockSetDelayAttachOnConnect(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_delay_attach_on_connect(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Type returns the current value of the socket's type option
func Type(s *Sock) int {
	val := C.zsock_type(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetSndhwm sets the sndhwm option for the socket
func SockSetSndhwm(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_sndhwm(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Sndhwm returns the current value of the socket's sndhwm option
func Sndhwm(s *Sock) int {
	val := C.zsock_sndhwm(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetRcvhwm sets the rcvhwm option for the socket
func SockSetRcvhwm(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_rcvhwm(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Rcvhwm returns the current value of the socket's rcvhwm option
func Rcvhwm(s *Sock) int {
	val := C.zsock_rcvhwm(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetAffinity sets the affinity option for the socket
func SockSetAffinity(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_affinity(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Affinity returns the current value of the socket's affinity option
func Affinity(s *Sock) int {
	val := C.zsock_affinity(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetSubscribe sets the subscribe option for the socket
func SockSetSubscribe(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_subscribe(unsafe.Pointer(s.zsockT), cV)
	}
}

// SockSetUnsubscribe sets the unsubscribe option for the socket
func SockSetUnsubscribe(v string) SockOption {
	return func(s *Sock) {
		cV := C.CString(v)
		defer C.free(unsafe.Pointer(cV))
		C.zsock_set_unsubscribe(unsafe.Pointer(s.zsockT), cV)
	}
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
func Identity(s *Sock) string {
	val := C.zsock_identity(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SockSetRate sets the rate option for the socket
func SockSetRate(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_rate(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Rate returns the current value of the socket's rate option
func Rate(s *Sock) int {
	val := C.zsock_rate(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetRecoveryIvl sets the recovery_ivl option for the socket
func SockSetRecoveryIvl(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_recovery_ivl(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// RecoveryIvl returns the current value of the socket's recovery_ivl option
func RecoveryIvl(s *Sock) int {
	val := C.zsock_recovery_ivl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetSndbuf sets the sndbuf option for the socket
func SockSetSndbuf(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_sndbuf(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Sndbuf returns the current value of the socket's sndbuf option
func Sndbuf(s *Sock) int {
	val := C.zsock_sndbuf(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetRcvbuf sets the rcvbuf option for the socket
func SockSetRcvbuf(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_rcvbuf(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Rcvbuf returns the current value of the socket's rcvbuf option
func Rcvbuf(s *Sock) int {
	val := C.zsock_rcvbuf(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetLinger sets the linger option for the socket
func SockSetLinger(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_linger(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Linger returns the current value of the socket's linger option
func Linger(s *Sock) int {
	val := C.zsock_linger(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetReconnectIvl sets the reconnect_ivl option for the socket
func SockSetReconnectIvl(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_reconnect_ivl(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// ReconnectIvl returns the current value of the socket's reconnect_ivl option
func ReconnectIvl(s *Sock) int {
	val := C.zsock_reconnect_ivl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetReconnectIvlMax sets the reconnect_ivl_max option for the socket
func SockSetReconnectIvlMax(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_reconnect_ivl_max(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// ReconnectIvlMax returns the current value of the socket's reconnect_ivl_max option
func ReconnectIvlMax(s *Sock) int {
	val := C.zsock_reconnect_ivl_max(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetBacklog sets the backlog option for the socket
func SockSetBacklog(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_backlog(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Backlog returns the current value of the socket's backlog option
func Backlog(s *Sock) int {
	val := C.zsock_backlog(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetMaxmsgsize sets the maxmsgsize option for the socket
func SockSetMaxmsgsize(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_maxmsgsize(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Maxmsgsize returns the current value of the socket's maxmsgsize option
func Maxmsgsize(s *Sock) int {
	val := C.zsock_maxmsgsize(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetMulticastHops sets the multicast_hops option for the socket
func SockSetMulticastHops(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_multicast_hops(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// MulticastHops returns the current value of the socket's multicast_hops option
func MulticastHops(s *Sock) int {
	val := C.zsock_multicast_hops(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetRcvtimeo sets the rcvtimeo option for the socket
func SockSetRcvtimeo(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_rcvtimeo(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Rcvtimeo returns the current value of the socket's rcvtimeo option
func Rcvtimeo(s *Sock) int {
	val := C.zsock_rcvtimeo(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetSndtimeo sets the sndtimeo option for the socket
func SockSetSndtimeo(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_sndtimeo(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// Sndtimeo returns the current value of the socket's sndtimeo option
func Sndtimeo(s *Sock) int {
	val := C.zsock_sndtimeo(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetXPubVerbose sets the xpub_verbose option for the socket
func SockSetXPubVerbose(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_xpub_verbose(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// SockSetTcpKeepalive sets the tcp_keepalive option for the socket
func SockSetTcpKeepalive(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_tcp_keepalive(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// TcpKeepalive returns the current value of the socket's tcp_keepalive option
func TcpKeepalive(s *Sock) int {
	val := C.zsock_tcp_keepalive(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetTcpKeepaliveIdle sets the tcp_keepalive_idle option for the socket
func SockSetTcpKeepaliveIdle(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_tcp_keepalive_idle(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// TcpKeepaliveIdle returns the current value of the socket's tcp_keepalive_idle option
func TcpKeepaliveIdle(s *Sock) int {
	val := C.zsock_tcp_keepalive_idle(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetTcpKeepaliveCnt sets the tcp_keepalive_cnt option for the socket
func SockSetTcpKeepaliveCnt(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_tcp_keepalive_cnt(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// TcpKeepaliveCnt returns the current value of the socket's tcp_keepalive_cnt option
func TcpKeepaliveCnt(s *Sock) int {
	val := C.zsock_tcp_keepalive_cnt(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SockSetTcpKeepaliveIntvl sets the tcp_keepalive_intvl option for the socket
func SockSetTcpKeepaliveIntvl(v int) SockOption {
	return func(s *Sock) {
		C.zsock_set_tcp_keepalive_intvl(unsafe.Pointer(s.zsockT), C.int(v))
	}
}

// TcpKeepaliveIntvl returns the current value of the socket's tcp_keepalive_intvl option
func TcpKeepaliveIntvl(s *Sock) int {
	val := C.zsock_tcp_keepalive_intvl(unsafe.Pointer(s.zsockT))
	return int(val)
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
func TcpAcceptFilter(s *Sock) string {
	val := C.zsock_tcp_accept_filter(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// Rcvmore returns the current value of the socket's rcvmore option
func Rcvmore(s *Sock) int {
	val := C.zsock_rcvmore(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Fd returns the current value of the socket's fd option
func Fd(s *Sock) int {
	val := C.zsock_fd(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Events returns the current value of the socket's events option
func Events(s *Sock) int {
	val := C.zsock_events(unsafe.Pointer(s.zsockT))
	return int(val)
}

// LastEndpoint returns the current value of the socket's last_endpoint option
func LastEndpoint(s *Sock) string {
	val := C.zsock_last_endpoint(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

