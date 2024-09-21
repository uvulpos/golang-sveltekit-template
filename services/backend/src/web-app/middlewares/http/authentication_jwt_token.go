package middlewares

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	customerrorconst "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/custom-error-const"
)

func (h *MiddlewareHandler) Authentication(requiredPermissions []string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		jwtToken := c.Cookies("jwt", "")

		authorizationHeader := c.Get("Authorization")

		if jwtToken == "" && len(authorizationHeader) > 7 && strings.HasPrefix(authorizationHeader, "Bearer ") {
			jwtToken = authorizationHeader[7:]
		} else if jwtToken == "" {
			return c.Status(http.StatusUnauthorized).SendString(customerrorconst.NOT_AUTHORIZED_ERROR_MESSAGE)
		}

		// 🚨 NOT THE RIGHT IMPLEMENTATION, USE INJECTION
		jwtData, jwtDataErr := h.jwtSvc.VerifyJWT(jwtToken)
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
		c.Locals("user-uuid", jwtData.UserUuid)
		c.Locals("permissions", permissions)

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
