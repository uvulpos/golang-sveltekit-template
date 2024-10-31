package middlewares

import (
	middlewareService "github.com/uvulpos/golang-sveltekit-template/src/web-app/middlewares/service"
)

type MiddlewareHandler struct {
	service *middlewareService.MiddlewareService
}

func NewMiddlewareHandler(service *middlewareService.MiddlewareService) *MiddlewareHandler {
	return &MiddlewareHandler{
		service,
	}
}
