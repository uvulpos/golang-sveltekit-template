package cookies

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func GenerateRefreshToken(value string, delete bool, expires time.Time) *fiber.Cookie {

	maxAge := 0

	if delete {
		value = ""
		expires = time.Now()
		maxAge = -1
	}

	// set refreshToken
	return &fiber.Cookie{
		Name:     "refresh_token",
		Path:     "/api/v1/auth",
		Value:    value,
		HTTPOnly: true,
		Expires:  expires,
		MaxAge:   maxAge,
		// Secure:   true,
		SameSite: "Strict",
	}
}
