package cookies

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
)

const CookieName_RefreshToken = "refresh_token"

func GenerateRefreshToken(value string, delete bool) *fiber.Cookie {

	maxAge := 0
	expires := time.Now().Add(time.Minute * time.Duration(configuration.REFRESH_TOKEN_VALIDITY_IN_DAYS))

	if delete {
		value = ""
		expires = time.Now().Add(-(time.Hour * 100))
		maxAge = -1
	}

	// set refreshToken
	return &fiber.Cookie{
		Name:     CookieName_RefreshToken,
		Path:     configuration.AUTHCOOKIES_PATH_API_V1,
		Value:    value,
		HTTPOnly: true,
		Expires:  expires,
		MaxAge:   maxAge,
		// Secure:   true,
		SameSite: "Strict",
	}
}
