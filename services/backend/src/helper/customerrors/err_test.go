package customerrors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/google/uuid"
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
	assert.Equal(t, errorIdentifier, "DatabaseError")
	assert.Equal(t, httpUserMessage, "An error occurred while processing your request. Please try again later.")
}

func TestHttpInternalServerErrorHandling(t *testing.T) {
	myDatabaseError := NewInternalServerError(
		errors.New("Internal server error"),
		uuid.New().String(),
		"Error context data",
	)

	httpStatus, errorIdentifier, httpUserMessage := GetHttpError(myDatabaseError)

	assert.Equal(t, httpStatus, http.StatusInternalServerError)
	assert.Equal(t, errorIdentifier, "InternalServerError")
	assert.Equal(t, httpUserMessage, "An error occurred while processing your request. Please try again later.")
}
