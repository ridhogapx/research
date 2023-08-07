package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

var privByte, _ = hex.DecodeString("b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZWQyNTUxOQAAACCX0Dw8NavIwFDX1811Tn4+35UNJT3qLdAkcbeZiyErYgAAAKh+4OANfuDgDQAAAAtzc2gtZWQyNTUxOQAAACCX0Dw8NavIwFDX1811Tn4+35UNJT3qLdAkcbeZiyErYgAAAED9XNY2RHnAGqv6gzGJ71SyjB5JWbiPxwJF5Rluc5nPupfQPDw1q8jAUNfXzXVOfj7flQ0lPeot0CRxt5mLIStiAAAAInJpZGhvZ2FweEByaWRob2dhcHgtQXNwaXJlLVIzLTEzMVQBAgM=")
var privKey = ed25519.PrivateKey(privByte)
var pubByte, _ = hex.DecodeString("AAAAC3NzaC1lZDI1NTE5AAAAIJfQPDw1q8jAUNfXzXVOfj7flQ0lPeot0CRxt5mLISti")
var pubKey = ed25519.PublicKey(pubByte)

func Encode() string {
	payload := paseto.JSONToken{
		Expiration: time.Now().Add(24 * time.Hour),
		Subject:    "Research",
	}

	payload.Set("data", "Life is Roblox!")

	footer := "Copyright"

	token, err := paseto.NewV2().Sign(privKey, payload, footer)

	if err != nil {
		panic(err)
	}
	return token
}

func main() {
	pubRaw := "AAAAC3NzaC1lZDI1NTE5AAAAIJfQPDw1q8jAUNfXzXVOfj7flQ0lPeot0CRxt5mLISti"
	fmt.Printf("Raw Public: %v", len(pubRaw))
	fmt.Printf("Public Key Length: %v \n", len(pubKey))
	fmt.Printf("Public Key Length: %v \n", len(privKey))

}
