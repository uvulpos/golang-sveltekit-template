package customerrors

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	badrequestconstraints "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/bad-request-constraints"
	errorconst "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/custom-error-const"
)

type BadRequestError struct {
	ID string

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	inputKey   string
	constraint string

	error error
}

func NewBadRequestError(badRequestContraint *badrequestconstraints.BadRequestContraint) *BadRequestError {
	model := &BadRequestError{
		ID: uuid.New().String(),

		errorIdentifier: errorconst.ERROR_IDENTIFIER_BAD_REQUEST,
		httpUserMessage: errorconst.BAD_REQUEST_ERROR_MESSAGE,
		httpStatus:      http.StatusBadRequest,

		inputKey:   badRequestContraint.KeyName,
		constraint: badRequestContraint.ConstraintMessage,

		error: fmt.Errorf("bad request [key: %s] %s", badRequestContraint.KeyName, badRequestContraint.ConstraintMessage),
	}

	fmt.Printf("🌪️ %s\n", model.GetDeveloperMessage())

	return model
}

func (e *BadRequestError) Error() string {
	return e.error.Error()
}

func (e *BadRequestError) ErrorType() errorconst.ErrorIdentifier {
	return e.errorIdentifier
}

func (e *BadRequestError) HttpError() (int, errorconst.ErrorIdentifier, string) {

	errormessage := e.GetUserMessage()
	if debugMode {
		errormessage = e.GetDeveloperMessage()
	}

	return e.httpStatus, e.errorIdentifier, errormessage
}

func (e *BadRequestError) LoggerError() (time.Time, int, errorconst.ErrorIdentifier, string, string, string) {
	return time.Now(), e.httpStatus, e.errorIdentifier, e.inputKey, e.constraint, ""
}

func (e *BadRequestError) GetDeveloperMessage() string {
	return fmt.Sprintf("[%s] #%s <br> %s <br> %s: %s", e.errorIdentifier, e.ID, e.httpUserMessage, e.inputKey, e.constraint)
}

func (e *BadRequestError) GetUserMessage() string {
	return fmt.Sprintf("(#%s) %s", e.ID, e.httpUserMessage)
}
