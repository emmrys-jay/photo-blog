package token

import (
	"errors"
	"time"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	Username  string
	IssuedAt  time.Time
	ExpiresAt time.Time
}

func NewPayload(username string) *Payload {
	return &Payload{
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(time.Minute * 5),
	}
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ErrExpiredToken
	}

	return nil
}
