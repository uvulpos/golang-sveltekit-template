package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) HandleGetProfile(c *fiber.Ctx) error {

	var permissions []string = []string{"read_posts", "write_posts"}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"username":    "uVulpos",
		"permissions": permissions,
	})

	// userID := c.Locals("user-uuid").(string)
	// h.service.GetProfileByUserUUID(userID)
	// return nil
}
