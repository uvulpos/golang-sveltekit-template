package http

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	"github.com/uvulpos/golang-sveltekit-template/src/resources/auth/http/cookies"
)

func (h *AuthHandler) CallbackHandler(c *fiber.Ctx) error {
	authCode := c.Query("code", "")
	state := c.Query("state", "")
	authProvider := c.Cookies(cookies.CookieName_AuthProvider, "")

	fmt.Println()
	fmt.Println("authProvider", authProvider)
	fmt.Println()

	jwtToken, refreshToken, err := h.service.CallbackFunction(
		authProvider,
		authCode,
		state,
		/* ---------------------------------------------
				GDPR sensible information
		--------------------------------------------- */
		"", // c.IP(),              // 	for GDPR reasons disabled be default
		"", // c.Get("User-Agent"), // 	for GDPR reasons disabled be default
	)

	if err != nil || jwtToken == "" {
		return err
	}

	c.Cookie(cookies.GenerateJwtToken(jwtToken, false))
	c.Cookie(cookies.GenerateRefreshToken(refreshToken, false))

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
		return "", customerrors.NewInternalServerError(err, "", fmt.Sprintf("Couldn't parse URL in oauth callback (func: generateRandomHashURL ; data: %v)", baseURL))
	}

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
