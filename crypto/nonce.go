package crypto

import (
	"encoding/binary"
	"bytes"
	//"fmt"

	"github.com/codahale/blake2"
)

// it implements and represents the nonce used to encrypt and decrypt the packets
type CryptoNonce struct {
    EncryptedNonce [24]byte
}

// it creates a new nonce object based on the generated public key and the server key
func NewNonce(publicKey, serverKey []byte) (CryptoNonce){
	h := blake2.New(&blake2.Config{Size: 24})
	h.Write(publicKey)
	h.Write(serverKey)
	hash := h.Sum(nil)

	o := CryptoNonce{}
	copy(o.EncryptedNonce[:], hash[:24])
	return o
}

// it creates a new nonce object based on the generated public key, the server key and the received nonce from the server
func NewNonceWithNonce(publicKey, serverKey, nonce []byte) (CryptoNonce){
	h := blake2.New(&blake2.Config{Size: 24})
	h.Write(nonce)
	h.Write(publicKey)
	h.Write(serverKey)
	hash := h.Sum(nil)

	o := CryptoNonce{}
	copy(o.EncryptedNonce[:], hash[:24])
	return o
}

// it creates a new nonce object based on the received nonce from the server
func NewNonceWithServerNonce(nonce []byte) (CryptoNonce){
	o := CryptoNonce{}
	copy(o.EncryptedNonce[:], nonce[:24])
	return o
}

// it increments the nonce (we have to increment it every time we encrypt or decrypt a new packet)
func (o *CryptoNonce) Increment(){
	var n int16
	var tmp [22]byte
	// read as int16le and increment
	buf := bytes.NewReader(o.EncryptedNonce[:])
	binary.Read(buf, binary.LittleEndian, &n)
	binary.Read(buf, binary.LittleEndian, &tmp)

	n = n + 2
	//n = n % 32767

	// create new byte buffer to save the incremented value
	newbuf := new(bytes.Buffer)
	binary.Write(newbuf, binary.LittleEndian, &n)
	binary.Write(newbuf, binary.LittleEndian, &tmp)

	copy(o.EncryptedNonce[:], newbuf.Bytes()[:24])
}

