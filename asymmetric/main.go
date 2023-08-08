package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

var privByte, _ = hex.DecodeString("275d307f075dcec1be9934a3d8fea24d62f1fcb445d7d44cb4d632a892283e50ece3386d8a511a903b7208d828059d0177fe66f98c9fa1eed562dd3b3846bb8f")
var privKey = ed25519.PrivateKey(privByte)
var pubByte, _ = hex.DecodeString("33b311fae5d4f4631f511bab8906afb8e6346a30df25e630b70fb4e8557b69e6")
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
	pubkey := make([]byte, 32)
	privKey := make([]byte, 64)

	_, err := rand.Read(pubkey)

	if err != nil {
		panic(err)
	}

	_, err = rand.Read(privKey)

	if err != nil {
		panic(err)
	}

	encodePub := hex.EncodeToString(pubkey)
	encodePriv := hex.EncodeToString(privKey)
	token := Encode()

	fmt.Printf("Public Key: %v \n", encodePub)
	fmt.Printf("Private Key: %v \n", encodePriv)
	fmt.Println("Token:", token)

}
