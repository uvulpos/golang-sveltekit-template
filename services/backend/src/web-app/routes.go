package webapp

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *App) createRoutes(router *fiber.App) {

	api := router.Group("/api")

	api.All("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	api.Get("/ping", a.GeneralHandler.Ping)
	api.Get("/ping-db", a.GeneralHandler.PingDatabase)

	apiV1 := api.Group("v1")

	apiV1.Use(Handle404)
	api.Use(Handle404)
}

func Handle404(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"code":    http.StatusNotFound,
		"message": "Not found ❌",
	})
}
