package http

import (
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) HandleGetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user-uuid").(string)
	h.service.GetProfileByUserUUID(userID)
	return nil
}
