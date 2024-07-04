package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

//	@Summary		Ping database
//	@Description	Ping the database if connection was successful.
//	@Description	Due to configuration, this ep might not be available
//
//	@Tags			health
//
//	@Produce		plain
//
//	@Success		200	{string}	string							"Database is connected"
//	@Failure		404	{object}	swagger.NotFoundErrorResponse	"code":	404,	"message":	"Not found ❌"
//	@Failure		500	{string}	string							"Database is not ready"
//
//	@Router			/api/ping-db [get]
func (h *GeneralHandler) PingDatabase(c *fiber.Ctx) error {
	err := h.service.Ping()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Database is not ready")
	}
	return c.Status(http.StatusOK).SendString("Database is connected")
}
