package cookies

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
)

const CookieName_JwtToken = "jwt"

func GenerateJwtToken(value string, delete bool) *fiber.Cookie {
	maxAge := 0
	expires := time.Now().Add(time.Minute * time.Duration(configuration.JWT_TOKEN_VALIDITY_IN_MINUTES))

	if delete {
		value = ""
		expires = time.Now().Add(-(time.Hour * 100))
		maxAge = -1
	}

	return &fiber.Cookie{
		Name:    CookieName_JwtToken,
		Path:    "/",
		Value:   value,
		Expires: expires,
		MaxAge:  maxAge,
	}
}
