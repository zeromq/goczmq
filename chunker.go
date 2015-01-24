package goczmq

import (
	"bytes"
	"fmt"
	"io"
)

// ReadChunker accepts a socket and a chunkSize, and implements
// the ReaderFrom interface.
type ReadChunker struct {
	sock      *Sock
	chunkSize int64
}

// NewReadChunker takes a socket and a chunkSize and returns a new
// ReadChunker instance
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

	expect := []byte("MOAR")

	p := make([]byte, c.chunkSize)

	for err == nil {
		n, err = r.Read(p)
		if err != nil {
			expect = []byte("NOMOAR")
		}

		err := c.sock.SendFrame(expect, 1)
		if err != nil {
			return total, err
		}

		err = c.sock.SendFrame(p[:n], 0)
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

// WriteChunker accepts a socket and implements
// the WriterTo interface
type WriteChunker struct {
	sock *Sock
}

// NewWriteChunker takes a socket and returns a new
// WriteChunker instance
func NewWriteChunker(s *Sock) *WriteChunker {
	return &WriteChunker{
		sock: s,
	}
}

// WriteTo to reads each frame of a multi part message one at
// a time and writes  them to an io.Writer.
func (c *WriteChunker) WriteTo(w io.Writer) (int64, error) {
	var total int64
	var n int
	var err error
	var chunk [][]byte

	expect := []byte("MOAR")

	for bytes.Compare(expect, []byte("MOAR")) == 0 {
		chunk, err = c.sock.RecvMessage()
		if len(chunk) != 2 {
			return total, fmt.Errorf("protocol error")
		}

		expect = chunk[0]

		if bytes.Compare(expect, []byte("MOAR")) != 0 &&
			bytes.Compare(expect, []byte("NOMOAR")) != 0 {
			return total, fmt.Errorf("protocol error")
		}

		n, err = w.Write(chunk[1])
		total += int64(n)
		if err != nil {
			return total, err
		}
	}

	return total, err
}

// Destroy calls destroy on the underlying socket to clean it up
func (c *WriteChunker) Destroy() {
	c.sock.Destroy()
}
