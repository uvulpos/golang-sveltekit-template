package customerrors

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	errorconst "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/custom-error-const"
)

type DatabaseTransactionCommitError struct {
	ID string

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	errorContextData string

	error error
}

func NewDatabaseTransactionCommitError(err error, errorContextData string) *DatabaseTransactionCommitError {
	model := &DatabaseTransactionCommitError{
		ID: uuid.New().String(),

		errorIdentifier: errorconst.ERROR_IDENTIFIER_DATABASE_TRANSACTION_NOT_STARTED,
		httpUserMessage: errorconst.INTERNAL_SERVER_ERROR_MESSAGE,
		httpStatus:      http.StatusInternalServerError,

		errorContextData: errorContextData,

		error: err,
	}

	fmt.Printf("🚨 %s\n", model.GetDeveloperMessage())

	return model
}

func (e *DatabaseTransactionCommitError) Error() string {
	return e.error.Error()
}

func (e *DatabaseTransactionCommitError) ErrorType() (errorIdentifier errorconst.ErrorIdentifier) {
	return e.errorIdentifier
}

func (e *DatabaseTransactionCommitError) HttpError() (int, errorconst.ErrorIdentifier, string) {

	errormessage := e.GetUserMessage()
	if debugMode {
		errormessage = e.GetDeveloperMessage()
	}

	return e.httpStatus, e.errorIdentifier, errormessage
}

func (e *DatabaseTransactionCommitError) LoggerError() (time.Time, int, errorconst.ErrorIdentifier, string, string, string) {
	return time.Now(), e.httpStatus, e.errorIdentifier, "", "", e.error.Error()
}

func (e *DatabaseTransactionCommitError) GetDeveloperMessage() string {
	return fmt.Sprintf("[%s] #%s <br> %s <br> %s", e.errorIdentifier, e.ID, e.httpUserMessage, e.error.Error())
}

func (e *DatabaseTransactionCommitError) GetUserMessage() string {
	return fmt.Sprintf("(#%s) %s", e.ID, e.httpUserMessage)
}
