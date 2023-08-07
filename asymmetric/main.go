package main

import (
	"crypto/ed25519"
	"fmt"
)

func main() {
	publicKey, privateKey, _ := ed25519.GenerateKey(nil)

	pub := []byte(publicKey)
	p := []byte(privateKey)

	fmt.Printf("Public Key: %v \n", string(pub))
	fmt.Printf("Private Key: %v \n", string(p))

}
