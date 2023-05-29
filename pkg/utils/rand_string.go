package utils

import (
	"crypto/rand"
)

// RandString generate random string
func RandString(n int) (string, error) {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	for i := range b {
		b[i] = letters[b[i]%byte(len(letters))]
	}

	return string(b), nil
}
