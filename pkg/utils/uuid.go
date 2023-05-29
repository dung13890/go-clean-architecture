package utils

import (
	"github.com/google/uuid"
)

// GenerateUUID returns uuid from google
func GenerateUUID() string {
	return uuid.New().String()
}
