package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
)

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	sessionID := c.Locals("session-uuid").(string)

	// to go to database and check if session is still valid

	jwtToken, refreshTokenErr := h.service.RecreateJwtFromSession(sessionID)
	if refreshTokenErr != nil {
		return refreshTokenErr
	}

	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Path:    "/",
		Value:   jwtToken,
		Expires: time.Now().Add(time.Minute * time.Duration(configuration.JWT_TOKEN_VALIDITY_IN_MINUTES)),
	})

	return c.SendString(jwtToken)
}
