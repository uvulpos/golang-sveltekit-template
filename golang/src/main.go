package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

// hello world

var appversion string = "dev"
var appbuildcommit string = "_"
var appbuilddate string

//go:embed frontend/*
var svelteFS embed.FS

func main() {

	publicFS, err := fs.Sub(svelteFS, "frontend")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Version ", appversion)
	fmt.Println("Build-Commit ", appbuildcommit)
	fmt.Println("Build-Date ", appbuilddate)

	app := fiber.New(fiber.Config{
		DisableStartupMessage: appversion == "",
	})

	api := app.Group("/api")
	api.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code": http.StatusNotFound,
			"age":  "status not found",
		})
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(publicFS),
		NotFoundFile: "index.html",
	}))

	app.Listen(":3000")
}
