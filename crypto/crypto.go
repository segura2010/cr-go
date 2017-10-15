package crypto

import (
	"crypto/rand"
	"golang.org/x/crypto/nacl/box"

    "github.com/segura2010/cr-go/packets"
)

type Crypto struct {
    PrivateKey [32]byte
    PublicKey [32]byte
    ServerKey []byte
    SharedKey []byte
    SessionKey []byte
    Nonce CryptoNonce
}

func NewCrypto(serverKey []byte) (Crypto){

	publicKey, privateKey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	o := Crypto{
		ServerKey: serverKey,
		PublicKey: *publicKey,
		PrivateKey: *privateKey,
	}

	return o
}

func (o *Crypto) DecryptPacket(pkt packets.Packet) (packets.Packet){
	if pkt.Type == packets.MessageType["ServerHello"]{

	}else if pkt.Type == packets.MessageType["ServerLoginFailed"]{

	}else if pkt.Type == packets.MessageType["ServerLoginOk"]{

	}else{

	}

	return pkt
}

func (o *Crypto) EncryptPacket(pkt packets.Packet) (packets.Packet){
	if pkt.Type == packets.MessageType["ClientLogin"]{
		// Generate initial Nonce
		o.Nonce = NewNonce(o.PublicKey[:], o.ServerKey)
	}else{

	}

	return pkt
}