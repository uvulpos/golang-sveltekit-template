package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	customerrorconst "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/custom-error-const"
)

func (h *MiddlewareHandler) AuthenticationRefreshToken() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		refreshToken := c.Cookies("refresh_token", "")

		authorizationHeader := c.Get("Authorization")

		if refreshToken == "" && len(authorizationHeader) > 7 && strings.HasPrefix(authorizationHeader, "Bearer ") {
			refreshToken = authorizationHeader[7:]
		} else if refreshToken == "" {
			return c.Status(http.StatusUnauthorized).SendString(customerrorconst.NOT_AUTHORIZED_ERROR_MESSAGE)
		}

		// 🚨 NOT THE RIGHT IMPLEMENTATION, USE INJECTION
		jwtData, jwtDataErr := h.jwtSvc.VerifyRefreshToken(refreshToken)
		if jwtDataErr != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
		}

		fmt.Println("jwtData", jwtData)

		// after jwt got evaluated, add information from jwt to fiber request context
		c.Locals("session-uuid", jwtData.SessionID)

		return c.Next()
	}
}
