package http

import (
	"github.com/uvulpos/golang-sveltekit-template/src/resources/user/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
