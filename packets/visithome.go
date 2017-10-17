package packets

import (
	"encoding/binary"
	"bytes"
)

type ClientVisitHome struct {
	Hi int32
	Lo int32
}

func NewClientVisitHomeFromBytes(buff []byte) (ClientVisitHome){
	o := ClientVisitHome{}

	buf := bytes.NewReader(buff)

	binary.Read(buf, binary.BigEndian, &o.Hi)
	binary.Read(buf, binary.BigEndian, &o.Lo)

	return o
}

func (o *ClientVisitHome) Bytes() ([]byte){
	// It creates the message bytes ready to be sent
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, o.Hi)
	binary.Write(buf, binary.BigEndian, o.Lo)

	return buf.Bytes()
}


type ServerVisitHome struct {
	Hi int32
	Lo int32
	Username string
}

func NewServerVisitHomeFromBytes(buff []byte) (ServerVisitHome){
	o := ServerVisitHome{}

	/* tests...
	var buf *bytes.Reader
	var fieldLen int32

	buf = bytes.NewReader(buff[0x58:])
	binary.Read(buf, binary.BigEndian, &o.Hi)
	binary.Read(buf, binary.BigEndian, &o.Lo)

	buf = bytes.NewReader(buff[0x71:])
	binary.Read(buf, binary.BigEndian, &fieldLen)
	name := make([]byte, fieldLen)
	binary.Read(buf, binary.BigEndian, &name)
	o.Username = string(name)
	*/

	return o
}

