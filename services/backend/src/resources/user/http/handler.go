package http

import (
	"github.com/uvulpos/go-svelte/src/resources/user/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
