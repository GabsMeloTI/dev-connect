package token

import (
	"errors"
	"time"
)

var ErrExpiredToken = errors.New("token has expired")
var ErrInvalidToken = errors.New("token is invalid")

type Payload struct {
	Username  string    `json:"username"`
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (payload *Payload) valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
