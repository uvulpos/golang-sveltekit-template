package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	jwtToken := c.Cookies("jwt", "")

	// after jwt got evaluated, add information from jwt to fiber request context
	c.Locals("jwt-token", jwtToken)

	return c.Next()
}
