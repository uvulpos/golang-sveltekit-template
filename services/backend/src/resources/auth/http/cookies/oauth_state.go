package cookies

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

const CookieName_OauthState = "auth_provder"

func GenerateOauthStateCookie(value string, delete bool) *fiber.Cookie {
	maxAge := 0
	expires := time.Now().Add(time.Minute * 15)

	if delete {
		value = ""
		expires = time.Now().Add(-(time.Hour * 100))
		maxAge = -1
	}

	return &fiber.Cookie{
		Name:    CookieName_AuthProvider,
		Value:   value,
		Path:    "/api/v1/oauth",
		Expires: expires,
		MaxAge:  maxAge,
	}
}
