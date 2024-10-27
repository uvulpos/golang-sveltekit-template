package http

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	"github.com/uvulpos/golang-sveltekit-template/src/resources/auth/http/cookies"
)

func (h *AuthHandler) CreateRedirect(c *fiber.Ctx) error {
	var expiration = time.Now().Add(10 * time.Minute)

	b := make([]byte, 16)
	_, readErr := rand.Read(b)
	if readErr != nil {
		httpStatus, _, httpMessage := customerrors.NewInternalServerError(readErr, "", "Could not create login hash").HttpError()
		return c.Status(httpStatus).SendString(httpMessage)
	}

	state := base64.URLEncoding.EncodeToString(b)
	redirectUrl := h.service.CreateRedirect(state)

	cookie := &fiber.Cookie{
		Name:    "oauthstate",
		Value:   state,
		Path:    "/api/v1/oauth",
		Expires: expiration,
	}
	c.Cookie(cookie)

	c.Cookie(cookies.GenerateAuthProviderCookie("authentik", false))

	return c.Redirect(redirectUrl, http.StatusPermanentRedirect)
}
