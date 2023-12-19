package webapp

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/go-svelte/src/web-app/middlewares"
)

func (a *App) createRoutes(router *fiber.App) {

	api := router.Group("/api")

	api.All("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	apiV1 := api.Group("v1")
	apiV1.Use(Handle404)

	api.Post("login", a.UserHandler.HandleLogin)
	api.Post("logout", a.UserHandler.HandleLogout)

	api.Use(middlewares.Authentication)
	api.Post("login/change-password", a.UserHandler.HandleChangePassword)
	api.Get("self/getProfile", a.UserHandler.HandleGetProfile)

	api.Use(Handle404)
}

func Handle404(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"code":    http.StatusNotFound,
		"message": "Not found ❌",
	})
}
