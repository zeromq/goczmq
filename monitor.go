package goczmq

/*
#include "czmq.h"

zactor_t *Monitor_new (zsock_t *sock) {
	zactor_t *monitor = zactor_new(zmonitor, sock);
	return monitor;
}

int Monitor_verbose (zactor_t *monitor) {
	return zstr_sendx(monitor, "VERBOSE", NULL);
}

int Monitor_listen (zactor_t *monitor, const char *type) {
	return zstr_sendx(monitor, "LISTEN", type, NULL);
}

void Monitor_destroy (zactor_t *monitor) {
	zactor_destroy(&monitor);
}
*/
import "C"

import (
	"unsafe"
)

// Monitor provides an API for obtaining socket events
type Monitor struct {
	zactorT *C.struct__zactor_t
}

// NewMonitor creates new Monitor actor.
func NewMonitor(socket *Sock) *Monitor {
	m := &Monitor{}
	m.zactorT = C.Monitor_new((*C.struct__zsock_t)(unsafe.Pointer(socket.zsockT)))
	return m
}

// Listen specifies which events to listen for. "ALL" is also supported.
func (m *Monitor) Listen(event string) error {
	cmd := C.CString("LISTEN")
	defer C.free(unsafe.Pointer(cmd))

	eventStr := C.CString(event)
	defer C.free(unsafe.Pointer(eventStr))

SendListen:
	rc, err := C.Monitor_listen((*C.struct__zactor_t)(unsafe.Pointer(m.zactorT)), eventStr)
	if rc == -1 {
		if isRetryableError(err) {
			goto SendListen
		}
		return ErrActorCmd
	}

	return nil
}

// Start activates the socket monitoring. Additional Listen() calls will not have an effect after this.
func (m *Monitor) Start() error {
	cmd := C.CString("START")
	defer C.free(unsafe.Pointer(cmd))

SendStart:
	rc, err := C.zstr_send(unsafe.Pointer(m.zactorT), cmd)
	if rc == -1 {
		if isRetryableError(err) {
			goto SendStart
		}
		return ErrActorCmd
	}
	C.zsock_wait(unsafe.Pointer(m.zactorT))

	return nil
}

// Verbose enables verbose mode, logging activity to stdout
func (m *Monitor) Verbose() error {
SendVerbose:
	rc, err := C.Monitor_verbose((*C.struct__zactor_t)(unsafe.Pointer(m.zactorT)))
	if rc == -1 {
		if isRetryableError(err) {
			goto SendVerbose
		}
		return ErrActorCmd
	}

	return nil
}

// Socket returns the actor as a Sock instance, useful and necessary for being able to receive messages
func (m *Monitor) Socket() *Sock {
	s := &Sock{}
	s.zsockT = (*C.struct__zsock_t)(unsafe.Pointer(m.zactorT))
	return s
}

// Destroy destroys the monitor instance
func (m *Monitor) Destroy() {
	C.Monitor_destroy((*C.struct__zactor_t)(unsafe.Pointer(m.zactorT)))
}
