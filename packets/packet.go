package packets

import(
	"encoding/binary"
	"bytes"

	"github.com/segura2010/cr-go/utils"
)

var MessageType = map[string]uint16{
	// Server
	"ServerHello": 20100,
	"ServerLoginFailed": 20103,
	"ServerLoginOk": 20104,
	"ServerVisitedHome": 24113,

	// Client
	"ClientHello": 10100,
	"ClientLogin": 10101,
	"ClientVisitHome": 14113,
}

type Packet struct {
    Type uint16
    Length int32
    Version uint16
    Payload []byte // the encrypted message
    DecryptedPayload []byte // the message itself
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

	var messageLength [6]byte // int24
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
