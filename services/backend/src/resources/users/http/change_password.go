package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ChangePasswordPayload struct {
	OldPassword string `json:"oldPassword" xml:"oldPassword" form:"oldPassword"`
	NewPassword string `json:"newPassword" xml:"newPassword" form:"newPassword"`
}

func (h *UserHandler) HandleChangePassword(c *fiber.Ctx) error {
	requestUser := c.Locals("username")
	username, ok := requestUser.(string)
	if !ok {
		return c.SendStatus(http.StatusUnauthorized)
	}

	payload := ChangePasswordPayload{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	authErr := h.service.UpdatePassword(nil, username, payload.NewPassword, payload.OldPassword)
	if authErr != nil {
		return c.Status(http.StatusBadRequest).SendString(authErr.Error())
	}

	return c.Status(http.StatusOK).SendString("password changed")
}
