package webapp

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	HandleLogin(c *fiber.Ctx) error
	HandleLogout(c *fiber.Ctx) error
}
