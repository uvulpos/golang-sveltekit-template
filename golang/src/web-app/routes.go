package webapp

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *App) createRoutes(router *fiber.App) {
	api := router.Group("/api")

	api.All("/test", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "api is working",
		})
	})

	api.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"message": "status not found",
		})
	})

}
