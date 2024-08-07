package customerrors

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	errorconst "github.com/uvulpos/go-svelte/basic-utils/customerrors/custom-error-const"
)

type DatabaseError struct {
	ID string

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	requestingUserID string

	sqlQuery         string
	sqlData          map[string]interface{}
	errorContextData string

	error error
}

func NewDatabaseError(err error, userID string, errorContextData string, sqlQuery string, sqlData SqlData) *DatabaseError {
	return &DatabaseError{
		ID: uuid.New().String(),

		errorIdentifier: errorconst.ERROR_IDENTIFIER_DATABASE,
		httpUserMessage: errorconst.INTERNAL_SERVER_ERROR_MESSAGE,
		httpStatus:      http.StatusInternalServerError,

		requestingUserID: userID,

		sqlQuery:         sqlQuery,
		sqlData:          sqlData,
		errorContextData: errorContextData,

		error: err,
	}
}

func (e *DatabaseError) Error() string {
	return e.error.Error()
}

func (e *DatabaseError) ErrorType() (errorIdentifier errorconst.ErrorIdentifier) {
	return e.errorIdentifier
}

func (e *DatabaseError) HttpError() (int, errorconst.ErrorIdentifier, string) {
	if debugMode {
		errormessage := fmt.Sprintf("[%s] #%s <br> %s <br> %s", e.errorIdentifier, e.ID, e.httpUserMessage, e.error.Error())
		return e.httpStatus, e.errorIdentifier, errormessage
	}
	return e.httpStatus, e.errorIdentifier, e.httpUserMessage
}

func (e *DatabaseError) LoggerError() (time.Time, int, errorconst.ErrorIdentifier, string, string, string) {
	contextData := fmt.Sprintf(
		"contextdata: %s, sqlQuery: %s, sqlData: %v",
		e.errorContextData,
		e.sqlQuery,
		e.sqlData,
	)
	return time.Now(), e.httpStatus, e.errorIdentifier, e.requestingUserID, contextData, e.error.Error()
}
