package customerrors

import (
	"time"

	errorconst "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/custom-error-const"
)

func GetHttpError(err ErrorInterface) (int, errorconst.ErrorIdentifier, string) {
	return err.HttpError()
}

func GetLoggerError(err ErrorInterface) (time.Time, int, errorconst.ErrorIdentifier, string, string, string) {
	return err.LoggerError()
}
