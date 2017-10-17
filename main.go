package main

import (
    "fmt"
    "encoding/hex"

    //"io/ioutil"

    "github.com/segura2010/cr-go/client"
    "github.com/segura2010/cr-go/packets"
    "github.com/segura2010/cr-go/utils"
)


func main(){
    fmt.Printf("Welcome to cr-go!")

    serverKey_202 := "72f1a4a4c48e44da0c42310f800e96624e6dc6a641a9d41c3b5039d8dfadc27e" // node-proxy
    //serverKey_202 := "980cf7bb7262b386fea61034aba7370613627919666b34e6ecf66307a381dd61"
    serverKey,_ := hex.DecodeString(serverKey_202)

    helloPayload := packets.NewDefaultClientHello()
    helloPkt := packets.Packet{
        Type: packets.MessageType["ClientHello"],
        Version: 0,
        Payload: helloPayload.Bytes(),
    }

    loginPayload := packets.NewDefaultClientLogin()
    HiLo := utils.Tag2HiLo("") // your account
    loginPayload.Hi = HiLo[0]
    loginPayload.Lo = HiLo[1]
    loginPayload.PassToken = "" // your account PassToken

    loginPkt := packets.Packet{
        Type: packets.MessageType["ClientLogin"],
        Version: 0,
        DecryptedPayload: loginPayload.Bytes(),
    }

    var basicPkt packets.Packet

    serverAddress := "0.0.0.0:9339"
    //serverAddress := "game.clashroyaleapp.com:9339"

    c := client.NewCRClient(serverKey)
    c.Connect(serverAddress)
    defer c.Close()

    c.SendPacket(helloPkt)
    c.RecvPacket() // receive hello response

    c.SendPacket(loginPkt)
    basicPkt = c.RecvPacket() // receive login response

    // receive multiple packets the server sends after login
    c.RecvPacket() // OwnHomeData
    c.RecvPacket() // InboxGlobal
    c.RecvPacket() // FriendList

    // send keepalive
    basicPkt = packets.Packet{
        Type: packets.MessageType["ClientKeepAlive"],
        Version: 0,
    }
    c.SendPacket(basicPkt)
    c.RecvPacket() // receive keepalive response

    // send visithome
    HiLo = utils.Tag2HiLo("P0C9QP8L")
    visitHomeMsg := packets.ClientVisitHome{
        Hi: HiLo[0],
        Lo: HiLo[1],
    }
    basicPkt = packets.Packet{
        Type: packets.MessageType["ClientVisitHome"],
        Version: 0,
        DecryptedPayload: visitHomeMsg.Bytes(),
    }
    c.SendPacket(basicPkt)
    basicPkt = c.RecvPacket() // receive visithome response
    fmt.Printf("\nVisitedHome: %x", basicPkt.DecryptedPayload)
    
}

