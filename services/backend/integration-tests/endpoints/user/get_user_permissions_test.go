package user

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/uvulpos/golang-sveltekit-template/integration-tests/helper/setup"
	"gotest.tools/assert"

	httpUserModel "github.com/uvulpos/golang-sveltekit-template/src/resources/user/http/http-models"
)

func TestGetUserPermissions(t *testing.T) {
	testData := setup.SetupTest(t)
	testData.AuthenticateAsUser(t, "TIRIEDL")

	response := testData.MakeRequest(t, http.MethodGet, "/api/v1/self", nil)
	defer response.Body.Close()

	responseBytes, responseBytesErr := io.ReadAll(response.Body)
	assert.Equal(t, responseBytesErr, nil, "could not read body of http request")

	var selfUserInformation httpUserModel.SelfInformationModel
	unmarshalErr := json.Unmarshal(responseBytes, &selfUserInformation)
	assert.Equal(t, unmarshalErr, nil, "could not unmarshal http request")

	assert.Equal(t, selfUserInformation.Username, "TIRIEDL", "blabla")
}
