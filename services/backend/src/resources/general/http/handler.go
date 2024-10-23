package http

import "github.com/uvulpos/golang-sveltekit-template/src/resources/general/service"

type GeneralHandler struct {
	service *service.GeneralSvc
}

func NewGeneralHandler(service *service.GeneralSvc) *GeneralHandler {
	return &GeneralHandler{
		service,
	}
}
