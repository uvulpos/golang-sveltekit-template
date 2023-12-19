package http

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) HandleGetProfile(c *fiber.Ctx) error {
	// get uuid from jwt
	requestUser := c.Locals("user-uuid")
	userUuid, ok := requestUser.(string)
	if !ok {
		return c.SendStatus(http.StatusUnauthorized)
	}

	user, userErr := h.service.GetUserByUUID(userUuid)
	if userErr != nil {
		return c.Status(http.StatusUnauthorized).SendString(userErr.Error())
	}

	httpUser := ToUserWithPermission(user)

	responeJson, responeJsonErr := json.Marshal(httpUser)
	if responeJsonErr != nil {
		return c.Status(http.StatusInternalServerError).SendString(responeJsonErr.Error())
	}

	return c.Status(http.StatusOK).Send(responeJson)
}
