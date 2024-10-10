package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// @Summary		System health-check
// @Description	Checks all depdent systems, if they are still up and responding
//
// @Tags			health
//
// @Produce		plain
//
// @Success		200	{string}	string							"✅ System healthy"
// @Failure		404	{object}	swagger.NotFoundErrorResponse	"code":	404,	"message":	"Not found ❌"
// @Failure		500	{string}	string							"❌ System unhealthy"
//
// @Router			/api/system-healthcheck [get]
func (h *GeneralHandler) SystemHealthCheck(c *fiber.Ctx) error {
	err := h.service.SystemHealthCheck()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("❌ System unhealthy")
	}
	return c.Status(http.StatusOK).SendString("✅ System healthy")
}
