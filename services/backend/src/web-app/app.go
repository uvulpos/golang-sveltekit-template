package webapp

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"

	"github.com/gofiber/swagger"
	_ "github.com/uvulpos/golang-sveltekit-template/swagger-docs"

	"github.com/uvulpos/golang-sveltekit-template/src/assets"
	swaggercss "github.com/uvulpos/golang-sveltekit-template/src/assets/swagger-css"
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
	dbHelper "github.com/uvulpos/golang-sveltekit-template/src/helper/database"

	customhttphandler "github.com/uvulpos/golang-sveltekit-template/src/web-app/custom-http-handler"
	middlewareHttp "github.com/uvulpos/golang-sveltekit-template/src/web-app/middlewares/http"
	middlewareService "github.com/uvulpos/golang-sveltekit-template/src/web-app/middlewares/service"

	authHttp "github.com/uvulpos/golang-sveltekit-template/src/resources/auth/http"
	authService "github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service"
	authStorage "github.com/uvulpos/golang-sveltekit-template/src/resources/auth/storage"

	generalHttp "github.com/uvulpos/golang-sveltekit-template/src/resources/general/http"
	generalService "github.com/uvulpos/golang-sveltekit-template/src/resources/general/service"
	generalStorage "github.com/uvulpos/golang-sveltekit-template/src/resources/general/storage"

	userHttp "github.com/uvulpos/golang-sveltekit-template/src/resources/user/http"
	userService "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service"
	userStorage "github.com/uvulpos/golang-sveltekit-template/src/resources/user/storage"
)

type App struct {
	MiddlewareHandler MiddlewareHandler
	AuthHandler       AuthHandler
	GeneralHandler    GeneralHandler
	UserHandler       UserHandler
}

func NewApp() *App {

	services, servicesErr := SetupServices()
	if servicesErr != nil {
		panic(servicesErr)
	}

	return &App{
		MiddlewareHandler: services.MiddlewareHandler,

		AuthHandler:    services.AuthHandler,
		GeneralHandler: services.GeneralHandler,
		UserHandler:    services.UserHandler,
	}
}

type AppServices struct {
	AuthHandler *authHttp.AuthHandler
	AuthService *authService.AuthService
	AuthStore   *authStorage.AuthStore

	GeneralHandler *generalHttp.GeneralHandler
	GeneralSvc     *generalService.GeneralSvc
	GeneralStore   *generalStorage.GeneralStore

	MiddlewareHandler *middlewareHttp.MiddlewareHandler
	MiddlewareSvc     *middlewareService.MiddlewareService

	UserHandler *userHttp.UserHandler
	UserSvc     *userService.UserService
	UserStore   *userStorage.UserStore
}

/*
*
*	Why is there a setup services function?
*
*	Because for the integration tests, they shouldn't access the service functions
*	but the reality is, some functions are needed, like generating a jwt session for the request
*	so use this opportunity as a cautious gift, but not as a priviledge
*
 */
func SetupServices() (*AppServices, error) {
	dbConn, dbConnErr := dbHelper.CreateDatabase()
	if dbConn == nil || dbConn.DB == nil || dbConnErr != nil {
		err := fmt.Errorf("could not connect to database")
		if err != nil {
			return nil, err
		}
		return nil, errors.New("cannot create the database for unknown reasons")
	}

	generalStore := generalStorage.NewGeneralStore(dbConn)
	generalSvc := generalService.NewGeneralSvc(generalStore)
	generalHandler := generalHttp.NewGeneralHandler(generalSvc)

	userStore := userStorage.NewUserStore(dbConn)
	userSvc := userService.NewUserService(userStore)
	userHandler := userHttp.NewUserHandler(userSvc)

	authStore := authStorage.NewAuthStore(dbConn)
	authService := authService.NewAuthService(authStore, userSvc)
	authHandler := authHttp.NewAuthHandler(authService)

	middlewareSvc := middlewareService.NewMiddlewareService(userSvc)
	middlewareHandler := middlewareHttp.NewMiddlewareHandler(middlewareSvc)

	return &AppServices{
		AuthHandler: authHandler,
		AuthService: authService,
		AuthStore:   authStore,

		GeneralHandler: generalHandler,
		GeneralSvc:     generalSvc,
		GeneralStore:   generalStore,

		MiddlewareHandler: middlewareHandler,
		MiddlewareSvc:     middlewareSvc,

		UserHandler: userHandler,
		UserSvc:     userSvc,
		UserStore:   userStore,
	}, nil
}

// @title		Golang + SvelteKit API
// @version	0.1
// @description
// @description	<img alt="coffee drinking gopher" src="/api/asset/gopher-coffee" height="200px">
// @description
// @description	This is a sample swagger for this template
// @description
// @description	[Return back to application](/) // [View on GitHub](https://github.com/uvulpos/golang-sveltekit-binary)
//
// @host			web.localhost
// @BasePath		/
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

	a.CreateRoutes(router)

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
	// err = router.ListenTLSWithCertificate(serverPort, tlsCert)
	if err != nil {
		panic(err)
	}
}

func (a *App) ReturnAppInE2EMode() *fiber.App {

	router := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          customhttphandler.UnexpectedErrorHandler,
	})

	a.CreateRoutes(router)
	router.Use(Handle404)

	return router
}
