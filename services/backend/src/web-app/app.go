package webapp

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"

	"github.com/gofiber/swagger"
	_ "github.com/uvulpos/go-svelte/swagger-docs"

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

// @title			Fiber Example API
// @version		1.0
// @description	This is a sample swagger for Fiber
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	fiber@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/
func (a *App) RunApp(showFrontend, showSwagger bool, webserverPort int) {

	publicFS, err := fs.Sub(assets.SvelteFS, "frontend")
	if err != nil {
		log.Fatal(err)
	}

	router := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	a.createRoutes(router)

	if showSwagger {
		router.Get("/swagger/*", swagger.HandlerDefault)
	}

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
