package customerrors

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	errorconst "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/custom-error-const"
)

type DatabaseNotFoundError struct {
	ID string

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	requestingUserID string

	sqlQuery string
	sqlData  map[string]interface{}

	error error
}

func NewDatabaseNotFoundError(err error, userID string, sqlQuery string, sqlData SqlData) *DatabaseNotFoundError {
	model := &DatabaseNotFoundError{
		ID: uuid.New().String(),

		errorIdentifier: errorconst.ERROR_IDENTIFIER_DATABASE_NOT_FOUND,
		httpUserMessage: errorconst.DATABASE_NOT_FOUND_ERROR_MESSAGE,
		httpStatus:      http.StatusNotFound,

		requestingUserID: userID,

		sqlQuery: sqlQuery,
		sqlData:  sqlData,

		error: err,
	}
	fmt.Printf("🚨 %s\n", model.GetDeveloperMessage())

	return model
}

func (e *DatabaseNotFoundError) Error() string {
	return e.error.Error()
}

func (e *DatabaseNotFoundError) ErrorType() (errorIdentifier errorconst.ErrorIdentifier) {
	return e.errorIdentifier
}

func (e *DatabaseNotFoundError) HttpError() (int, errorconst.ErrorIdentifier, string) {

	errormessage := e.GetUserMessage()
	if debugMode {
		errormessage = e.GetDeveloperMessage()
	}

	return e.httpStatus, e.errorIdentifier, errormessage
}

func (e *DatabaseNotFoundError) LoggerError() (time.Time, int, errorconst.ErrorIdentifier, string, string, string) {
	contextData := fmt.Sprintf(
		"sqlQuery: %s, sqlData: %v",
		e.sqlQuery,
		e.sqlData,
	)
	return time.Now(), e.httpStatus, e.errorIdentifier, e.requestingUserID, contextData, e.error.Error()
}

func (e *DatabaseNotFoundError) GetDeveloperMessage() string {
	return fmt.Sprintf("[%s] #%s <br>\n %s <br>\n %s", e.errorIdentifier, e.ID, e.httpUserMessage, e.error.Error())
}

func (e *DatabaseNotFoundError) GetUserMessage() string {
	return fmt.Sprintf("(#%s) %s", e.ID, e.httpUserMessage)
}
