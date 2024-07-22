package service

import (
	authService "github.com/uvulpos/go-svelte/authentication-api/ressources/auth/service"
	jwtService "github.com/uvulpos/go-svelte/authentication-api/ressources/jwt/service"
	userService "github.com/uvulpos/go-svelte/user-api/ressources/user/service"
)

type AuthService struct {
	auth *authService.AuthService
	jwt  *jwtService.JwtService
	user *userService.UserService
}

func NewAuthService(authService *authService.AuthService, jwtService *jwtService.JwtService) *AuthService {
	return &AuthService{
		auth: authService,
		jwt:  jwtService,
	}
}
