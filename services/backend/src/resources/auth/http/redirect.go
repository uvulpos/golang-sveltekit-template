package http

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/go-svelte/basic-utils/customerrors"
)

func (h *AuthHandler) CreateRedirect(c *fiber.Ctx) error {
	var expiration = time.Now().Add(10 * time.Second)

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
		Expires: expiration,
	}
	c.Cookie(cookie)

	return c.Redirect(redirectUrl, http.StatusFound)
}
