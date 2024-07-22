package http

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/go-svelte/src/configuration"
)

func (h *AuthHandler) CallbackHandler(c *fiber.Ctx) error {
	authCode := c.Query("code", "")
	state := c.Query("state", "")

	oauthstate := c.Cookies("oauthstate", "")
	if state != oauthstate {
		return fmt.Errorf("not the same oauth session")
	}

	jwtToken, err := h.service.CallbackFunction(
		authCode,
		state,
		"",
	)
	if err != nil || jwtToken == "" {
		fmt.Println("#ERROR JWT / CALLBACK", err)
		return err
	}

	// set jwt
	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: jwtToken,
	})

	c.Redirect(configuration.WEBSERVER_DISPLAYNAME, http.StatusMovedPermanently)

	return nil
}
