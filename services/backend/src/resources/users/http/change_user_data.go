package http

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/go-svelte/src/resources/users/http/models"
)

func (h *UserHandler) HandleUpdateUserData(c *fiber.Ctx) error {
	requestUser := c.Locals("user-uuid")
	userUuid, ok := requestUser.(string)
	if !ok {
		return c.SendStatus(http.StatusUnauthorized)
	}

	payload := models.ChangeUserDataPayload{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	newUser, isUpdatedUser := h.service.UpdateUserData(payload, userUuid)
	if isUpdatedUser != nil {
		return c.Status(http.StatusInternalServerError).SendString(isUpdatedUser.Error())
	}

	httpUser := ToSelfUserWithPermission(newUser)

	responeJson, responeJsonErr := json.Marshal(httpUser)
	if responeJsonErr != nil {
		return c.Status(http.StatusInternalServerError).SendString(responeJsonErr.Error())
	}

	return c.Status(http.StatusOK).Send(responeJson)
}
