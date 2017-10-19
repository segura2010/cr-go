package client

import (
	"net"
    //"fmt"

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
    var err error

    _, err = o.Socket.Read(buf[:7])
    if err != nil{
    	panic(err)
    }

    pkt := packets.NewPacketFromBytes(buf[:])

    if int(pkt.Length) > MAX_LENGTH{
        return pkt
    }

    _, err = o.Socket.Read(buf[:pkt.Length])
    if err != nil{
        panic(err)
    }
    pkt.Payload = buf[:pkt.Length]

    pkt = o.Crypt.DecryptPacket(pkt)

    return pkt
}