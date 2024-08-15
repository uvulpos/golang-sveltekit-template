package http

import (
	"fmt"
	"net/http"
	"time"

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

		/* ---------------------------------------------
				GDPR sensible information
		--------------------------------------------- */
		"", // c.IP(),              // 	for GDPR reasons disabled be default
		"", // c.Get("User-Agent"), // 	for GDPR reasons disabled be default

	)
	if err != nil || jwtToken == "" {
		fmt.Println("#ERROR JWT / CALLBACK", err)
		return err
	}

	// set jwt
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Path:     "/",
		Value:    jwtToken,
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
	})

	c.Redirect(configuration.WEBSERVER_DISPLAYNAME, http.StatusMovedPermanently)

	return nil
}
