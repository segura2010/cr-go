package crypto

import (
	"crypto/rand"
	"golang.org/x/crypto/nacl/box"
	"fmt"
	"bytes"
	"encoding/binary"

    "github.com/segura2010/cr-go/packets"
)

// it implements all the needed methods to encrypt and decrypt the messages
type Crypto struct {
    PrivateKey [32]byte
    PublicKey [32]byte
    ServerKey [32]byte
    SharedKey [32]byte
    SessionKey []byte
    DecryptionNonce CryptoNonce
    EncryptionNonce CryptoNonce
}

func NewCrypto(serverKey []byte) (Crypto){

	publicKey, privateKey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	o := Crypto{
		PublicKey: *publicKey,
		PrivateKey: *privateKey,
	}

	copy(o.ServerKey[:], serverKey[:32])

	return o
}

func (o *Crypto) DecryptPacket(pkt packets.Packet) (packets.Packet){
	if pkt.Type == packets.MessageType["ServerHello"]{
		// hello packet is not encrypted
		hello := packets.NewServerHelloFromBytes(pkt.Payload)
		o.SessionKey = hello.SessionKey
		//fmt.Printf("\nServerHello")
		//fmt.Printf("\nReceived sessionkey: %x", o.SessionKey)
	}else if pkt.Type == packets.MessageType["ServerLoginFailed"]{
		//fmt.Printf("\nServerLoginFailed")
	}else if pkt.Type == packets.MessageType["ServerLoginOk"]{
		//fmt.Printf("\nServerLoginOK")
		tmpNonce := NewNonceWithNonce(o.PublicKey[:], o.ServerKey[:], o.EncryptionNonce.EncryptedNonce[:])
		out, decrypted := box.OpenAfterPrecomputation(nil, pkt.Payload, &tmpNonce.EncryptedNonce, &o.SharedKey)
		
		//fmt.Printf("\n\n%x", out)

		if decrypted{
			var nonce [24]byte
			var sharedKey [32]byte

			buf := bytes.NewReader(out)

			binary.Read(buf, binary.BigEndian, &nonce)
			binary.Read(buf, binary.BigEndian, &sharedKey)

			o.DecryptionNonce = NewNonceWithServerNonce(nonce[:])
			o.SharedKey = sharedKey

			//fmt.Printf("\nReceived DecryptionNonce: %x", o.DecryptionNonce.EncryptedNonce[:])
			//fmt.Printf("\nReceived SharedKey: %x", o.SharedKey)

			pkt.DecryptedPayload = out[56:] // remove nonce and sharedKey
		}
	}else{
		o.DecryptionNonce.Increment()

		// fmt.Printf("\n\n--- Decrypt ---")
		// fmt.Printf("\nMessage: %x", pkt.Payload)
		// fmt.Printf("\nNonce: %x", o.DecryptionNonce.EncryptedNonce[:])
		// fmt.Printf("\nSharedKey: %x", o.SharedKey[:])
		// fmt.Printf("\n--- /Decrypt ---")
		
		out, decrypted := box.OpenAfterPrecomputation(nil, pkt.Payload, &o.DecryptionNonce.EncryptedNonce, &o.SharedKey)
		pkt.DecryptedPayload = out
		if !decrypted{
			fmt.Printf("\nNot decrypted!")
		}
	}

	return pkt
}

func (o *Crypto) EncryptPacket(pkt packets.Packet) (packets.Packet){
	if pkt.Type == packets.MessageType["ClientHello"]{
		// Client hello is uncrypted
	}else if pkt.Type == packets.MessageType["ClientLogin"]{
		// Generate initial Nonce
		o.EncryptionNonce = NewNonce(o.PublicKey[:], o.ServerKey[:])
		//fmt.Printf("\n\nEncryptionNonce: %x", o.EncryptionNonce.EncryptedNonce[:])
		// generate initial sharedkey
		box.Precompute(&o.SharedKey, &o.ServerKey, &o.PrivateKey)

		// generate message [sessionKey][nonce][payload]
		message := append(o.SessionKey, o.EncryptionNonce.EncryptedNonce[:]...)
		message = append(message, pkt.DecryptedPayload...)

		// encrypt appending the client public key to the encrypted message
		out := box.SealAfterPrecomputation(o.PublicKey[:], message, &o.EncryptionNonce.EncryptedNonce, &o.SharedKey)
		pkt.Payload = out
	}else{
		o.EncryptionNonce.Increment()

		// fmt.Printf("\n\n--- Encrypt ---")
		// fmt.Printf("\nMessage: %x", pkt.Payload)
		// fmt.Printf("\nNonce: %x", o.EncryptionNonce.EncryptedNonce[:])
		// fmt.Printf("\nSharedKey: %x", o.SharedKey[:])
		// fmt.Printf("\n--- /Encrypt ---")

		out := box.SealAfterPrecomputation(nil, pkt.DecryptedPayload, &o.EncryptionNonce.EncryptedNonce, &o.SharedKey)
		pkt.Payload = out
	}

	return pkt
}