package main

import (
    "fmt"
    "encoding/hex"

    "github.com/segura2010/cr-go/crypto"
)


func main(){
    fmt.Printf("Welcome to cr-go!")

    serverKey_202 := "980cf7bb7262b386fea61034aba7370613627919666b34e6ecf66307a381dd61"
    serverKey,_ := hex.DecodeString(serverKey_202)
    c := crypto.NewCrypto(serverKey)

    fmt.Printf("\nServerKey: %x", c.ServerKey)
    fmt.Printf("\nPrivateKey: %x", c.PrivateKey)
    fmt.Printf("\nPublicKey: %x", c.PublicKey)

    testPublicKey,_ := hex.DecodeString("4af8c34a51ac604f2ccb3ea249d4ed0e70f6cc2f39b70f6a5b50bcaaf5ca3f24")
    fmt.Printf("\nTest PublicKey: %x", testPublicKey)
    nonce := crypto.NewNonce(testPublicKey, serverKey)
    fmt.Printf("\nNonce: %x", nonce.EncryptedNonce)
}

