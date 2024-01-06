package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) HandleLogout(c *fiber.Ctx) error {
	c.ClearCookie("jwt")
	return c.SendStatus(http.StatusOK)
}
