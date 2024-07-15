package http

import (
	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) GetAuthCallbackFunction(c *fiber.Ctx) error {
	authCode := c.Query("code", "")
	h.auth.CallbackFunction(authCode)
	return nil
}
