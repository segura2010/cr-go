package packets

import (
	"encoding/binary"
	"bytes"
	"fmt"

	"github.com/segura2010/cr-go/utils"
	"github.com/segura2010/cr-go/packets/components"
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
	Deck [8]components.Card
	Hi int32
	Lo int32
	Username string
}

func NewServerVisitHomeFromBytes(buff []byte) (ServerVisitHome){
	o := ServerVisitHome{}

	// tests...
	var buf *bytes.Reader
	var tmp int32
	var btmp byte
	var isPresent byte
	var tmpbuf [30]byte

	buf = bytes.NewReader(buff)

	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	binary.Read(buf, binary.BigEndian, &btmp)
	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)

	// read actual deck cards
	for i:=0;i<8;i++{
		utils.ReadRrsInt32(buf, binary.BigEndian, &o.Deck[i].Id)
		utils.ReadRrsInt32(buf, binary.BigEndian, &o.Deck[i].Level)
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
		utils.ReadRrsInt32(buf, binary.BigEndian, &o.Deck[i].Count)
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	}

	binary.Read(buf, binary.BigEndian, &o.Hi)
	binary.Read(buf, binary.BigEndian, &o.Lo)

	// HomeUnknownSeason optional (read it if present, continue if not)
	binary.Read(buf, binary.BigEndian, &isPresent)
	fmt.Printf("\nseason?: %d", isPresent)
	if isPresent > 0{
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	}

	// HomeSeason[]
	// read length, then read each season info
	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	for i:=0;i<int(tmp);i++{
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	}

	// more unknowns..
	for i:=0;i<8;i++{
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	}

	binary.Read(buf, binary.BigEndian, &tmpbuf)
	fmt.Printf("\n%s", tmpbuf[:])

	//utils.ReadString(buf, binary.BigEndian, &o.Username)

	/*
	buf = bytes.NewReader(buff[0x71:])
	binary.Read(buf, binary.BigEndian, &fieldLen)
	name := make([]byte, fieldLen)
	binary.Read(buf, binary.BigEndian, &name)
	o.Username = string(name)
	*/

	return o
}

