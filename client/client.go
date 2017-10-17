package client

import (
	"net"

    "github.com/segura2010/cr-go/crypto"
    "github.com/segura2010/cr-go/packets"
)

type CRClient struct {
    Crypt crypto.Crypto
    Socket net.Conn 
}

func NewCRClient(serverKey []byte) (CRClient){

	o := CRClient{}

	o.Crypt = crypto.NewCrypto(serverKey)

	return o
}

// It starts the connection with the server
func (o *CRClient) Connect(address string){
	var err error
	o.Socket, err = net.Dial("tcp", address)
    if err != nil{
        panic(err)
    }
}

// It closes the connection (we recommend you use this with defer)
func (o *CRClient) Close(){
	o.Socket.Close()
}

// It encrypts and sends the packet
func (o *CRClient) SendPacket(pkt packets.Packet) (packets.Packet){

	pkt = o.Crypt.EncryptPacket(pkt)

	o.Socket.Write(pkt.Bytes())

	return pkt
}

// It receives and decrypts the received packet
func (o *CRClient) RecvPacket() (packets.Packet){

	var MAX_LENGTH = 20000
    var buf [20000]byte
    var n int // bytes read
    var err error

    n, err = o.Socket.Read(buf[:MAX_LENGTH])
    if err != nil{
    	panic(err)
    }

    pkt := packets.NewPacketFromBytes(buf[:], n)
    pkt = o.Crypt.DecryptPacket(pkt)

    return pkt
}