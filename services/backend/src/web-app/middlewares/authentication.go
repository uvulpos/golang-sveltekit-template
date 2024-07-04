package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	// jwtToken := c.Cookies("jwt", "")
	// errormessage := "could not validate authentication"

	// if jwtToken == "" {
	// 	_ = c.Status(http.StatusUnauthorized).SendString(errormessage)
	// }

	// c.Locals("user-uuid", jwtData.Uuid)
	// c.Locals("username", jwtData.Username)
	// c.Locals("permissions", jwtData.Permissions)

	return c.Next()
}
