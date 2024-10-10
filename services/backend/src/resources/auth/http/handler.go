package http

import (
	"github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service"
	jwtService "github.com/uvulpos/golang-sveltekit-template/src/resources/jwt/service"
)

type AuthHandler struct {
	service *service.AuthService
	jwtSvc  *jwtService.JwtService
}

func NewAuthHandler(service *service.AuthService, jwtSvc *jwtService.JwtService) *AuthHandler {
	return &AuthHandler{
		service: service,
		jwtSvc:  jwtSvc,
	}
}
