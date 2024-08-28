package webapp

import (
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

	authHttp "github.com/uvulpos/golang-sveltekit-template/src/resources/auth/http"
	authService "github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service"
	authStorage "github.com/uvulpos/golang-sveltekit-template/src/resources/auth/storage"

	generalHttp "github.com/uvulpos/golang-sveltekit-template/src/resources/general/http"
	generalService "github.com/uvulpos/golang-sveltekit-template/src/resources/general/service"
	generalStorage "github.com/uvulpos/golang-sveltekit-template/src/resources/general/storage"

	userHttp "github.com/uvulpos/golang-sveltekit-template/src/resources/user/http"
	userService "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service"
	userStorage "github.com/uvulpos/golang-sveltekit-template/src/resources/user/storage"

	authPackageService "github.com/uvulpos/golang-sveltekit-template/src/resources/identity-provider/auth/service"
	authPackageStorage "github.com/uvulpos/golang-sveltekit-template/src/resources/identity-provider/auth/storage"

	jwtPackageService "github.com/uvulpos/golang-sveltekit-template/src/resources/identity-provider/jwt/service"
	jwtPackageStorage "github.com/uvulpos/golang-sveltekit-template/src/resources/identity-provider/jwt/storage"
)

type App struct {
	AuthHandler    AuthHandler
	GeneralHandler GeneralHandler
	UserHandler    UserHandler
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

	jwtPackageStorage := jwtPackageStorage.NewJwtStorage(dbConn)
	jwtPackageSvc := jwtPackageService.NewJwtService(jwtPackageStorage, "somethingNice")

	authPackageStore := authPackageStorage.NewAuthStorage(dbConn)
	authPackageSvc := authPackageService.NewAuthService(
		authPackageStore,
		configuration.AUTHORIZATION_OAUTH_KEY,
		configuration.AUTHORIZATION_OAUTH_SECRET,
		configuration.AUTHORIZATION_OAUTH_CALLBACK_URL,
		configuration.AUTHORIZATION_OAUTH_AUTHORIZATION_URL,
		configuration.AUTHORIZATION_OAUTH_TOKEN_URL,
		configuration.AUTHORIZATION_OAUTH_USERINFO_URL,
		configuration.AUTHORIZATION_OAUTH_LOGOUT_URL,
		configuration.AUTHORIZATION_OAUTH_SCOPES...,
	)

	authStore := authStorage.NewAuthStore(dbConn)
	authService := authService.NewAuthService(authStore, authPackageSvc, jwtPackageSvc)
	authHandler := authHttp.NewAuthHandler(authService)

	generalStore := generalStorage.NewGeneralStore(dbConn)
	generalSvc := generalService.NewGeneralSvc(generalStore)
	generalHandler := generalHttp.NewGeneralHandler(generalSvc)

	userStore := userStorage.NewUserStore(dbConn)
	userSvc := userService.NewUserService(userStore)
	userHandler := userHttp.NewUserHandler(userSvc)

	return &App{
		AuthHandler:    authHandler,
		GeneralHandler: generalHandler,
		UserHandler:    userHandler,
	}
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

	// tlsCert, err := tls.X509KeyPair(certificatePEM, privateKeyPEM)
	// if err != nil {
	// 	panic("Fehler beim Erstellen des tls.Certificate:" + err.Error())
	// }

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
