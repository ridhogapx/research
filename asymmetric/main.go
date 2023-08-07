package main

import (
	"crypto/ed25519"
	"fmt"
)

func main() {
	publicKey, privateKey, _ := ed25519.GenerateKey(nil)

	fmt.Printf("Public Key: %v", publicKey)
	fmt.Printf("Private Key: %v", privateKey)

}
