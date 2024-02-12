package webapp

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/spf13/viper"

	"github.com/gofiber/swagger"
	_ "github.com/uvulpos/go-svelte/swagger-docs"

	"github.com/uvulpos/go-svelte/src/assets"
	"github.com/uvulpos/go-svelte/src/helper/config"
	dbHelper "github.com/uvulpos/go-svelte/src/helper/database"
	"github.com/uvulpos/go-svelte/src/helper/webauthn"

	userHttp "github.com/uvulpos/go-svelte/src/resources/users/http"
	userService "github.com/uvulpos/go-svelte/src/resources/users/service"
	userStorage "github.com/uvulpos/go-svelte/src/resources/users/storage"

	passkeyHttp "github.com/uvulpos/go-svelte/src/resources/passkeys-fido/http"
	passkeyService "github.com/uvulpos/go-svelte/src/resources/passkeys-fido/service"
	passkeyStorage "github.com/uvulpos/go-svelte/src/resources/passkeys-fido/storage"
)

type App struct {
	UserHandler    UserHandler
	PasskeyHandler PasskeyHandler
}

func NewApp(configuration *config.Configuration) *App {

	dbConn := dbHelper.CreateDatabase(configuration)

	webAuthNHandler := webauthn.CreateNewWebAuthN(
		"Go Svelte Binary Localhost",
		"web.localhost",
		"http://web.localhost/",
	)

	userStore := userStorage.NewUserStore(dbConn)
	userSvc := userService.NewUserSvc(userStore)
	userHandler := userHttp.NewUserHandler(userSvc)

	passkeyStore := passkeyStorage.NewUserStore(dbConn)
	passkeySvc := passkeyService.NewPasskeySvc(passkeyStore, userSvc, webAuthNHandler)
	passkeyHandler := passkeyHttp.NewPasskeyHandler(passkeySvc)

	return &App{
		UserHandler:    userHandler,
		PasskeyHandler: passkeyHandler,
	}
}

func (a *App) RunApp(configuration *config.Configuration) {

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

	if viper.GetBool("show-swagger") {
		router.Get("/swagger/*", swagger.HandlerDefault)
	}

	if !configuration.Webserver.NoFrontend {
		router.Use("/", filesystem.New(filesystem.Config{
			Root:         http.FS(publicFS),
			NotFoundFile: "index.html",
		}))
	}

	router.Use(Handle404)

	serverPort := fmt.Sprintf(":%d", configuration.Webserver.Port)
	log.Printf("server listens on %s\n", serverPort)
	router.Listen(serverPort)
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
