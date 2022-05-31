package token

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("myPersonal_secret_key")

type JWTMaker struct {
	secretKey []byte
}

func NewJWTMaker() *JWTMaker {
	return &JWTMaker{
		secretKey: secretKey,
	}
}

func (jm *JWTMaker) CreateToken(payload *Payload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(jm.secretKey))
}

func (jm *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(jm.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
