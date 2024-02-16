package http

import "github.com/uvulpos/go-svelte/src/resources/passkeys-fido/service"

type PasskeyHandler struct {
	service *service.PasskeySvc
}

func NewPasskeyHandler(service *service.PasskeySvc) *PasskeyHandler {
	return &PasskeyHandler{
		service,
	}
}
