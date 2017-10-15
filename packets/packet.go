package packets



type Packet struct {
    Type uint16
    Length int32
    Version uint16
    Payload []byte // the encrypted message
    DecryptedPayload []byte // the message itself
}


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
