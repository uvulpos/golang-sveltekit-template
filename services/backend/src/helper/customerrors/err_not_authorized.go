package customerrors

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	errorconst "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/custom-error-const"
)

type NotAuthorizedError struct {
	ID string

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	error error
}

func NewNotAuthorizedError() *NotAuthorizedError {
	model := &NotAuthorizedError{
		ID: uuid.New().String(),

		errorIdentifier: errorconst.ERROR_IDENTIFIER_UNAUTHORIZED,
		httpUserMessage: errorconst.NOT_AUTHORIZED_ERROR_MESSAGE,
		httpStatus:      http.StatusUnauthorized,

		error: errors.New("user is unauthorized"),
	}

	fmt.Printf("🚨 %s\n", model.GetDeveloperMessage())

	return model
}

func (e *NotAuthorizedError) Error() string {
	return e.error.Error()
}

func (e *NotAuthorizedError) ErrorType() (errorIdentifier errorconst.ErrorIdentifier) {
	return e.errorIdentifier
}

func (e *NotAuthorizedError) HttpError() (int, errorconst.ErrorIdentifier, string) {

	errormessage := e.GetUserMessage()
	if debugMode {
		errormessage = e.GetDeveloperMessage()
	}

	return e.httpStatus, e.errorIdentifier, errormessage
}

func (e *NotAuthorizedError) LoggerError() (time.Time, int, errorconst.ErrorIdentifier, string, string, string) {
	return time.Now(), e.httpStatus, e.errorIdentifier, "", "", e.error.Error()
}

func (e *NotAuthorizedError) GetDeveloperMessage() string {
	return fmt.Sprintf("[%s] #%s <br> %s <br> %s", e.errorIdentifier, e.ID, e.httpUserMessage, e.error.Error())
}

func (e *NotAuthorizedError) GetUserMessage() string {
	return fmt.Sprintf("(#%s) %s", e.ID, e.httpUserMessage)
}
