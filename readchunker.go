package goczmq

import "io"

// ReadChunker accepts a socket and a chunkSize, and implements
// the ReadFrom interface.
type ReadChunker struct {
	sock      *Sock
	chunkSize int64
}

// NewReadChunker takes a socket and a chunkSize and returns a new
// chunker instance
func NewReadChunker(s *Sock, cs int64) *ReadChunker {
	return &ReadChunker{
		sock:      s,
		chunkSize: cs,
	}
}

// ReadFrom reads from an io.Reader into a []byte of chunkSize.
// It writes each chunk of data as a frame to the socket.
// Each frame will have the more flag set until the last frame.
// All data from the io.Reader is sent in one atomic multi
// frame message
func (c *ReadChunker) ReadFrom(r io.Reader) (int64, error) {
	var total int64
	var n int
	var err error

	flag := Flag(1)

	p := make([]byte, c.chunkSize)

	for err == nil {
		n, err = r.Read(p)
		if err != nil {
			flag = Flag(0)
		}

		err := c.sock.SendFrame(p[:n], flag)
		if err != nil {
			return total, err
		}

		total += int64(n)
	}

	return total, err
}

// Destroy calls destroy on the underlying socket to clean it up
func (c *ReadChunker) Destroy() {
	c.sock.Destroy()
}
