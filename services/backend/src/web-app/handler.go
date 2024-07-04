package webapp

import "github.com/gofiber/fiber/v2"

type GeneralHandler interface {
	Ping(c *fiber.Ctx) error
	PingDatabase(c *fiber.Ctx) error
}
