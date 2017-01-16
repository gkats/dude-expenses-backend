package handler

import (
	"errors"
	"net/http"
)

type Error interface {
	error
	Status() int
	ValidationErrors() map[string][]string
}

type HandlerError struct {
	statusCode       int
	validationErrors map[string][]string
	Err              error
}

func (err HandlerError) Error() string {
	return err.Err.Error()
}

func (err HandlerError) Status() int {
	return err.statusCode
}

func (err HandlerError) ValidationErrors() map[string][]string {
	return err.validationErrors
}

func (err *HandlerError) SetValidationErrors(errors map[string][]string) {
	err.validationErrors = errors
}

func newHandlerError(status int, err error) HandlerError {
	return HandlerError{statusCode: status, Err: err}
}

func BadRequest() HandlerError {
	return newHandlerError(http.StatusBadRequest, errors.New("Bad request"))
}

func UnprocessableEntity(validationErrors map[string][]string) HandlerError {
	apiError := newHandlerError(http.StatusUnprocessableEntity, errors.New("Resource invalid"))
	apiError.SetValidationErrors(validationErrors)
	return apiError
}

func InternalServerError() HandlerError {
	return newHandlerError(http.StatusInternalServerError, errors.New("Something went wrong"))
}
