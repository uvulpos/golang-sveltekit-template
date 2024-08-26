package customerrors

import (
	"time"

	errorconst "github.com/uvulpos/go-svelte/basic-utils/customerrors/custom-error-const"
)

var debugMode bool = true

func SetDebugMode(isDebug bool) {
	debugMode = isDebug
}

type SqlData map[string]interface{}

type ErrorInterface interface {
	Error() (errormessage string)
	ErrorType() (errorIdentifier errorconst.ErrorIdentifier)
	HttpError() (httpStatus int, errorIdentifier errorconst.ErrorIdentifier, userMessage string)
	LoggerError() (now time.Time, httpStatus int, errorIdentifier errorconst.ErrorIdentifier, requestingUserID string, contextData string, nativeErrorMessage string)
}
