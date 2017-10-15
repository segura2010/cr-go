package packets

import (
	"encoding/binary"
	"bytes"
)

type ClientLogin struct {
	Hi int32
	Lo int32
	PassToken string
	MajorVersion int32
	MinorVersion int32
	Build int32
	ContentHash string
	Udid string
	OpenUdid string
	MacAddress string
	Device string
	AdvertisingGuid string
	OsVersion string
	IsAndroid byte
	U13 string
	AndroidId string
	DeviceLang string
	U16 byte
	Language byte
	FacebookId string
	AdvertisingEnabled byte
	AppleIfv string
	AppStore int32
	KunlunSso string
	KunlunUid string
	U24 string
	U25 string
	U26 byte
}

func NewDefaultClientLogin() (ClientLogin){
	// Default values for version 2.0.2 Android
	o := ClientLogin{
		Hi: 0,
		Lo: 0,
		MajorVersion: 3,
		MinorVersion: 0,
		Build: 690,
		ContentHash: "5765c06d5996ebf4a15b258903c3a0de922a57dd",
		Udid: "",
		AppStore: 29,
		Device: "D2303",
		DeviceLang: "es-ES",
		Language: 3,
		AdvertisingEnabled: 1,
		AdvertisingGuid: "cc823b32-3dbc-4455-8d00-f6b1ef6ad4b1",
		OsVersion: "4.4.4",
		AndroidId: "64e666532f11a6f4",
		U16: 1,
		U13: "",
		U24: "",
		U25: "",
		U26: 1,
	}

	return o
}

func NewClientLoginFromBytes(buff []byte) (ClientLogin){
	o := ClientLogin{}

	buf := bytes.NewReader(buff)

	binary.Read(buf, binary.BigEndian, &o.Hi)

	return o
}

func (o *ClientLogin) Bytes() ([]byte){
	// It creates the message bytes ready to be sent
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, o.Hi)
	binary.Write(buf, binary.BigEndian, o.Lo)
	binary.Write(buf, binary.BigEndian, []byte(o.PassToken))
	binary.Write(buf, binary.BigEndian, o.MajorVersion)
	binary.Write(buf, binary.BigEndian, o.MinorVersion)
	binary.Write(buf, binary.BigEndian, o.Build)
	binary.Write(buf, binary.BigEndian, []byte(o.ContentHash))
	binary.Write(buf, binary.BigEndian, []byte(o.Udid))
	binary.Write(buf, binary.BigEndian, []byte(o.OpenUdid))
	binary.Write(buf, binary.BigEndian, []byte(o.MacAddress))
	binary.Write(buf, binary.BigEndian, []byte(o.Device))
	binary.Write(buf, binary.BigEndian, []byte(o.AdvertisingGuid))
	binary.Write(buf, binary.BigEndian, []byte(o.OsVersion))
	binary.Write(buf, binary.BigEndian, o.IsAndroid)
	binary.Write(buf, binary.BigEndian, []byte(o.U13))
	binary.Write(buf, binary.BigEndian, []byte(o.AndroidId))
	binary.Write(buf, binary.BigEndian, []byte(o.DeviceLang))
	binary.Write(buf, binary.BigEndian, o.U16)
	binary.Write(buf, binary.BigEndian, o.Language)
	binary.Write(buf, binary.BigEndian, []byte(o.FacebookId))
	binary.Write(buf, binary.BigEndian, o.AdvertisingEnabled)
	binary.Write(buf, binary.BigEndian, []byte(o.AppleIfv))
	binary.Write(buf, binary.BigEndian, o.AppStore)
	binary.Write(buf, binary.BigEndian, []byte(o.KunlunSso))
	binary.Write(buf, binary.BigEndian, []byte(o.KunlunUid))
	binary.Write(buf, binary.BigEndian, []byte(o.U24))
	binary.Write(buf, binary.BigEndian, []byte(o.U25))
	binary.Write(buf, binary.BigEndian, o.U26)

	return buf.Bytes()
}

