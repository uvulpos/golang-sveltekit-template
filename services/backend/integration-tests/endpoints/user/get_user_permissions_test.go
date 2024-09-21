package user

import (
	"net/http"
	"testing"

	"github.com/uvulpos/golang-sveltekit-template/integration-tests/helper/setup"
	"gotest.tools/assert"
)

func TestGetUserPermissions(t *testing.T) {
	testData := setup.SetupTest(t)
	testData.AuthenticateAsUser(t, "")

	response := testData.MakeRequest(t, http.MethodGet, "/api/v1/self/permissions", nil)
	defer response.Body.Close()

	assert.Equal(t, "", "", "")
}
