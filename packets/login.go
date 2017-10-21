package packets

import (
	"encoding/binary"
	"bytes"

	"github.com/segura2010/cr-go/utils"
)

// It represents the ClientLogin message with all the fields we need to send to the server
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
	MacAddress [4]byte
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

// it creates a new ClientLogin message using the default values
// you have to set Hi, Lo and PassToken with your account details
func NewDefaultClientLogin() (ClientLogin){
	// Default values for version 2.0.2 Android
	o := ClientLogin{
		Hi: 0,
		Lo: 0,
		MajorVersion: 3,
		MinorVersion: 0,
		Build: 690,
		ContentHash: "ccfe9f95663453bc252f4367055d3ec3e022ae65",
		Udid: "",
		OpenUdid: "68e646598f73a6f5",
		AppStore: 29,
		Device: "D2303",
		DeviceLang: "es-ES",
		Language: 3,
		FacebookId: "",
		AdvertisingEnabled: 1,
		AdvertisingGuid: "cc823b32-3dbc-4455-8d00-f6b1ef6ad4b1",
		OsVersion: "7.1.2",
		IsAndroid: 1,
		AndroidId: "68e646598f73a6f5",
		MacAddress: [4]byte{0xff,0xff,0xff,0xff},
		AppleIfv: "",
		KunlunSso: "",
		KunlunUid: "",
		U16: 1,
		U13: "",
		U24: "",
		U25: "",
		U26: 0,
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
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.PassToken))
	binary.Write(buf, binary.BigEndian, utils.GetRrsInt32(o.MajorVersion))
	binary.Write(buf, binary.BigEndian, utils.GetRrsInt32(o.MinorVersion))
	binary.Write(buf, binary.BigEndian, utils.GetRrsInt32(o.Build))
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.ContentHash))
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.Udid))
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.OpenUdid))
	binary.Write(buf, binary.BigEndian, o.MacAddress)
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.Device))
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.AdvertisingGuid))
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.OsVersion))
	binary.Write(buf, binary.BigEndian, o.IsAndroid)
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.U13))
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.AndroidId))
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.DeviceLang))
	binary.Write(buf, binary.BigEndian, o.U16)
	binary.Write(buf, binary.BigEndian, o.Language)
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.FacebookId))
	binary.Write(buf, binary.BigEndian, o.AdvertisingEnabled)
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.AppleIfv))
	binary.Write(buf, binary.BigEndian, utils.GetRrsInt32(o.AppStore))
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.KunlunSso))
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.KunlunUid))
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.U24))
	utils.WriteBytes(buf, binary.BigEndian, []byte(o.U25))
	binary.Write(buf, binary.BigEndian, o.U26)

	return buf.Bytes()
}

// It represents the ServerLoginOk message received when we successfully logged in
type ServerLoginOk struct {
	Hi int32
	Lo int32
	PassToken string
	GameCenterId string
	FacebookId string
	MajorVersion int32
	Build int32
	ContentVersion int32
	Environment string
	SessionCount int32
	PlayTimeSeconds int32
	DaysSinceStartedPlaying int32
	FacebookAppId string
	ServerTime string
	AccountCreatedDate string
	U16 int32
	GoogleServiceId string
	U18 string
	U19 string
	Region string
	ContentURL string
	EventAssetsURL string
	U23 byte
}

func NewServerLoginOkFromBytes(buff []byte) (ServerLoginOk){
	o := ServerLoginOk{}

	var buf *bytes.Reader

	buf = bytes.NewReader(buff)
	binary.Read(buf, binary.BigEndian, &o.Hi)
	binary.Read(buf, binary.BigEndian, &o.Lo)
	binary.Read(buf, binary.BigEndian, &o.Hi) // it is repeated..
	binary.Read(buf, binary.BigEndian, &o.Lo)
	utils.ReadString(buf, binary.BigEndian, &o.PassToken)
	utils.ReadString(buf, binary.BigEndian, &o.GameCenterId)
	utils.ReadString(buf, binary.BigEndian, &o.FacebookId)
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.MajorVersion)
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.Build)
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.Build) // it is repeated..
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.ContentVersion)
	utils.ReadString(buf, binary.BigEndian, &o.Environment)
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.SessionCount)
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.PlayTimeSeconds)
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.DaysSinceStartedPlaying)
	utils.ReadString(buf, binary.BigEndian, &o.FacebookAppId)
	utils.ReadString(buf, binary.BigEndian, &o.ServerTime)
	utils.ReadString(buf, binary.BigEndian, &o.AccountCreatedDate)
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.U16)
	utils.ReadString(buf, binary.BigEndian, &o.GoogleServiceId)
	utils.ReadString(buf, binary.BigEndian, &o.U18)
	utils.ReadString(buf, binary.BigEndian, &o.U19)
	utils.ReadString(buf, binary.BigEndian, &o.Region)
	utils.ReadString(buf, binary.BigEndian, &o.ContentURL)
	utils.ReadString(buf, binary.BigEndian, &o.EventAssetsURL)
	binary.Read(buf, binary.BigEndian, &o.U23)

	return o
}

// It represents the ServerLoginFailed message
type ServerLoginFailed struct {
	ErrorCode int32
	ResourceFingerprintData string
	RedirectDomain string
	ContentURL string
	UpdateURL string
	Reason string
	SecondsUntilMantinanceEnd int32
}

func NewServerLoginFailedFromBytes(buff []byte) (ServerLoginFailed){
	o := ServerLoginFailed{}

	var buf *bytes.Reader
	//var tmp16 int16

	buf = bytes.NewReader(buff)
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.ErrorCode)
	utils.ReadString(buf, binary.BigEndian, &o.ResourceFingerprintData)
	utils.ReadString(buf, binary.BigEndian, &o.RedirectDomain)
	utils.ReadString(buf, binary.BigEndian, &o.ContentURL)
	utils.ReadString(buf, binary.BigEndian, &o.UpdateURL)
	//binary.Read(buf, binary.BigEndian, &tmp16)
	utils.ReadString(buf, binary.BigEndian, &o.Reason)
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.SecondsUntilMantinanceEnd)

	return o
}

