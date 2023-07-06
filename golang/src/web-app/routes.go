package webapp

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *App) createRoutes(router *fiber.App) {
	api := router.Group("/api/v1")

	api.All("/profile/getProfile", func(c *fiber.Ctx) error {
		var permissions []string = []string{"read_posts", "write_posts"}
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"username":    "uVulpos",
			"permissions": permissions,
		})
	})

	api.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"message": "status not found",
		})

	})

}
