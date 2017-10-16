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

