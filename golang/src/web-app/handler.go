package webapp

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	HandleGetProfile(c *fiber.Ctx) error
}
