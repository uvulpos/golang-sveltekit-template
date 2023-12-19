package middlewares

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/go-svelte/src/helper/jwt"
)

func Authentication(c *fiber.Ctx) error {
	jwtToken := c.Cookies("jwt", "")
	errormessage := "could not validate authentication"

	if jwtToken == "" {
		c.Status(http.StatusUnauthorized).SendString(errormessage)
	}

	isJwtValid, jwtData, isJwtValidErr := jwt.VerifyJWToken(jwtToken)
	if isJwtValidErr != nil || !isJwtValid {
		c.Status(http.StatusUnauthorized).SendString(errormessage)
	}

	c.Locals("user-uuid", jwtData.Uuid)
	c.Locals("username", jwtData.Username)
	c.Locals("permissions", jwtData.Permissions)

	return c.Next()
}
