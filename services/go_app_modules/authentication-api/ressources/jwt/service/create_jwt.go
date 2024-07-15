package service

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

func (s *JwtService) CreateJWT() {
	// todo create jwt
	ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
}
