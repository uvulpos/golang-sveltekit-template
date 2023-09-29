package http

import (
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) HandleLogout(c *fiber.Ctx) error {
	c.ClearCookie("jwt")
	return nil
}
