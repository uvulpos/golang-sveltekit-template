package service

import userService "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service"

type MiddlewareService struct {
	userSvc *userService.UserService
}

func NewMiddlewareService(userSvc *userService.UserService) *MiddlewareService {
	return &MiddlewareService{
		userSvc: userSvc,
	}
}
