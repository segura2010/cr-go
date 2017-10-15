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

	buf := bytes.NewReader(buff)

	binary.Read(buf, binary.BigEndian, &o.Protocol)
	binary.Read(buf, binary.BigEndian, &o.KeyVersion)
	binary.Read(buf, binary.BigEndian, &o.MajorVersion)
	binary.Read(buf, binary.BigEndian, &o.MinorVersion)
	binary.Read(buf, binary.BigEndian, &o.Build)

	var contentHashLen int32
	binary.Read(buf, binary.BigEndian, &contentHashLen)
	contentHash := make([]byte, contentHashLen)
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

	contentHashLen := int32(len(o.ContentHash))
	binary.Write(buf, binary.BigEndian, contentHashLen)
	binary.Write(buf, binary.BigEndian, []byte(o.ContentHash))

	binary.Write(buf, binary.BigEndian, o.DeviceType)
	binary.Write(buf, binary.BigEndian, o.AppStore)

	return buf.Bytes()
}

type ServerHello struct {
	SessionKey []byte
}

func NewServerHelloFromBytes(buff []byte) (ServerHello){
	o := ServerHello{}

	buf := bytes.NewReader(buff)
	var keyLen int32

	binary.Read(buf, binary.BigEndian, &keyLen)

	sessionKey := make([]byte, keyLen)
	binary.Read(buf, binary.BigEndian, &sessionKey)

	o.SessionKey = sessionKey[:]

	return o
}

