package webapp

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	HandleLogin(c *fiber.Ctx) error
	HandleLogout(c *fiber.Ctx) error
	HandleChangePassword(c *fiber.Ctx) error
	HandleGetProfile(c *fiber.Ctx) error
	HandleGetUser(c *fiber.Ctx) error
	HandleGetUsers(c *fiber.Ctx) error
	HandleCheckUsername(c *fiber.Ctx) error
	HandleCheckEmail(c *fiber.Ctx) error
	HandleUpdateUserData(c *fiber.Ctx) error
	HandleJWTRefresh(c *fiber.Ctx) error
}
