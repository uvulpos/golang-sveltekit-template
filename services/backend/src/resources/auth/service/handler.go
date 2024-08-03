package service

import (
	authService "github.com/uvulpos/go-svelte/authentication-api/ressources/auth/service"
	jwtService "github.com/uvulpos/go-svelte/authentication-api/ressources/jwt/service"
)

type AuthService struct {
	auth *authService.AuthService
	jwt  *jwtService.JwtService
}

func NewAuthService(authService *authService.AuthService, jwtService *jwtService.JwtService) *AuthService {
	return &AuthService{
		auth: authService,
		jwt:  jwtService,
	}
}
