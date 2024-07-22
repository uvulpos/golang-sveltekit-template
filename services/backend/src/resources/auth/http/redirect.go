package http

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) CreateRedirect(c *fiber.Ctx) error {
	var expiration = time.Now().Add(10 * time.Second)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	redirectUrl := h.service.CreateRedirect(state)

	cookie := &fiber.Cookie{
		Name:    "oauthstate",
		Value:   state,
		Expires: expiration,
	}
	c.Cookie(cookie)
	c.Redirect(redirectUrl, http.StatusFound)
	return nil
}
