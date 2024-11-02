package middlewares

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/uvulpos/golang-sveltekit-template/src/resources/auth/http/cookies"

	customerrorconst "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/custom-error-const"
	jwt "github.com/uvulpos/golang-sveltekit-template/src/helper/jwt"
	sessionLocals "github.com/uvulpos/golang-sveltekit-template/src/web-app/middlewares/http/consts/session-locals"
)

func (h *MiddlewareHandler) Authentication(requiredPermissions []string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		jwtToken := c.Cookies(cookies.CookieName_JwtToken, "")

		authorizationHeader := c.Get("Authorization")

		if jwtToken == "" && len(authorizationHeader) > 7 && strings.HasPrefix(authorizationHeader, "Bearer ") {
			jwtToken = authorizationHeader[7:]
		} else if jwtToken == "" {
			return c.Status(http.StatusUnauthorized).SendString(customerrorconst.NOT_AUTHORIZED_ERROR_MESSAGE)
		}

		jwtData, jwtDataErr := jwt.VerifyJWT(jwtToken)
		if jwtDataErr != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
		}

		permissions, permissionErr := h.service.GetUserPermissionsByID(jwtData.UserUuid)
		if permissionErr != nil {
			status, _, message := permissionErr.HttpError()
			return c.Status(status).SendString(message)
		}

		if !arePermissionsSatisfied(requiredPermissions, permissions) {
			c.Status(http.StatusForbidden).SendString(customerrorconst.FORBIDDEN_ERROR_MESSAGE)
		}

		// after jwt got evaluated, add information from jwt to fiber request context
		c.Locals(sessionLocals.UserUUID, jwtData.UserUuid)
		c.Locals(sessionLocals.Permissions, permissions)

		return c.Next()
	}
}

func arePermissionsSatisfied(wantPermissions, havePermissions []string) bool {
	foundPermission := 0

	for _, wantPermission := range wantPermissions {
	HaveLoop:
		for _, havePermission := range havePermissions {
			if wantPermission == havePermission {
				foundPermission++
				break HaveLoop
			}
		}
	}

	return foundPermission == len(wantPermissions)
}
