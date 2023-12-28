package http

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CheckEmailPayload struct {
	Email string `json:"email" xml:"email" form:"email"`
}

type CheckEmailResponseBody struct {
	Status bool `json:"status" xml:"status" form:"status"`
}

func (h *UserHandler) HandleCheckEmail(c *fiber.Ctx) error {
	requestUser := c.Locals("user-uuid")
	userUuid, ok := requestUser.(string)
	if !ok {
		return c.SendStatus(http.StatusUnauthorized)
	}

	payload := CheckEmailPayload{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	isAvailable, isAvailableErr := h.service.IsAvailableEmail(nil, payload.Email, userUuid)
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
