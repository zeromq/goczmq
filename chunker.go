package goczmq

import (
	"fmt"
	"io"

	iomsg "github.com/taotetek/goczmq/iochunk/msg"
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
// It writes each chunk of data as a two frame message. The first
// frame is used to indicate if this is the last message or if
// there are more messages containing file data to come.
func (c *ReadChunker) ReadFrom(r io.Reader) (int64, error) {
	var total int64
	var n int
	var err error

	expect := []byte("MOAR")

	p := make([]byte, c.chunkSize)

	for err == nil {
		n, err = r.Read(p)

		msg := iomsg.NewChunk()
		msg.More = 1

		if err != nil {
			msg.More = 0
		}

		msg.Payload = p

		err = msg.Send(c.sock)
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

// WriteTo to reads each chunk message one at
// a time and writes  them to an io.Writer.
func (c *WriteChunker) WriteTo(w io.Writer) (int64, error) {
	var total int64
	var n int
	var err error
	var chunk [][]byte

	expect := []byte("MOAR")

	more := 1
	for more == 1 {
		transit, err := iomsg.Recv(c.sock)
		if err != nil {
			return total, fmt.Errorf("protocol error")
		}

		chunk := transit.(*iomsg.Chunk)
		n, err = w.Write(chunk.Payload)
		total += int64(n)
		if err != nil {
			return total, err
		}

		more = chunk.More
	}

	return total, err
}

// Destroy calls destroy on the underlying socket to clean it up
func (c *WriteChunker) Destroy() {
	c.sock.Destroy()
}
