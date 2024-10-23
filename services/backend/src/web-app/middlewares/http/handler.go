package middlewares

import (
	jwtService "github.com/uvulpos/golang-sveltekit-template/src/resources/jwt/service"
	middlewareService "github.com/uvulpos/golang-sveltekit-template/src/web-app/middlewares/service"
)

type MiddlewareHandler struct {
	service *middlewareService.MiddlewareService
	jwtSvc  *jwtService.JwtService
}

func NewMiddlewareHandler(service *middlewareService.MiddlewareService, jwtSvc *jwtService.JwtService) *MiddlewareHandler {
	return &MiddlewareHandler{
		service,
		jwtSvc,
	}
}
