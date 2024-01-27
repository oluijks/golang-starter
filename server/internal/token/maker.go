package token

import "time"

type Maker interface {
	MakeToken(email string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
