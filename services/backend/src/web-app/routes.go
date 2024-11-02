package webapp

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *App) CreateRoutes(router *fiber.App) {

	api := router.Group("/api")

	api.All("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	api.Get("/ping", a.GeneralHandler.Ping)
	api.Get("/system-healthcheck", a.GeneralHandler.SystemHealthCheck)
	api.Get("/asset/gopher-coffee", a.GeneralHandler.AssetGopherCoffee)
	api.Get("/asset/logo", a.GeneralHandler.AssetLogo)
	api.Get("/asset/logo-branding", a.GeneralHandler.AssetLogoBranding)

	apiV1 := api.Group("v1")

	apiV1.Get("/oauth/redirect", a.AuthHandler.CreateRedirect)
	apiV1.Get("/oauth/callback", a.AuthHandler.CallbackHandler)
	// apiV1.Get("/oauth/logout", a.AuthHandler.Logout) // deprecated
	apiV1.Get("/auth/logout", a.AuthHandler.Logout)
	apiV1.Get("/auth/refresh", a.MiddlewareHandler.AuthenticationRefreshToken(), a.AuthHandler.Refresh)

	// apiV1Admin := apiV1.Group("/admin-shit/", a.MiddlewareHandler.Authentication([]string{"admin:read", "admin:write"}))
	// apiV1Admin.Get("/", TestEP)

	apiV1.Use(a.MiddlewareHandler.Authentication([]string{}))

	apiV1.Get("/test", TestEP)
	apiV1.Get("/self", a.UserHandler.GetSelfInformation)
	apiV1.Get("/self/permissions", a.UserHandler.GetSelfPermissions)

	// apiV1Admin.Use(Handle404)
	apiV1.Use(Handle404)
	api.Use(Handle404)
}

func Handle404(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"code":    http.StatusNotFound,
		"message": "Not found ❌",
	})
}

func TestEP(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Hello World",
	})
}
