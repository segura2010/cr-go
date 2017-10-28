package client

import (
	"net"

    "github.com/segura2010/cr-go/crypto"
    "github.com/segura2010/cr-go/packets"
)

// it implements the communication with the server (using sockets) and the calls to the encryption/decryption
// functions every time we send/receive a packet
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
    var n int

    _, err = o.Socket.Read(buf[:7])
    if err != nil{
    	panic(err)
    }

    pkt := packets.NewPacketFromBytes(buf[:])

    if int(pkt.Length) > MAX_LENGTH{
        // bad length... ???
        pkt.Length = int32(MAX_LENGTH)
        n, err = o.Socket.Read(buf[:pkt.Length])
        if err != nil{
            panic(err)
        }
        pkt.Length = int32(n)
        pkt.Payload = buf[:pkt.Length]
    }else{
        // the packet could be sent in multiple chunks, so I have to
        // read until I read the full packet
        total := 0
        var tmpbuf []byte
        for{
            // read until we read the full packet
            diff := int(pkt.Length) - total
            n, err = o.Socket.Read(buf[:diff])
            if err != nil{
                panic(err)
            }
            total += n
            tmpbuf = append(tmpbuf, buf[:n]...)
            if int32(total) >= pkt.Length{
                break
            }
        }
        pkt.Payload = tmpbuf
    }

    pkt = o.Crypt.DecryptPacket(pkt)

    return pkt
    
}