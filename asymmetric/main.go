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
var symmetricKey, _ = hex.DecodeString("84e79d7b-e588-4978-a67f-f577bd39fb3d")

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

func Symmetric() string {
	now := time.Now()
	exp := now.Add(24 * time.Hour)

	jsonToken := paseto.JSONToken{
		Expiration: exp,
	}

	jsonToken.Set("data", "Kumalala")

	token, err := paseto.NewV1().Encrypt(symmetricKey, jsonToken, "Some footer")

	if err != nil {
		panic(err)
	}

	return token

}

func main() {
	token := Symmetric()
	fmt.Println("Token:", token)

}
