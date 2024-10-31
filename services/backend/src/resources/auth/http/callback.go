package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/random"
	"github.com/uvulpos/golang-sveltekit-template/src/resources/auth/http/cookies"
)

func (h *AuthHandler) CallbackHandler(c *fiber.Ctx) error {
	authCode := c.Query("code", "")
	state := c.Query("state", "")
	authProvider := c.Cookies(cookies.CookieName_AuthProvider)

	jwtToken, refreshToken, err := h.service.CallbackFunction(
		authProvider,
		authCode,
		state,
	)

	if err != nil || jwtToken == "" {
		return err
	}

	c.Cookie(cookies.GenerateJwtToken(jwtToken, false))
	c.Cookie(cookies.GenerateRefreshToken(refreshToken, false))

	redirectURL, redirectURLErr := random.GenerateRandomHashURL(configuration.WEBSERVER_DISPLAYNAME)
	if redirectURLErr != nil {
		status, _, usermessage := redirectURLErr.HttpError()
		return c.Status(status).SendString(usermessage)
	}

	return c.Redirect(redirectURL, http.StatusMovedPermanently)
}
