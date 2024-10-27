package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/golang-sveltekit-template/src/resources/auth/http/cookies"
)

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(cookies.GenerateRefreshToken("", true))
	c.Cookie(cookies.GenerateJwtToken("", true))

	return c.Redirect("/", http.StatusMovedPermanently)
}
