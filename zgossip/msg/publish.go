package msg

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"

	"github.com/zeromq/goczmq"
)

// Client or server announces a new tuple
type Publish struct {
	routingId []byte
	version   byte
	Key       string
	Value     string
	Ttl       uint32
}

// New creates new Publish message.
func NewPublish() *Publish {
	publish := &Publish{}
	return publish
}

// String returns print friendly name.
func (p *Publish) String() string {
	str := "ZGOSSIP_MSG_PUBLISH:\n"
	str += fmt.Sprintf("    version = %v\n", p.version)
	str += fmt.Sprintf("    Key = %v\n", p.Key)
	str += fmt.Sprintf("    Value = %v\n", p.Value)
	str += fmt.Sprintf("    Ttl = %v\n", p.Ttl)
	return str
}

// Marshal serializes the message.
func (p *Publish) Marshal() ([]byte, error) {
	// Calculate size of serialized data
	bufferSize := 2 + 1 // Signature and message ID

	// version is a 1-byte integer
	bufferSize += 1

	// Key is a string with 1-byte length
	bufferSize++ // Size is one byte
	bufferSize += len(p.Key)

	// Value is a string with 4-byte length
	bufferSize += 4 // Size is 4 bytes
	bufferSize += len(p.Value)

	// Ttl is a 4-byte integer
	bufferSize += 4

	// Now serialize the message
	tmpBuf := make([]byte, bufferSize)
	tmpBuf = tmpBuf[:0]
	buffer := bytes.NewBuffer(tmpBuf)
	binary.Write(buffer, binary.BigEndian, Signature)
	binary.Write(buffer, binary.BigEndian, PublishId)

	// version
	value, _ := strconv.ParseUint("1", 10, 1*8)
	binary.Write(buffer, binary.BigEndian, byte(value))

	// Key
	putString(buffer, p.Key)

	// Value
	putLongString(buffer, p.Value)

	// Ttl
	binary.Write(buffer, binary.BigEndian, p.Ttl)

	return buffer.Bytes(), nil
}

// Unmarshal unmarshals the message.
func (p *Publish) Unmarshal(frames ...[]byte) error {
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
	if id != PublishId {
		return errors.New("malformed Publish message")
	}
	// version
	binary.Read(buffer, binary.BigEndian, &p.version)
	if p.version != 1 {
		return errors.New("malformed version message")
	}
	// Key
	p.Key = getString(buffer)
	// Value
	p.Value = getLongString(buffer)
	// Ttl
	binary.Read(buffer, binary.BigEndian, &p.Ttl)

	return nil
}

// Send sends marshaled data through 0mq socket.
func (p *Publish) Send(sock *goczmq.Sock) (err error) {
	frame, err := p.Marshal()
	if err != nil {
		return err
	}

	socType := sock.GetType()
	if err != nil {
		return err
	}

	// If we're sending to a ROUTER, we send the routingId first
	if socType == goczmq.ROUTER {
		err = sock.SendFrame(p.routingId, goczmq.MORE)
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
func (p *Publish) RoutingId() []byte {
	return p.routingId
}

// SetRoutingId sets the routingId for this message, routingId should be set
// whenever talking to a ROUTER.
func (p *Publish) SetRoutingId(routingId []byte) {
	p.routingId = routingId
}

// Setversion sets the version.
func (p *Publish) SetVersion(version byte) {
	p.version = version
}

// version returns the version.
func (p *Publish) Version() byte {
	return p.version
}
