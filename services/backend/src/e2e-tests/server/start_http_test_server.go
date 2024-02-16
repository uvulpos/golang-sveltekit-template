package e2etests

import (
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/uvulpos/go-svelte/src/helper/config"
	webapp "github.com/uvulpos/go-svelte/src/web-app"
)

type httpTestServer struct {
	server *fiber.App
}

func NewHttpTestServer() *httpTestServer {
	configuration := config.LoadData()
	server := webapp.NewApp(configuration).ReturnAppInE2EMode()

	return &httpTestServer{
		server: server,
	}
}

func (s *httpTestServer) MakeRequest(methode, url string) (*http.Response, error) {
	req := httptest.NewRequest(methode, url, nil)
	response, responseErr := s.server.Test(req, 2)
	return response, responseErr
}
