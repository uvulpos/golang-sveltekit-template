package middlewares

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	customerrorconst "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/custom-error-const"
	"github.com/uvulpos/golang-sveltekit-template/src/resources/auth/http/cookies"
	sessionLocals "github.com/uvulpos/golang-sveltekit-template/src/web-app/middlewares/http/consts/session-locals"
)

func (h *MiddlewareHandler) AuthenticationRefreshToken() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		refreshToken := c.Cookies(cookies.CookieName_RefreshToken, "")

		authorizationHeader := c.Get("Authorization")

		if refreshToken == "" && len(authorizationHeader) > 7 && strings.HasPrefix(authorizationHeader, "Bearer ") {
			refreshToken = authorizationHeader[7:]
		} else if refreshToken == "" {
			return c.Status(http.StatusUnauthorized).SendString(customerrorconst.NOT_AUTHORIZED_ERROR_MESSAGE)
		}

		jwtData, jwtDataErr := h.jwtSvc.VerifyRefreshToken(refreshToken)
		if jwtDataErr != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
		}

		// after jwt got evaluated, add information from jwt to fiber request context
		c.Locals(sessionLocals.SessionUUID, jwtData.SessionID)

		return c.Next()
	}
}
