package utils

import (
	"go-app/config"
	"go-app/internal/domain"
	"go-app/pkg/logger"
	"time"

	"github.com/golang-jwt/jwt"

	"golang.org/x/crypto/bcrypt"
)

// JwtKey is secret key fow singed
var JwtKey = []byte(config.GetAppConfig().AppJWTKey)

// Claims is struct claims for jwt
type Claims struct {
	Name   string `json:"name"`
	ID     uint   `json:"id"`
	RoleID int    `json:"role_id"`
	jwt.StandardClaims
}

// GeneratePassword returns hashed password
func GeneratePassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	return string(bytes), err
}

// ComparePassword used to compare password with hashed password
func ComparePassword(pass string, hashPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(pass)) == nil
}

// GenerateToken returns token string
func GenerateToken(user *domain.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		Name:   user.Name,
		ID:     user.ID,
		RoleID: user.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(JwtKey)
	if err != nil {
		logger.Error().Println("error when new claims: ", err)

		return "", err
	}

	return tokenString, nil
}
