package http

import "github.com/uvulpos/go-svelte/src/resources/users/service"

type UserHandler struct {
	service *service.UserSvc
}

func NewUserHandler(service *service.UserSvc) *UserHandler {
	return &UserHandler{
		service,
	}
}
