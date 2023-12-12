package errors

import (
	"fmt"
	"net/http"

	"github.com/uvulpos/go-svelte/src/helper/logger"
)

type Error struct {
	code    int
	msg     string
	errType ErrorType
}

func newError(code int, msg string, errType ErrorType) *Error {
	return &Error{
		code:    code,
		msg:     msg,
		errType: errType,
	}
}

func NewError(code int, msg string, errType ErrorType) *Error {
	newErr := newError(code, msg, errType)
	var std = logger.Error
	outErr := std.Output(3, newErr.Log())
	if outErr != nil {
		logger.Error.Println(outErr)
	}

	return newErr
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Message() string {
	return e.msg
}

func (e *Error) Type() string {
	return e.errType.value
}

func (e *Error) Log() string {
	return fmt.Sprintf("[msg='%s' err='%s']", e.msg, e.Type())
}

func NewBadRequestError(msg string, errType ErrorType) *Error {
	return NewError(http.StatusBadRequest, msg, errType)
}

func NewNotFoundError(msg string, errType ErrorType) *Error {
	return NewError(http.StatusNotFound, msg, errType)
}

func NewIDNotFoundError(ID int, entityName string, errType ErrorType) *Error {
	msg := fmt.Sprintf("%s of ID %d not found", entityName, ID)
	return NewError(http.StatusNotFound, msg, errType)
}

func NewInternalServerError(msg string, errType ErrorType) *Error {
	return NewError(http.StatusInternalServerError, msg, errType)
}

func NewInternalServerErrorApp(msg string) *Error {
	return NewError(http.StatusInternalServerError, msg, AppError)
}

func (e *Error) ToHttpError() error {
	return fmt.Errorf("[ %s ] %s", e.Type(), e.Message())
}
