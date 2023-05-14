package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed frontend/*
var svelteFS embed.FS

func main() {
	app := fiber.New()

	publicFS, err := fs.Sub(svelteFS, "frontend")
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(publicFS),
		NotFoundFile: "index.html",
	}))

	app.Listen(":3000")
}
