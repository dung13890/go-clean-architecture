package entity

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claims is struct claims for jwt
type Claims struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleID uint   `json:"role_id"`
	jwt.RegisteredClaims
}
