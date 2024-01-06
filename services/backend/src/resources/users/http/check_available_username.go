package http

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CheckUsernamePayload struct {
	Username string `json:"username" xml:"username" form:"username"`
}

type CheckUsernameResponseBody struct {
	Status bool `json:"status" xml:"status" form:"status"`
}

func (h *UserHandler) HandleCheckUsername(c *fiber.Ctx) error {
	requestUser := c.Locals("user-uuid")
	userUuid, ok := requestUser.(string)
	if !ok {
		return c.SendStatus(http.StatusUnauthorized)
	}

	payload := CheckUsernamePayload{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	isAvailable, isAvailableErr := h.service.IsAvailableUsername(nil, payload.Username, userUuid)
	if isAvailableErr != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	response := CheckUsernameResponseBody{Status: isAvailable}
	responseJson, responseErr := json.Marshal(response)
	if responseErr != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).Send(responseJson)
}
