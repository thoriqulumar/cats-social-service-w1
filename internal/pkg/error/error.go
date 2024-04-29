package error

import "errors"

// CustomError represents a custom error with code and message
type CustomError struct {
	Code    int
	Message string
}

// New creates a new CustomError instance with the specified code and message
func New(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

// Error returns the error message
func (e *CustomError) Error() string {
	return e.Message
}

func GetCode(err error) int {
	var cerr *CustomError
	ok := errors.As(err, &cerr)
	if !ok {
		return 0
	}
	return cerr.Code
}
