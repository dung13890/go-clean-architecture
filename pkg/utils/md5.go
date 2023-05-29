package utils

import (
	"crypto/md5" //#nosec
	"encoding/hex"
)

// MD5Hash is a function to hash the string using MD5 algorithm
func MD5Hash(s string) string {
	hash := md5.Sum([]byte(s)) //#nosec

	return hex.EncodeToString(hash[:])
}
