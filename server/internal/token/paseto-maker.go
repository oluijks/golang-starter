package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoTokenMaker struct {
	paseto      *paseto.V2
	symmeticKey []byte
}

func NewPasetoTokenMaker() *PasetoTokenMaker {
	return &PasetoTokenMaker{
		paseto:      &paseto.V2{},
		symmeticKey: []byte{},
	}
}

func PasetoMaker(symmeticKey string) (Maker, error) {
	if len(symmeticKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoTokenMaker{
		paseto:      paseto.NewV2(),
		symmeticKey: []byte(symmeticKey),
	}
	return maker, nil
}

func (maker *PasetoTokenMaker) MakeToken(email string, duration time.Duration) (string, error) {
	payload, err := NewPayload(email, duration)
	if err != nil {
		return "", err
	}
	return maker.paseto.Encrypt(maker.symmeticKey, payload, nil)
}

func (maker *PasetoTokenMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmeticKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
