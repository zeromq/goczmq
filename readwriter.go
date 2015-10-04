package goczmq

import "C"

// ReadWriter provides an io.ReadWriter compatible
// interface for goczmq.Sock

type ReadWriter struct {
	sock         *Sock
	clientIDs    []string
	frame        []byte
	currentIndex int
}

// NewReadWriter accepts a sock and returns a goczmq.ReadWriter. The
// io.ReadWriter should now be considered responsible for this
// Sock.
func NewReadWriter(sock *Sock) *ReadWriter {
	return &ReadWriter{
		sock: sock,
	}
}

// Read satisifies io.Read
func (r *ReadWriter) Read(p []byte) (int, error) {
	var totalRead int
	var totalFrame int

	var flag int
	var err error

	if r.currentIndex == 0 {
		r.frame, flag, err = r.sock.RecvFrame()

		if r.sock.GetType() == Router && r.currentIndex == 0 {
			r.clientIDs = append(r.clientIDs, string(r.frame))
			r.frame = []byte{0}
			r.frame, flag, err = r.sock.RecvFrame()
		}

		if flag == FlagMore && r.sock.GetType() != Router {
			return totalRead, ErrMultiPartUnsupported
		}

		if err != nil {
			return totalRead, err
		}
	}

	totalRead += copy(p[:], r.frame[r.currentIndex:])
	totalFrame += len(r.frame)

	if totalFrame > len(p) {
		r.currentIndex = totalRead
	} else {
		r.currentIndex = 0
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
