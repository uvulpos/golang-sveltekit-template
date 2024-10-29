package http

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	"github.com/uvulpos/golang-sveltekit-template/src/resources/auth/http/cookies"
	providerConst "github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service/provider-const"
)

func (h *AuthHandler) CreateRedirect(c *fiber.Ctx) error {

	b := make([]byte, 16)
	_, readErr := rand.Read(b)
	if readErr != nil {
		httpStatus, _, httpMessage := customerrors.NewInternalServerError(readErr, "", "Could not create login hash").HttpError()
		return c.Status(httpStatus).SendString(httpMessage)
	}

	state := base64.URLEncoding.EncodeToString(b)
	redirectUrl := h.service.CreateRedirect(state)

	c.Cookie(cookies.GenerateOauthStateCookie(state, false))
	c.Cookie(cookies.GenerateAuthProviderCookie(string(providerConst.Authentik), false))

	return c.Redirect(redirectUrl, http.StatusTemporaryRedirect)
}
