package users

import (
	"net/http"
	"testing"

	testServer "github.com/uvulpos/go-svelte/src/e2e-tests/server"
)

func TestUserLogin(t *testing.T) {
	_, respErr := testServer.NewHttpTestServer().MakeRequest(http.MethodGet, "/api/v1/hallo")
	if respErr != nil {
		t.FailNow()
	}
}
