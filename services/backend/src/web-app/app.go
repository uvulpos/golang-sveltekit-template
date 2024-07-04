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
	swaggercss "github.com/uvulpos/go-svelte/src/assets/swagger-css"
	"github.com/uvulpos/go-svelte/src/configuration"
	dbHelper "github.com/uvulpos/go-svelte/src/helper/database"

	generalHttp "github.com/uvulpos/go-svelte/src/resources/general/http"
	generalService "github.com/uvulpos/go-svelte/src/resources/general/service"
	generalStorage "github.com/uvulpos/go-svelte/src/resources/general/storage"
)

type App struct {
	GeneralHandler GeneralHandler
}

func NewApp() *App {

	dbConn, dbConnErr := dbHelper.CreateDatabase()
	if dbConn == nil || dbConn.DB == nil || dbConnErr != nil {
		err := fmt.Errorf("could not connect to database")
		if err != nil {
			panic(err)
		}
		return nil
	}

	generalStore := generalStorage.NewGeneralStore(*dbConn)
	generalSvc := generalService.NewGeneralSvc(generalStore)
	generalHandler := generalHttp.NewGeneralHandler(generalSvc)

	return &App{
		GeneralHandler: generalHandler,
	}
}

//	@title			Golang + SvelteKit API
//	@version		1.0
//	@description	This is a sample swagger for this template
//
//	@contact.name	GitHub Template
//	@contact.url	https://www.github.com/uvulpos
//
//	@host			web.localhost
//	@BasePath		/
func (a *App) RunApp() {
	publicFS, err := fs.Sub(assets.SvelteFS, "frontend")
	if err != nil {
		log.Fatal(err)
	}

	router := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// for internal use you can print error to response
			// "Unexpected Error Occured " + "\n" + err.Error()
			return c.Status(fiber.StatusInternalServerError).SendString("Unexpected Error Occured")
		},
	})

	a.createRoutes(router)

	if configuration.WEBSERVER_SHOW_SWAGGER {
		router.Get("/swagger/*", swagger.New(swagger.Config{
			CustomStyle: swaggercss.DefaultCSS,
		}))
	}

	if configuration.WEBSERVER_SHOW_FRONTEND {
		router.Use("/", filesystem.New(filesystem.Config{
			Root:         http.FS(publicFS),
			NotFoundFile: "index.html",
		}))
	}

	router.Use(Handle404)

	serverPort := fmt.Sprintf(":%d", configuration.WEBSERVER_PORT)
	err = router.Listen(serverPort)
	if err != nil {
		panic(err)
	}
}

func (a *App) ReturnAppInE2EMode() *fiber.App {

	router := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler: func(c *fiber.Ctx, err error) error {

			// for internal use you can print error to response
			// "Unexpected Error Occured " + "\n" + err.Error()
			return c.Status(fiber.StatusInternalServerError).SendString("Unexpected Error Occured")
		},
	})

	a.createRoutes(router)
	router.Use(Handle404)

	return router
}
