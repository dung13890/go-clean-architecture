package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword returns hashed password
func GeneratePassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	return string(bytes), err
}

// ComparePassword used to compare new password with hashed password
func ComparePassword(pass string, hashPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(pass)) == nil
}
