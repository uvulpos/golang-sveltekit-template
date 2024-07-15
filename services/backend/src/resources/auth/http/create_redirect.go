package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) CreateRedirect(c *fiber.Ctx) error {
	redirectUrl := h.auth.CreateRedirect()
	c.Redirect(redirectUrl, http.StatusFound)
	return nil
}
