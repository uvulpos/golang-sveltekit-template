package http

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/go-svelte/basic-utils/customerrors"
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

	redirectURL, redirectURLErr := generateRandomHashURL(configuration.WEBSERVER_DISPLAYNAME)
	if redirectURLErr != nil {
		status, _, usermessage := redirectURLErr.HttpError()
		return c.Status(status).SendString(usermessage)
	}

	return c.Redirect(redirectURL, http.StatusMovedPermanently)
}

func generateRandomHashURL(baseURL string) (string, customerrors.ErrorInterface) {
	hash, err := generateRandomHash()
	if err != nil {
		return "", customerrors.NewInternalServerError(err, "", "Could not generate a simple oauth redirect get parameter hash")
	}

	// URL parsen
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return "", customerrors.NewInternalServerError(err, "", "Couldn't parse URL in oauth callback")
	}

	fmt.Println(parsedURL)

	query := parsedURL.Query()
	query.Set("refresh-hash", hash)
	parsedURL.RawQuery = query.Encode()

	return parsedURL.String(), nil
}

func generateRandomHash() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}