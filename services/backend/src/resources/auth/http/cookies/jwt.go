package cookies

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func GenerateJwtToken(value string, delete bool, expires time.Time) *fiber.Cookie {
	maxAge := 0

	if delete {
		value = ""
		expires = time.Now()
		maxAge = -1
	}

	return &fiber.Cookie{
		Name:    "jwt",
		Path:    "/",
		Value:   value,
		Expires: expires,
		MaxAge:  maxAge,
	}
}
