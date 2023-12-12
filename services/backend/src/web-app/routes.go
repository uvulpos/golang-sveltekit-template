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

	api.Post("/login", a.UserHandler.HandleLogin)
	api.Post("/logout", a.UserHandler.HandleLogout)

	api.Use(Handle404)

	apiV1 := api.Group("/v1")
	apiV1.Use(Handle404)

}

func Handle404(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"code":    http.StatusNotFound,
		"message": "Not found ❌",
	})
}
