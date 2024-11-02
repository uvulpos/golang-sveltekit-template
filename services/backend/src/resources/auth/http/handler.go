package http

import (
	"github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}
