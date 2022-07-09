// error library takes an id and an error message and we send this customized error object
// removing error logic from the handlers so if changes happen it won't break the code
// we can just swap it with new
package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"` // omits when you don't set it
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound, // return 404
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError, // returns 500
	}
}
