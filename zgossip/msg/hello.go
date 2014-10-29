package msg

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"

	"github.com/zeromq/goczmq"
)

// Client says hello to server
type Hello struct {
	routingId []byte
	version   byte
}

// New creates new Hello message.
func NewHello() *Hello {
	hello := &Hello{}
	return hello
}

// String returns print friendly name.
func (h *Hello) String() string {
	str := "ZGOSSIP_MSG_HELLO:\n"
	str += fmt.Sprintf("    version = %v\n", h.version)
	return str
}

// Marshal serializes the message.
func (h *Hello) Marshal() ([]byte, error) {
	// Calculate size of serialized data
	bufferSize := 2 + 1 // Signature and message ID

	// version is a 1-byte integer
	bufferSize += 1

	// Now serialize the message
	tmpBuf := make([]byte, bufferSize)
	tmpBuf = tmpBuf[:0]
	buffer := bytes.NewBuffer(tmpBuf)
	binary.Write(buffer, binary.BigEndian, Signature)
	binary.Write(buffer, binary.BigEndian, HelloId)

	// version
	value, _ := strconv.ParseUint("1", 10, 1*8)
	binary.Write(buffer, binary.BigEndian, byte(value))

	return buffer.Bytes(), nil
}

// Unmarshals the message.
func (h *Hello) Unmarshal(frames ...[]byte) error {
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
		return errors.New("invalid signature")
	}

	// Get message id and parse per message type
	var id uint8
	binary.Read(buffer, binary.BigEndian, &id)
	if id != HelloId {
		return errors.New("malformed Hello message")
	}
	// version
	binary.Read(buffer, binary.BigEndian, &h.version)
	if h.version != 1 {
		return errors.New("malformed version message")
	}

	return nil
}

// Sends marshaled data through 0mq socket.
func (h *Hello) Send(socket *goczmq.Sock) (err error) {
	frame, err := h.Marshal()
	if err != nil {
		return err
	}

	socType := socket.GetType()
	if err != nil {
		return err
	}

	// If we're sending to a ROUTER, we send the routingId first
	if socType == goczmq.ROUTER {
		err = socket.SendBytes(h.routingId, goczmq.MORE)
		if err != nil {
			return err
		}
	}

	// Now send the data frame
	err = socket.SendBytes(frame, 0)
	if err != nil {
		return err
	}

	return err
}

// RoutingId returns the routingId for this message, routingId should be set
// whenever talking to a ROUTER.
func (h *Hello) RoutingId() []byte {
	return h.routingId
}

// SetRoutingId sets the routingId for this message, routingId should be set
// whenever talking to a ROUTER.
func (h *Hello) SetRoutingId(routingId []byte) {
	h.routingId = routingId
}

// Setversion sets the version.
func (h *Hello) SetVersion(version byte) {
	h.version = version
}

// version returns the version.
func (h *Hello) Version() byte {
	return h.version
}
