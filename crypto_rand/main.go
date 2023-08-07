package main

import (
	"crypto/rand"
	"fmt"
)

const CHAR = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func random(length int) (string, error) {
	b := make([]byte, length)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	for i, itemB := range b {
		b[i] = CHAR[itemB%byte(len(CHAR))]
	}

	return string(b), nil
}

func main() {
	ran, err := random(32)

	if err != nil {
		panic(err)
	}

	fmt.Println(ran)
}
