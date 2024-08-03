package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	logoutURL, logoutErr := h.service.Logout()
	if logoutErr != nil {
		return logoutErr
	}
	return c.Redirect(logoutURL, http.StatusMovedPermanently)
}
