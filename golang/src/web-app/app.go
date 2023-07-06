package webapp

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/uvulpos/go-svelte/src/assets"

	dbHelper "github.com/uvulpos/go-svelte/src/helper/database"

	userHttp "github.com/uvulpos/go-svelte/src/resources/users/http"
	userService "github.com/uvulpos/go-svelte/src/resources/users/service"
	userStorage "github.com/uvulpos/go-svelte/src/resources/users/storage"
)

type App struct {
	UserHandler UserHandler
}

func NewApp() *App {

	dbConn := dbHelper.CreateDatabase()

	userStore := userStorage.NewUserStore(dbConn)
	userSvc := userService.NewUserSvc(userStore)
	userHandler := userHttp.NewUserHandler(userSvc)

	return &App{
		UserHandler: userHandler,
	}
}

func (a *App) RunApp(showFrontend bool, webserverPort int) {

	publicFS, err := fs.Sub(assets.SvelteFS, "frontend")
	if err != nil {
		log.Fatal(err)
	}

	router := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	a.createRoutes(router)

	if showFrontend {
		router.Use("/", filesystem.New(filesystem.Config{
			Root:         http.FS(publicFS),
			NotFoundFile: "index.html",
		}))
	}

	serverPort := fmt.Sprintf(":%d", webserverPort)
	log.Printf("server listens on %s\n", serverPort)
	router.Listen(serverPort)
}
