package token

import (
	"time"
)

type Maker interface {
	CreateToken(userID, username, name, email, bio, avatar string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
