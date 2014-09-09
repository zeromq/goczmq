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

// SetTos sets the tos option for the socket
func (z *Zsock) SetTos(val int) {
	C.zsock_set_tos(unsafe.Pointer(z.zsock_t), C.int(val))
}

// Tos returns the current value of the socket's tos option
func (z *Zsock) Tos() int {
	val := C.zsock_tos(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetRouterHandover sets the router_handover option for the socket
func (z *Zsock) SetRouterHandover(val int) {
	C.zsock_set_router_handover(unsafe.Pointer(z.zsock_t), C.int(val))
}

// SetRouterMandatory sets the router_mandatory option for the socket
func (z *Zsock) SetRouterMandatory(val int) {
	C.zsock_set_router_mandatory(unsafe.Pointer(z.zsock_t), C.int(val))
}

// SetProbeRouter sets the probe_router option for the socket
func (z *Zsock) SetProbeRouter(val int) {
	C.zsock_set_probe_router(unsafe.Pointer(z.zsock_t), C.int(val))
}

// SetReqRelaxed sets the req_relaxed option for the socket
func (z *Zsock) SetReqRelaxed(val int) {
	C.zsock_set_req_relaxed(unsafe.Pointer(z.zsock_t), C.int(val))
}

// SetReqCorrelate sets the req_correlate option for the socket
func (z *Zsock) SetReqCorrelate(val int) {
	C.zsock_set_req_correlate(unsafe.Pointer(z.zsock_t), C.int(val))
}

// SetConflate sets the conflate option for the socket
func (z *Zsock) SetConflate(val int) {
	C.zsock_set_conflate(unsafe.Pointer(z.zsock_t), C.int(val))
}

// SetZapDomain sets the zap_domain option for the socket
func (z *Zsock) SetZapDomain(val string) {
	C.zsock_set_zap_domain(unsafe.Pointer(z.zsock_t), C.CString(val))
}

// ZapDomain returns the current value of the socket's zap_domain option
func (z *Zsock) ZapDomain() string {
	val := C.zsock_zap_domain(unsafe.Pointer(z.zsock_t))
	return C.GoString(val)
}

// Mechanism returns the current value of the socket's mechanism option
func (z *Zsock) Mechanism() int {
	val := C.zsock_mechanism(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetPlainServer sets the plain_server option for the socket
func (z *Zsock) SetPlainServer(val int) {
	C.zsock_set_plain_server(unsafe.Pointer(z.zsock_t), C.int(val))
}

// PlainServer returns the current value of the socket's plain_server option
func (z *Zsock) PlainServer() int {
	val := C.zsock_plain_server(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetPlainUsername sets the plain_username option for the socket
func (z *Zsock) SetPlainUsername(val string) {
	C.zsock_set_plain_username(unsafe.Pointer(z.zsock_t), C.CString(val))
}

// PlainUsername returns the current value of the socket's plain_username option
func (z *Zsock) PlainUsername() string {
	val := C.zsock_plain_username(unsafe.Pointer(z.zsock_t))
	return C.GoString(val)
}

// SetPlainPassword sets the plain_password option for the socket
func (z *Zsock) SetPlainPassword(val string) {
	C.zsock_set_plain_password(unsafe.Pointer(z.zsock_t), C.CString(val))
}

// PlainPassword returns the current value of the socket's plain_password option
func (z *Zsock) PlainPassword() string {
	val := C.zsock_plain_password(unsafe.Pointer(z.zsock_t))
	return C.GoString(val)
}

// SetCurveServer sets the curve_server option for the socket
func (z *Zsock) SetCurveServer(val int) {
	C.zsock_set_curve_server(unsafe.Pointer(z.zsock_t), C.int(val))
}

// CurveServer returns the current value of the socket's curve_server option
func (z *Zsock) CurveServer() int {
	val := C.zsock_curve_server(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetCurvePublickey sets the curve_publickey option for the socket
func (z *Zsock) SetCurvePublickey(val string) {
	C.zsock_set_curve_publickey(unsafe.Pointer(z.zsock_t), C.CString(val))
}

// CurvePublickey returns the current value of the socket's curve_publickey option
func (z *Zsock) CurvePublickey() string {
	val := C.zsock_curve_publickey(unsafe.Pointer(z.zsock_t))
	return C.GoString(val)
}

// SetCurveSecretkey sets the curve_secretkey option for the socket
func (z *Zsock) SetCurveSecretkey(val string) {
	C.zsock_set_curve_secretkey(unsafe.Pointer(z.zsock_t), C.CString(val))
}

// CurveSecretkey returns the current value of the socket's curve_secretkey option
func (z *Zsock) CurveSecretkey() string {
	val := C.zsock_curve_secretkey(unsafe.Pointer(z.zsock_t))
	return C.GoString(val)
}

// SetCurveServerkey sets the curve_serverkey option for the socket
func (z *Zsock) SetCurveServerkey(val string) {
	C.zsock_set_curve_serverkey(unsafe.Pointer(z.zsock_t), C.CString(val))
}

// CurveServerkey returns the current value of the socket's curve_serverkey option
func (z *Zsock) CurveServerkey() string {
	val := C.zsock_curve_serverkey(unsafe.Pointer(z.zsock_t))
	return C.GoString(val)
}

// SetGssapiServer sets the gssapi_server option for the socket
func (z *Zsock) SetGssapiServer(val int) {
	C.zsock_set_gssapi_server(unsafe.Pointer(z.zsock_t), C.int(val))
}

// GssapiServer returns the current value of the socket's gssapi_server option
func (z *Zsock) GssapiServer() int {
	val := C.zsock_gssapi_server(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetGssapiPlaintext sets the gssapi_plaintext option for the socket
func (z *Zsock) SetGssapiPlaintext(val int) {
	C.zsock_set_gssapi_plaintext(unsafe.Pointer(z.zsock_t), C.int(val))
}

// GssapiPlaintext returns the current value of the socket's gssapi_plaintext option
func (z *Zsock) GssapiPlaintext() int {
	val := C.zsock_gssapi_plaintext(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetGssapiPrincipal sets the gssapi_principal option for the socket
func (z *Zsock) SetGssapiPrincipal(val string) {
	C.zsock_set_gssapi_principal(unsafe.Pointer(z.zsock_t), C.CString(val))
}

// GssapiPrincipal returns the current value of the socket's gssapi_principal option
func (z *Zsock) GssapiPrincipal() string {
	val := C.zsock_gssapi_principal(unsafe.Pointer(z.zsock_t))
	return C.GoString(val)
}

// SetGssapiServicePrincipal sets the gssapi_service_principal option for the socket
func (z *Zsock) SetGssapiServicePrincipal(val string) {
	C.zsock_set_gssapi_service_principal(unsafe.Pointer(z.zsock_t), C.CString(val))
}

// GssapiServicePrincipal returns the current value of the socket's gssapi_service_principal option
func (z *Zsock) GssapiServicePrincipal() string {
	val := C.zsock_gssapi_service_principal(unsafe.Pointer(z.zsock_t))
	return C.GoString(val)
}

// SetIpv6 sets the ipv6 option for the socket
func (z *Zsock) SetIpv6(val int) {
	C.zsock_set_ipv6(unsafe.Pointer(z.zsock_t), C.int(val))
}

// Ipv6 returns the current value of the socket's ipv6 option
func (z *Zsock) Ipv6() int {
	val := C.zsock_ipv6(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetImmediate sets the immediate option for the socket
func (z *Zsock) SetImmediate(val int) {
	C.zsock_set_immediate(unsafe.Pointer(z.zsock_t), C.int(val))
}

// Immediate returns the current value of the socket's immediate option
func (z *Zsock) Immediate() int {
	val := C.zsock_immediate(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetRouterRaw sets the router_raw option for the socket
func (z *Zsock) SetRouterRaw(val int) {
	C.zsock_set_router_raw(unsafe.Pointer(z.zsock_t), C.int(val))
}

// SetIpv4only sets the ipv4only option for the socket
func (z *Zsock) SetIpv4only(val int) {
	C.zsock_set_ipv4only(unsafe.Pointer(z.zsock_t), C.int(val))
}

// Ipv4only returns the current value of the socket's ipv4only option
func (z *Zsock) Ipv4only() int {
	val := C.zsock_ipv4only(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetDelayAttachOnConnect sets the delay_attach_on_connect option for the socket
func (z *Zsock) SetDelayAttachOnConnect(val int) {
	C.zsock_set_delay_attach_on_connect(unsafe.Pointer(z.zsock_t), C.int(val))
}

// Type returns the current value of the socket's type option
func (z *Zsock) Type() int {
	val := C.zsock_type(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetSndhwm sets the sndhwm option for the socket
func (z *Zsock) SetSndhwm(val int) {
	C.zsock_set_sndhwm(unsafe.Pointer(z.zsock_t), C.int(val))
}

// Sndhwm returns the current value of the socket's sndhwm option
func (z *Zsock) Sndhwm() int {
	val := C.zsock_sndhwm(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetRcvhwm sets the rcvhwm option for the socket
func (z *Zsock) SetRcvhwm(val int) {
	C.zsock_set_rcvhwm(unsafe.Pointer(z.zsock_t), C.int(val))
}

// Rcvhwm returns the current value of the socket's rcvhwm option
func (z *Zsock) Rcvhwm() int {
	val := C.zsock_rcvhwm(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetAffinity sets the affinity option for the socket
func (z *Zsock) SetAffinity(val int) {
	C.zsock_set_affinity(unsafe.Pointer(z.zsock_t), C.int(val))
}

// Affinity returns the current value of the socket's affinity option
func (z *Zsock) Affinity() int {
	val := C.zsock_affinity(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetSubscribe sets the subscribe option for the socket
func (z *Zsock) SetSubscribe(val string) {
	C.zsock_set_subscribe(unsafe.Pointer(z.zsock_t), C.CString(val))
}

// SetUnsubscribe sets the unsubscribe option for the socket
func (z *Zsock) SetUnsubscribe(val string) {
	C.zsock_set_unsubscribe(unsafe.Pointer(z.zsock_t), C.CString(val))
}

// SetIdentity sets the identity option for the socket
func (z *Zsock) SetIdentity(val string) {
	C.zsock_set_identity(unsafe.Pointer(z.zsock_t), C.CString(val))
}

// Identity returns the current value of the socket's identity option
func (z *Zsock) Identity() string {
	val := C.zsock_identity(unsafe.Pointer(z.zsock_t))
	return C.GoString(val)
}

// SetRate sets the rate option for the socket
func (z *Zsock) SetRate(val int) {
	C.zsock_set_rate(unsafe.Pointer(z.zsock_t), C.int(val))
}

// Rate returns the current value of the socket's rate option
func (z *Zsock) Rate() int {
	val := C.zsock_rate(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetRecoveryIvl sets the recovery_ivl option for the socket
func (z *Zsock) SetRecoveryIvl(val int) {
	C.zsock_set_recovery_ivl(unsafe.Pointer(z.zsock_t), C.int(val))
}

// RecoveryIvl returns the current value of the socket's recovery_ivl option
func (z *Zsock) RecoveryIvl() int {
	val := C.zsock_recovery_ivl(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetSndbuf sets the sndbuf option for the socket
func (z *Zsock) SetSndbuf(val int) {
	C.zsock_set_sndbuf(unsafe.Pointer(z.zsock_t), C.int(val))
}

// Sndbuf returns the current value of the socket's sndbuf option
func (z *Zsock) Sndbuf() int {
	val := C.zsock_sndbuf(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetRcvbuf sets the rcvbuf option for the socket
func (z *Zsock) SetRcvbuf(val int) {
	C.zsock_set_rcvbuf(unsafe.Pointer(z.zsock_t), C.int(val))
}

// Rcvbuf returns the current value of the socket's rcvbuf option
func (z *Zsock) Rcvbuf() int {
	val := C.zsock_rcvbuf(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetLinger sets the linger option for the socket
func (z *Zsock) SetLinger(val int) {
	C.zsock_set_linger(unsafe.Pointer(z.zsock_t), C.int(val))
}

// Linger returns the current value of the socket's linger option
func (z *Zsock) Linger() int {
	val := C.zsock_linger(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetReconnectIvl sets the reconnect_ivl option for the socket
func (z *Zsock) SetReconnectIvl(val int) {
	C.zsock_set_reconnect_ivl(unsafe.Pointer(z.zsock_t), C.int(val))
}

// ReconnectIvl returns the current value of the socket's reconnect_ivl option
func (z *Zsock) ReconnectIvl() int {
	val := C.zsock_reconnect_ivl(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetReconnectIvlMax sets the reconnect_ivl_max option for the socket
func (z *Zsock) SetReconnectIvlMax(val int) {
	C.zsock_set_reconnect_ivl_max(unsafe.Pointer(z.zsock_t), C.int(val))
}

// ReconnectIvlMax returns the current value of the socket's reconnect_ivl_max option
func (z *Zsock) ReconnectIvlMax() int {
	val := C.zsock_reconnect_ivl_max(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetBacklog sets the backlog option for the socket
func (z *Zsock) SetBacklog(val int) {
	C.zsock_set_backlog(unsafe.Pointer(z.zsock_t), C.int(val))
}

// Backlog returns the current value of the socket's backlog option
func (z *Zsock) Backlog() int {
	val := C.zsock_backlog(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetMaxmsgsize sets the maxmsgsize option for the socket
func (z *Zsock) SetMaxmsgsize(val int) {
	C.zsock_set_maxmsgsize(unsafe.Pointer(z.zsock_t), C.int(val))
}

// Maxmsgsize returns the current value of the socket's maxmsgsize option
func (z *Zsock) Maxmsgsize() int {
	val := C.zsock_maxmsgsize(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetMulticastHops sets the multicast_hops option for the socket
func (z *Zsock) SetMulticastHops(val int) {
	C.zsock_set_multicast_hops(unsafe.Pointer(z.zsock_t), C.int(val))
}

// MulticastHops returns the current value of the socket's multicast_hops option
func (z *Zsock) MulticastHops() int {
	val := C.zsock_multicast_hops(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetRcvtimeo sets the rcvtimeo option for the socket
func (z *Zsock) SetRcvtimeo(val int) {
	C.zsock_set_rcvtimeo(unsafe.Pointer(z.zsock_t), C.int(val))
}

// Rcvtimeo returns the current value of the socket's rcvtimeo option
func (z *Zsock) Rcvtimeo() int {
	val := C.zsock_rcvtimeo(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetSndtimeo sets the sndtimeo option for the socket
func (z *Zsock) SetSndtimeo(val int) {
	C.zsock_set_sndtimeo(unsafe.Pointer(z.zsock_t), C.int(val))
}

// Sndtimeo returns the current value of the socket's sndtimeo option
func (z *Zsock) Sndtimeo() int {
	val := C.zsock_sndtimeo(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetXpubVerbose sets the xpub_verbose option for the socket
func (z *Zsock) SetXpubVerbose(val int) {
	C.zsock_set_xpub_verbose(unsafe.Pointer(z.zsock_t), C.int(val))
}

// SetTcpKeepalive sets the tcp_keepalive option for the socket
func (z *Zsock) SetTcpKeepalive(val int) {
	C.zsock_set_tcp_keepalive(unsafe.Pointer(z.zsock_t), C.int(val))
}

// TcpKeepalive returns the current value of the socket's tcp_keepalive option
func (z *Zsock) TcpKeepalive() int {
	val := C.zsock_tcp_keepalive(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetTcpKeepaliveIdle sets the tcp_keepalive_idle option for the socket
func (z *Zsock) SetTcpKeepaliveIdle(val int) {
	C.zsock_set_tcp_keepalive_idle(unsafe.Pointer(z.zsock_t), C.int(val))
}

// TcpKeepaliveIdle returns the current value of the socket's tcp_keepalive_idle option
func (z *Zsock) TcpKeepaliveIdle() int {
	val := C.zsock_tcp_keepalive_idle(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetTcpKeepaliveCnt sets the tcp_keepalive_cnt option for the socket
func (z *Zsock) SetTcpKeepaliveCnt(val int) {
	C.zsock_set_tcp_keepalive_cnt(unsafe.Pointer(z.zsock_t), C.int(val))
}

// TcpKeepaliveCnt returns the current value of the socket's tcp_keepalive_cnt option
func (z *Zsock) TcpKeepaliveCnt() int {
	val := C.zsock_tcp_keepalive_cnt(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetTcpKeepaliveIntvl sets the tcp_keepalive_intvl option for the socket
func (z *Zsock) SetTcpKeepaliveIntvl(val int) {
	C.zsock_set_tcp_keepalive_intvl(unsafe.Pointer(z.zsock_t), C.int(val))
}

// TcpKeepaliveIntvl returns the current value of the socket's tcp_keepalive_intvl option
func (z *Zsock) TcpKeepaliveIntvl() int {
	val := C.zsock_tcp_keepalive_intvl(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// SetTcpAcceptFilter sets the tcp_accept_filter option for the socket
func (z *Zsock) SetTcpAcceptFilter(val string) {
	C.zsock_set_tcp_accept_filter(unsafe.Pointer(z.zsock_t), C.CString(val))
}

// TcpAcceptFilter returns the current value of the socket's tcp_accept_filter option
func (z *Zsock) TcpAcceptFilter() string {
	val := C.zsock_tcp_accept_filter(unsafe.Pointer(z.zsock_t))
	return C.GoString(val)
}

// Rcvmore returns the current value of the socket's rcvmore option
func (z *Zsock) Rcvmore() int {
	val := C.zsock_rcvmore(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// Fd returns the current value of the socket's fd option
func (z *Zsock) Fd() int {
	val := C.zsock_fd(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// Events returns the current value of the socket's events option
func (z *Zsock) Events() int {
	val := C.zsock_events(unsafe.Pointer(z.zsock_t))
	return int(val)
}

// LastEndpoint returns the current value of the socket's last_endpoint option
func (z *Zsock) LastEndpoint() string {
	val := C.zsock_last_endpoint(unsafe.Pointer(z.zsock_t))
	return C.GoString(val)
}
