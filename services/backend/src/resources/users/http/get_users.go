package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	var requiredPermissions []string = []string{"MANAGE_USERS"}

	// get uuid from jwt
	requestUser := c.Locals("user-uuid")
	userUuid, ok := requestUser.(string)
	if !ok {
		return c.SendStatus(http.StatusUnauthorized)
	}

	// todo: read offset from request + filter

	user, userErr := h.service.GetUserByUUID(nil, userUuid)
	if userErr != nil {
		return c.SendStatus(http.StatusForbidden)
	}

	for _, requiredPermission := range requiredPermissions {
		manageUsersPermissions := user.Permissions.IncludePermission(requiredPermission)
		if !manageUsersPermissions {
			fmt.Println("PERMISSION NOT FOUND! needed:", requiredPermission, " have: ")
			fmt.Printf("%v\n", user.Permissions)
			return c.SendStatus(http.StatusForbidden)
		}
	}

	users, usersErr := h.service.GetUsers()
	if usersErr != nil {
		// fmt.Println(usersErr)
		return c.SendStatus(http.StatusInternalServerError)
	}

	httpUsers := ToUsersForManager(users)

	responeJson, responeJsonErr := json.Marshal(httpUsers)
	if responeJsonErr != nil {
		return c.Status(http.StatusInternalServerError).SendString(responeJsonErr.Error())
	}

	return c.Status(http.StatusOK).Send(responeJson)
}
