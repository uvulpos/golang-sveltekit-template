package webapp

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	CreateRedirect(c *fiber.Ctx) error
	CallbackHandler(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type GeneralHandler interface {
	Ping(c *fiber.Ctx) error
	SystemHealthCheck(c *fiber.Ctx) error
	AssetGopherCoffee(c *fiber.Ctx) error
}
