package setup

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	jwtService "github.com/uvulpos/golang-sveltekit-template/src/resources/jwt/service"
	webApp "github.com/uvulpos/golang-sveltekit-template/src/web-app"
	customhttphandler "github.com/uvulpos/golang-sveltekit-template/src/web-app/custom-http-handler"
	"gotest.tools/assert"
)

type TestSetupStruct struct {
	FiberRouter         *fiber.App
	ApplicationStruct   *webApp.App
	ApplicationServices *webApp.AppServices
	Cookies             []*http.Cookie
}

/*
* ApplicationStruct gets exposed, in case you really need to access service layers
* If there is an endpoint to do the required action, please prefer this way
 */
func SetupTest(t *testing.T) *TestSetupStruct {
	applicationStruct := webApp.NewApp()

	fiberRouter := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          customhttphandler.UnexpectedErrorHandler,
	})

	applicationStruct.CreateRoutes(fiberRouter)

	applicationServices, applicationServicesErr := webApp.SetupServices()
	if applicationServicesErr != nil {
		panic("could not create services model")
	}

	return &TestSetupStruct{
		FiberRouter:         fiberRouter,
		ApplicationStruct:   applicationStruct,
		ApplicationServices: applicationServices,
	}
}

func (s *TestSetupStruct) AuthenticateAsUser(t *testing.T, identifier string) {
	s.ApplicationServices.UserSvc.GetUserByUsernameWithPermissions(nil, identifier)
	user, userErr := s.ApplicationServices.UserSvc.GetUserByUsernameWithPermissions(nil, identifier)
	assert.Equal(t, userErr, nil, fmt.Sprintf("could not find user for test authentication %s", identifier))

	jwtData := jwtService.NewJwtDataModel(user.ID, "", user.Permissions)
	jwtToken, jwtTokenErr := s.ApplicationServices.JwtPackageSvc.CreateJWT(jwtData)
	assert.Equal(t, jwtTokenErr, nil, fmt.Sprintf("could not create jwt token for test authentication for user %s", identifier))

	s.Cookies = append(s.Cookies, &http.Cookie{
		Name:  "jwt",
		Value: jwtToken,
	})
}

func (s *TestSetupStruct) MakeRequest(t *testing.T, applicationMethod, applicationURL string, requestBody io.Reader) *http.Response {
	req := httptest.NewRequest(applicationMethod, applicationURL, requestBody)

	for _, cookie := range s.Cookies {
		req.AddCookie(cookie)
	}

	response, responseErr := s.FiberRouter.Test(req, 1)
	assert.Equal(t, responseErr, nil, "")

	return response
}
