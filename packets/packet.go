package packets

import(
	"encoding/binary"
	"bytes"

	"github.com/segura2010/cr-go/utils"
)

// It is a map with the definition of some client and server messages ID
var MessageType = map[string]uint16{
	// Server
	"ServerHello": 20100,
	"ServerLoginFailed": 20103,
	"ServerLoginOk": 20104,
	"ServerVisitedHome": 24113,
	"ServerKeepAliveOk": 20108,

	// Client
	"ClientHello": 10100,
	"ClientLogin": 10101,
	"ClientVisitHome": 14113,
	"ClientKeepAlive": 10108,
}

// It represents a packet
type Packet struct {
	// message type
    Type uint16
    // encrypted message length
    Length int32
    Version uint16
    // the encrypted message
    Payload []byte
    // the message itself
    DecryptedPayload []byte
}

func (o *Packet) Bytes() ([]byte){
	buf := new(bytes.Buffer)

	// prepare int24 length
	o.Length = int32(len(o.Payload))
	tmpbuf := new(bytes.Buffer)
	binary.Write(tmpbuf, binary.BigEndian, o.Length)
	fixedLen := tmpbuf.Bytes()[1:]

	binary.Write(buf, binary.BigEndian, o.Type)
	binary.Write(buf, binary.BigEndian, fixedLen)
	binary.Write(buf, binary.BigEndian, o.Version)
	binary.Write(buf, binary.BigEndian, o.Payload)

	return buf.Bytes()
}

func NewPacketFromBytes(buff []byte) (Packet){
	// It creates the message bytes ready to be sent
	buf := bytes.NewReader(buff)

	var messageLength [3]byte // int24
	o := Packet{}

	binary.Read(buf, binary.BigEndian, &o.Type)
	binary.Read(buf, binary.BigEndian, &messageLength)
	binary.Read(buf, binary.BigEndian, &o.Version)
	
	o.Length = int32(utils.Int24ToInt32(messageLength[:]))

	var payload = make([]byte, o.Length)
	binary.Read(buf, binary.BigEndian, &payload)
	o.Payload = payload[:]

	return o
}
