package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/golang-sveltekit-template/src/resources/auth/http/cookies"
	sessionLocals "github.com/uvulpos/golang-sveltekit-template/src/web-app/middlewares/http/consts/session-locals"
)

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	refreshTokenString := c.Locals(sessionLocals.AuthProvider)
	fmt.Println(refreshTokenString)

	c.Cookie(cookies.GenerateJwtToken("", true, time.Now()))
	c.Cookie(cookies.GenerateRefreshToken("", true, time.Now()))

	logoutURL, logoutErr := h.service.Logout()
	if logoutErr != nil {
		return logoutErr
	}
	return c.Redirect(logoutURL, http.StatusMovedPermanently)
}
