package main

import "crypto/rand"

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
