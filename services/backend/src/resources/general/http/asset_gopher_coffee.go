package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/go-svelte/src/assets"
)

func (h *GeneralHandler) AssetGopherCoffee(c *fiber.Ctx) error {
	c.Set("Content-Type", "image/gif")
	return c.SendString(assets.GopherCoffee)
}
