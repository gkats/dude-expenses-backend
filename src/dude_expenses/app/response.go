package app

import (
	"net/http"
)

type Response interface {
	Status() int
	Body() interface{}
}

type handlerResponse struct {
	status int
	body   interface{}
}

func (r handlerResponse) Status() int {
	return r.status
}

func (r handlerResponse) Body() interface{} {
	return r.body
}

type ErrorResponse struct {
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors,omitempty"`
}

func HandlerResponse(status int, body interface{}) Response {
	return handlerResponse{status: status, body: body}
}

func BadRequest() Response {
	body := ErrorResponse{Message: "Bad request"}
	return HandlerResponse(http.StatusBadRequest, body)
}

func Unauthorized() Response {
	body := ErrorResponse{Message: "Unauthorized"}
	return HandlerResponse(http.StatusUnauthorized, body)
}

func NotFound() Response {
	body := ErrorResponse{Message: "Resource not found"}
	return HandlerResponse(http.StatusNotFound, body)
}

func UnprocessableEntity(errors map[string][]string) Response {
	body := ErrorResponse{Message: "Resource invalid", Errors: errors}
	return HandlerResponse(http.StatusUnprocessableEntity, body)
}

func InternalServerError() Response {
	body := ErrorResponse{Message: "Something went wrong"}
	return HandlerResponse(http.StatusInternalServerError, body)
}

func OK(body interface{}) Response {
	return HandlerResponse(http.StatusOK, body)
}

func Created(body interface{}) Response {
	return HandlerResponse(http.StatusCreated, body)
}
