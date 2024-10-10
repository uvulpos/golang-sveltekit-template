package http

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Path:     "/",
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
		Expires:  time.Now(),
	})
	logoutURL, logoutErr := h.service.Logout()
	if logoutErr != nil {
		return logoutErr
	}
	return c.Redirect(logoutURL, http.StatusMovedPermanently)
}
