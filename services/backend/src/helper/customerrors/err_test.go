package customerrors

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	customerrorconst "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/custom-error-const"
	"gotest.tools/assert"
)

func TestHttpDatabaseErrorHandling(t *testing.T) {
	myDatabaseError := NewDatabaseError(
		errors.New("Database error"),
		uuid.New().String(),
		"SELECT * FROM users",
		"",
		map[string]interface{}{"id": 123},
	)

	httpStatus, errorIdentifier, httpUserMessage := GetHttpError(myDatabaseError)

	assert.Equal(t, httpStatus, http.StatusInternalServerError)
	assert.Equal(t, errorIdentifier, customerrorconst.ERROR_IDENTIFIER_DATABASE)
	assert.Assert(
		t,
		strings.HasSuffix(httpUserMessage, ") "+customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE),
		"Errormessage does not have defined error type suffix",
	)
}

func TestHttpInternalServerErrorHandling(t *testing.T) {
	myDatabaseError := NewInternalServerError(
		errors.New("Internal server error"),
		uuid.New().String(),
		"Error context data",
	)

	httpStatus, errorIdentifier, httpUserMessage := GetHttpError(myDatabaseError)

	assert.Equal(t, httpStatus, http.StatusInternalServerError)
	assert.Equal(t, errorIdentifier, customerrorconst.ERROR_IDENTIFIER_INTERNAL_SERVER_ERROR)

	assert.Assert(
		t,
		strings.HasSuffix(httpUserMessage, ") "+customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE),
		"Errormessage does not have defined error type suffix",
	)
}
