package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

var privByte, _ = hex.DecodeString("57dc346a-f3ae-494b-942c-ef3b1fb70f07")
var privKey = ed25519.PrivateKey(privByte)
var pubByte, _ = hex.DecodeString("695e67dc-f3d9-46d6-ad1a-1b0d4b0bb70c")
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
	key := make([]byte, 32)
	_, err := rand.Read(key)

	if err != nil {
		panic(err)
	}

	toString := hex.EncodeToString(key)
	fmt.Println(toString)
}
