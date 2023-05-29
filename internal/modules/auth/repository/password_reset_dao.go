package repository

import (
	"time"
)

// PasswordReset DAO model
type PasswordReset struct {
	Email     string `json:"email" gorm:"primaryKey"`
	Token     string `json:"token"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
