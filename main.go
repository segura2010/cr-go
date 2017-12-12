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

    //serverKey_202 := "72f1a4a4c48e44da0c42310f800e96624e6dc6a641a9d41c3b5039d8dfadc27e" // node-proxy
    serverKey_202 := "99b61876f3ff18caeca0aec1f326d9981bbcaf64e7daa317a7f10966867af968"
    serverKey,_ := hex.DecodeString(serverKey_202)

    helloPayload := packets.NewDefaultClientHello()
    helloPayload.ContentHash = "3f6063a89e32b2403462696cbb5e0f68f8e58ea2"
    helloPkt := packets.Packet{
        Type: packets.MessageType["ClientHello"],
        Version: 0,
        Payload: helloPayload.Bytes(),
    }

    loginPayload := packets.NewDefaultClientLogin()
    HiLo := utils.Tag2HiLo("") // test account
    loginPayload.Hi = HiLo[0]
    loginPayload.Lo = HiLo[1]
    loginPayload.PassToken = ""
    loginPayload.ContentHash = "3f6063a89e32b2403462696cbb5e0f68f8e58ea2"

    loginPkt := packets.Packet{
        Type: packets.MessageType["ClientLogin"],
        Version: 0,
        DecryptedPayload: loginPayload.Bytes(),
    }

    var basicPkt packets.Packet

    //serverAddress := "0.0.0.0:9339"
    serverAddress := "game.clashroyaleapp.com:9339"

    c := client.NewCRClient(serverKey)
    c.Connect(serverAddress)
    defer c.Close()

    c.SendPacket(helloPkt)
    basicPkt = c.RecvPacket() // receive hello response
    fmt.Printf("\nOn ServerHello I received: %d", basicPkt.Type)
    fmt.Printf("\nPayload: %x", basicPkt.Payload)
    if basicPkt.Type == packets.MessageType["ServerLoginFailed"]{
        loginFailed := packets.NewServerLoginFailedFromBytes(basicPkt.DecryptedPayload)
        fmt.Printf("\nLoginFailedOnHello: %v", loginFailed)
        fmt.Printf("\nLoginFailedOnHello: %s", basicPkt.Payload)
        return
    }

    c.SendPacket(loginPkt)
    basicPkt = c.RecvPacket() // receive login response
    fmt.Printf("\nOn LoginOk I received: %d", basicPkt.Type)
    fmt.Printf("\nPayload: %x", basicPkt.Payload)
    if basicPkt.Type == packets.MessageType["ServerLoginFailed"]{
        loginFailed := packets.NewServerLoginFailedFromBytes(basicPkt.DecryptedPayload)
        fmt.Printf("\nLoginFailed: %v", loginFailed)
        return
    }

    loginOk := packets.NewServerLoginOkFromBytes(basicPkt.DecryptedPayload)
    fmt.Printf("\nLoginOk: %v", loginOk)

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
}

