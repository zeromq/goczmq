package msg

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"

	"github.com/zeromq/goczmq"
)

// This is a simple chunk protocol for streaming large messages broken into chunks
type Chunk struct {
	routingId []byte
	version   byte
	More      byte
	Payload   []byte
}

// New creates new Chunk message.
func NewChunk() *Chunk {
	chunk := &Chunk{}
	return chunk
}

// String returns print friendly name.
func (c *Chunk) String() string {
	str := "IOCHUNK_MSG_CHUNK:\n"
	str += fmt.Sprintf("    version = %v\n", c.version)
	str += fmt.Sprintf("    More = %v\n", c.More)
	str += fmt.Sprintf("    Payload = %v\n", c.Payload)
	return str
}

// Marshal serializes the message.
func (c *Chunk) Marshal() ([]byte, error) {
	// Calculate size of serialized data
	bufferSize := 2 + 1 // Signature and message ID

	// version is a 1-byte integer
	bufferSize += 1

	// More is a 1-byte integer
	bufferSize += 1

	// Payload is a block of []byte with one byte length
	bufferSize += 1 + len(c.Payload)

	// Now serialize the message
	tmpBuf := make([]byte, bufferSize)
	tmpBuf = tmpBuf[:0]
	buffer := bytes.NewBuffer(tmpBuf)
	binary.Write(buffer, binary.BigEndian, Signature)
	binary.Write(buffer, binary.BigEndian, ChunkId)

	// version
	value, _ := strconv.ParseUint("1", 10, 1*8)
	binary.Write(buffer, binary.BigEndian, byte(value))

	// More
	binary.Write(buffer, binary.BigEndian, c.More)

	// Payload
	putBytes(buffer, c.Payload)

	return buffer.Bytes(), nil
}

// Unmarshal unmarshals the message.
func (c *Chunk) Unmarshal(frames ...[]byte) error {
	if frames == nil {
		return errors.New("Can't unmarshal empty message")
	}

	frame := frames[0]
	frames = frames[1:]

	buffer := bytes.NewBuffer(frame)

	// Get and check protocol signature
	var signature uint16
	binary.Read(buffer, binary.BigEndian, &signature)
	if signature != Signature {
		return fmt.Errorf("invalid signature %X != %X", Signature, signature)
	}

	// Get message id and parse per message type
	var id uint8
	binary.Read(buffer, binary.BigEndian, &id)
	if id != ChunkId {
		return errors.New("malformed Chunk message")
	}
	// version
	binary.Read(buffer, binary.BigEndian, &c.version)
	if c.version != 1 {
		return errors.New("malformed version message")
	}
	// More
	binary.Read(buffer, binary.BigEndian, &c.More)
	// Payload
	c.Payload = getBytes(buffer)

	return nil
}

// Send sends marshaled data through 0mq socket.
func (c *Chunk) Send(sock goczmqSock) (err error) {
	frame, err := c.Marshal()
	if err != nil {
		return err
	}

	socType := sock.GetType()
	if err != nil {
		return err
	}

	// If we're sending to a ROUTER, we send the routingId first
	if socType == goczmq.ROUTER {
		err = sock.SendFrame(c.routingId, goczmq.MORE)
		if err != nil {
			return err
		}
	}

	// Now send the data frame
	err = sock.SendFrame(frame, 0)
	if err != nil {
		return err
	}

	return err
}

// RoutingId returns the routingId for this message, routingId should be set
// whenever talking to a ROUTER.
func (c *Chunk) RoutingId() []byte {
	return c.routingId
}

// SetRoutingId sets the routingId for this message, routingId should be set
// whenever talking to a ROUTER.
func (c *Chunk) SetRoutingId(routingId []byte) {
	c.routingId = routingId
}

// Setversion sets the version.
func (c *Chunk) SetVersion(version byte) {
	c.version = version
}

// version returns the version.
func (c *Chunk) Version() byte {
	return c.version
}
