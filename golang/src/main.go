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

	fmt.Println("Version ", appversion)
	fmt.Println("Build-Commit ", appbuildcommit)
	fmt.Println("Build-Date ", appbuilddate)

	app := fiber.New(fiber.Config{
		DisableStartupMessage: appversion == "",
	})

	publicFS, err := fs.Sub(svelteFS, "frontend")
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/hello-world", func(f *fiber.Ctx) error {
		return f.SendString("Hello World")
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(publicFS),
		NotFoundFile: "index.html",
	}))

	app.Listen(":3000")
}

func existsID(id int, ids []int) bool {
	for _, searchingID := range ids {
		if searchingID == id {
			return true
		}
	}
	return false
}
