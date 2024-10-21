package model

import (
	"time"
)

type PayloadDTO struct {
	UserID       string    `json:"user_id"`
	UserNickname string    `json:"user_nickname"`
	ExpiryAt     time.Time `json:"expiry_at"`
	Email        string    `json:"email"`
}
