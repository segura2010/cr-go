package packets

import (
	"encoding/binary"
	"bytes"

	"github.com/segura2010/cr-go/utils"
)

// It represents the Hello packet sent by the client at the beginning of the communication
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

// creates a new hello message with default values
func NewDefaultClientHello() (ClientHello){
	// Default values for version 2.0.2 Android
	o := ClientHello{
		Protocol: 1,
		KeyVersion: 15,
		MajorVersion: 3,
		MinorVersion: 0,
		Build: 830,
		ContentHash: "3f6063a89e32b2403462696cbb5e0f68f8e58ea2",
		DeviceType: 2,
		AppStore: 2,
	}

	return o
}

// creates a ClientHello object from a byte array
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

// Returns the byte array representation of the ClientHello message (ready to send)
func (o *ClientHello) Bytes() ([]byte){
	// It creates the message bytes ready to be sent
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, o.Protocol)
	binary.Write(buf, binary.BigEndian, o.KeyVersion)
	binary.Write(buf, binary.BigEndian, o.MajorVersion)
	binary.Write(buf, binary.BigEndian, o.MinorVersion)
	binary.Write(buf, binary.BigEndian, o.Build)
	
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.ContentHash))

	binary.Write(buf, binary.BigEndian, o.DeviceType)
	binary.Write(buf, binary.BigEndian, o.AppStore)

	return buf.Bytes()
}

// It represents the server response for the hello message
type ServerHello struct {
	SessionKey []byte
}

// creates a ServerHello object from a byte array
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

