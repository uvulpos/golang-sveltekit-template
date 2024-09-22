package webapp

import "github.com/gofiber/fiber/v2"

type MiddlewareHandler interface {
	Authentication(requiredPermissions []string) func(c *fiber.Ctx) error
	AuthenticationRefreshToken() func(c *fiber.Ctx) error
}

type AuthHandler interface {
	CreateRedirect(c *fiber.Ctx) error
	CallbackHandler(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx) error
}

type GeneralHandler interface {
	Ping(c *fiber.Ctx) error
	SystemHealthCheck(c *fiber.Ctx) error
	AssetGopherCoffee(c *fiber.Ctx) error
	AssetLogo(c *fiber.Ctx) error
	AssetLogoBranding(c *fiber.Ctx) error
}

type UserHandler interface {
	GetSelfPermissions(c *fiber.Ctx) error
	GetSelfInformation(c *fiber.Ctx) error
}
