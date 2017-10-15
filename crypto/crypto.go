package crypto

import (
	"crypto/rand"
	"golang.org/x/crypto/nacl/box"
	"fmt"

    "github.com/segura2010/cr-go/packets"
)

type Crypto struct {
    PrivateKey [32]byte
    PublicKey [32]byte
    ServerKey []byte
    SharedKey [32]byte
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
		// hello packet is not encrypted
		hello := packets.NewServerHelloFromBytes(pkt.Payload)
		o.SessionKey = hello.SessionKey
		fmt.Printf("\nReceived sessionkey: %x", o.SessionKey)
	}else if pkt.Type == packets.MessageType["ServerLoginFailed"]{
		fmt.Printf("\nServerLoginFailed")
	}else if pkt.Type == packets.MessageType["ServerLoginOk"]{
		fmt.Printf("\nServerLoginFailed")
	}else{

	}

	return pkt
}

func (o *Crypto) EncryptPacket(pkt packets.Packet) (packets.Packet){
	if pkt.Type == packets.MessageType["ClientLogin"]{
		// Generate initial Nonce
		o.Nonce = NewNonce(o.PublicKey[:], o.ServerKey)
		message := append(o.SessionKey, o.Nonce.EncryptedNonce[:]...)
		message = append(message, pkt.DecryptedPayload...)
		
		var out []byte
		out,_ = box.OpenAfterPrecomputation(out, message, &o.Nonce.EncryptedNonce, &o.SharedKey)
		
		pkt.Payload = append(o.PublicKey[:], out...)
	}else{

	}

	return pkt
}