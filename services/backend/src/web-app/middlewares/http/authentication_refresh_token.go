package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/golang-sveltekit-template/src/resources/auth/http/cookies"

	customerrorconst "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/custom-error-const"
	jwt "github.com/uvulpos/golang-sveltekit-template/src/helper/jwt"
	sessionLocals "github.com/uvulpos/golang-sveltekit-template/src/web-app/middlewares/http/consts/session-locals"
)

func (h *MiddlewareHandler) AuthenticationRefreshToken() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		refreshToken := c.Cookies(cookies.CookieName_RefreshToken, "")

		jwtData, jwtDataErr := jwt.VerifyRefreshToken(refreshToken)
		if jwtDataErr != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
		}

		// after jwt got evaluated, add information from jwt to fiber request context
		c.Locals(sessionLocals.SessionUUID, jwtData.SessionID)

		return c.Next()
	}
}
