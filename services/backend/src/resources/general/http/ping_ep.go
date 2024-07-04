package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

//	@Summary		Ping the server if startup was successful
//	@Description	get string by ID
//	@Tags			health
//
//	@Produce		plain
//
//	@Success		200	{string}	string							"API is online"
//	@Failure		404	{object}	swagger.NotFoundErrorResponse	"code":	404,	"message":	"Not found ❌"
//
//	@Router			/api/ping [get]
func (h *GeneralHandler) Ping(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("API is online")
}
