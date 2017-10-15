package crypto

import (
	"encoding/binary"
	"bytes"
	//"fmt"

	"github.com/codahale/blake2"
)

type CryptoNonce struct {
    EncryptedNonce []byte
}

func NewNonce(publicKey, serverKey []byte) (CryptoNonce){
	h := blake2.New(&blake2.Config{Size: 24})
	h.Write(publicKey)
	h.Write(serverKey)
	hash := h.Sum(nil)

	o := CryptoNonce{
		EncryptedNonce: hash,
	}
	return o
}

func NewNonceWithServerNonce(nonce []byte) (CryptoNonce){
	o := CryptoNonce{
		EncryptedNonce: nonce,
	}
	return o
}

func (o *CryptoNonce) Increment(){
	var n int16
	var tmp [22]byte
	// read as int16le and increment
	buf := bytes.NewReader(o.EncryptedNonce)
	binary.Read(buf, binary.LittleEndian, &n)
	binary.Read(buf, binary.LittleEndian, &tmp)
	n += 2
	// create new byte buffer to save the incremented value
	newbuf := new(bytes.Buffer)
	binary.Write(newbuf, binary.LittleEndian, &n)
	binary.Write(newbuf, binary.LittleEndian, &tmp)

	o.EncryptedNonce = newbuf.Bytes()
}