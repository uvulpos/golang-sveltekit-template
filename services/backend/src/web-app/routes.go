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

	apiV1.Post("login", a.UserHandler.HandleLogin)
	apiV1.Post("logout", a.UserHandler.HandleLogout)

	apiV1.Use(middlewares.Authentication)
	apiV1.Post("login/is-available-email", a.UserHandler.HandleCheckEmail)
	apiV1.Post("login/is-available-username", a.UserHandler.HandleCheckUsername)
	apiV1.Post("login/change-password", a.UserHandler.HandleChangePassword)
	apiV1.Get("self/get-user-data", a.UserHandler.HandleGetProfile)
	apiV1.Post("self/update-user-data", a.UserHandler.HandleUpdateUserData)

	apiV1.Use(Handle404)
	api.Use(Handle404)
}

func Handle404(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"code":    http.StatusNotFound,
		"message": "Not found ❌",
	})
}
