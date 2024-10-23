package customerrors

import (
	"errors"
	"net/http"
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
	assert.Equal(t, httpUserMessage, customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
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
	assert.Equal(t, httpUserMessage, customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
}
