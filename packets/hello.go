package packets

import (
	"encoding/binary"
	"bytes"
)

type ClientHello struct {
	Protocol int32
	KeyVersion int32
	MajorVersion int32
	MinorVersion int32
	Build int32
	ContentHash string
	DeviceType int32
	AppStore int32
}

func NewDefaultClientHello() (ClientHello){
	// Default values for version 2.0.2 Android
	o := ClientHello{
		Protocol: 1,
		KeyVersion: 14,
		MajorVersion: 3,
		MinorVersion: 0,
		Build: 690,
		ContentHash: "5765c06d5996ebf4a15b258903c3a0de922a57dd",
		DeviceType: 2,
		AppStore: 29,
	}

	return o
}

func NewClientHelloFromBytes(buff []byte) (ClientHello){
	o := ClientHello{}

	// calculate the length of contentHash in order to be able to read it
	buffLen := len(buff)
	contentHashLen := buffLen - (7*4)
	contentHash := make([]byte, contentHashLen)

	buf := bytes.NewReader(buff)

	binary.Read(buf, binary.BigEndian, &o.Protocol)
	binary.Read(buf, binary.BigEndian, &o.KeyVersion)
	binary.Read(buf, binary.BigEndian, &o.MajorVersion)
	binary.Read(buf, binary.BigEndian, &o.MinorVersion)
	binary.Read(buf, binary.BigEndian, &o.Build)
	binary.Read(buf, binary.BigEndian, &contentHash)
	binary.Read(buf, binary.BigEndian, &o.DeviceType)
	binary.Read(buf, binary.BigEndian, &o.AppStore)

	o.ContentHash = string(contentHash[:])

	return o
}

func (o *ClientHello) Bytes() ([]byte){
	// It creates the message bytes ready to be sent
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, o.Protocol)
	binary.Write(buf, binary.BigEndian, o.KeyVersion)
	binary.Write(buf, binary.BigEndian, o.MajorVersion)
	binary.Write(buf, binary.BigEndian, o.MinorVersion)
	binary.Write(buf, binary.BigEndian, o.Build)
	binary.Write(buf, binary.BigEndian, []byte(o.ContentHash))
	binary.Write(buf, binary.BigEndian, o.DeviceType)
	binary.Write(buf, binary.BigEndian, o.AppStore)

	return buf.Bytes()
}

