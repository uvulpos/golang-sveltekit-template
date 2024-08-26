package customerrors

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	errorconst "github.com/uvulpos/go-svelte/basic-utils/customerrors/custom-error-const"
)

type InternalServerError struct {
	ID string

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	errorContextData string

	requestingUserID string

	error error
}

func NewInternalServerError(err error, userID string, errorContextData string) *InternalServerError {
	model := &InternalServerError{
		ID: uuid.New().String(),

		errorIdentifier: errorconst.INTERNAL_SERVER_ERROR_MESSAGE,
		httpUserMessage: errorconst.INTERNAL_SERVER_ERROR_MESSAGE,
		httpStatus:      http.StatusInternalServerError,

		errorContextData: errorContextData,

		requestingUserID: userID,

		error: err,
	}

	fmt.Println("🚨")

	return model
}

func (e *InternalServerError) Error() string {
	return e.error.Error()
}

func (e *InternalServerError) ErrorType() (errorIdentifier errorconst.ErrorIdentifier) {
	return e.errorIdentifier
}

func (e *InternalServerError) HttpError() (int, errorconst.ErrorIdentifier, string) {
	if debugMode {
		errormessage := fmt.Sprintf("[%s] #%s <br> %s <br> %s", e.errorIdentifier, e.ID, e.httpUserMessage, e.error.Error())
		return e.httpStatus, e.errorIdentifier, errormessage
	}
	return e.httpStatus, e.errorIdentifier, e.httpUserMessage
}

func (e *InternalServerError) LoggerError() (time.Time, int, errorconst.ErrorIdentifier, string, string, string) {
	return time.Now(), e.httpStatus, e.errorIdentifier, e.requestingUserID, e.errorContextData, e.error.Error()
}
