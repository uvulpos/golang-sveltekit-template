package http

import (
	jwtService "github.com/uvulpos/golang-sveltekit-template/src/helper/jwt"
	"github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service"
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
