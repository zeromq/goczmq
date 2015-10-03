package goczmq

import "C"

// ReadWriter provides an io.ReadWriter compatible
// interface for goczmq.Sock
type ReadWriter struct {
	sock      *Sock
	clientIDs []string
}

// NewReadWriter accepts a sock and returns a goczmq.ReadWriter. The
// io.ReadWriter should now be considered responsible for this
// Sock.
func NewReadWriter(sock *Sock) *ReadWriter {
	return &ReadWriter{
		sock:      sock,
		clientIDs: make([]string, 0),
	}
}

// Read satisifies io.Read
func (r *ReadWriter) Read(p []byte) (int, error) {
	var totalRead int
	var totalFrame int

	frame, flag, err := r.sock.RecvFrame()
	if err != nil {
		return totalRead, err
	}

	if r.sock.GetType() == Router {
		r.clientIDs = append(r.clientIDs, string(frame))
	} else {
		totalRead += copy(p[:], frame[:])
		totalFrame += len(frame)
	}

	for flag == FlagMore {
		frame, flag, err = r.sock.RecvFrame()
		if err != nil {
			return totalRead, err
		}
		totalRead += copy(p[totalRead:], frame[:])
		totalFrame += len(frame)
	}

	if totalFrame > len(p) {
		err = ErrSliceFull
	} else {
		err = nil
	}

	return totalRead, err
}

// Write satisfies io.Write
func (r *ReadWriter) Write(p []byte) (int, error) {
	var total int
	if r.sock.GetType() == Router {
		err := r.sock.SendFrame(r.GetLastClientID(), FlagMore)
		if err != nil {
			return total, err
		}
	}
	err := r.sock.SendFrame(p, 0)
	if err != nil {
		return total, err
	}

	return len(p), nil

}

// GetLastClientID returns the id of the last client you received
// a message from if the underlying socket is a Router socket
func (r *ReadWriter) GetLastClientID() []byte {
	id := []byte(r.clientIDs[0])
	r.clientIDs = r.clientIDs[1:]
	return id
}

// SetLastClientID lets you manually set the id of the client
// you last received a message from if the underlying socket
// is a Router socket
func (r *ReadWriter) SetLastClientID(id []byte) {
	r.clientIDs = append(r.clientIDs, string(id))
}

// Destroy destroys both the ReadWriter and the underlying Sock
func (r *ReadWriter) Destroy() {
	r.sock.Destroy()
}
