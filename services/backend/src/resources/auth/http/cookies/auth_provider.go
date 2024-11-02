package cookies

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
)

const CookieName_AuthProvider = "auth_provder"

func GenerateAuthProviderCookie(value string, delete bool) *fiber.Cookie {
	maxAge := 0
	expires := time.Now().Add(time.Minute * 10)

	if delete {
		value = ""
		expires = time.Now().Add(-(time.Hour * 100))
		maxAge = -1
	}

	return &fiber.Cookie{
		Name:    CookieName_AuthProvider,
		Path:    configuration.AUTHCOOKIES_PATH_API_V1,
		Value:   value,
		Expires: expires,
		MaxAge:  maxAge,
	}
}
