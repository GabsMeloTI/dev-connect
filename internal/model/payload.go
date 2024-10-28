package model

import (
	"time"
)

type PayloadDTO struct {
	UserID   string    `json:"user_id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	ExpiryAt time.Time `json:"expiry_at"`
}
