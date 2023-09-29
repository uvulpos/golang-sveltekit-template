package cookies

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateAuthenticationFiberCookie(jwt string) *fiber.Cookie {
	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = jwt
	cookie.HTTPOnly = true
	cookie.Secure = true
	cookie.SameSite = "strict"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	return cookie
}
