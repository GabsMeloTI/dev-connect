package token

import (
	"errors"
	"time"
)

var ErrExpiredToken = errors.New("token has expired")
var ErrInvalidToken = errors.New("token is invalid")

type Payload struct {
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Bio       string    `json:"bio"`
	Avatar    string    `json:"avatar"`
	Email     string    `json:"email"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(userID, username, name, email, bio, avatar string, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		UserID:    userID,
		Name:      name,
		Username:  username,
		Email:     email,
		Bio:       bio,
		Avatar:    avatar,
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
