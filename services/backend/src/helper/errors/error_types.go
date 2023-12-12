package errors

type ErrorType struct {
	value string
}

var (
	Forbidden = newErrorType("FORBIDDEN")

	AppError = newErrorType("APP-ERROR")
)

func newErrorType(name string) ErrorType {
	return ErrorType{
		value: name,
	}
}
