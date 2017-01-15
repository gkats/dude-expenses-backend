package handler

import (
	"app"
	"encoding/json"
	"net/http"
)

type AppHandlerFunc func(w http.ResponseWriter, r *http.Request) (HandlerResponse, Error)

type appHandler struct {
	env *app.Env
	h   func(*app.Env) AppHandlerFunc
}

func (handler appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	response, err := handler.h(handler.env)(w, r)
	if err != nil {
		w.WriteHeader(err.Status())
		encoder.Encode(newErrorResponse(err))
	} else {
		w.WriteHeader(response.Status)
		encoder.Encode(&response.Body)
	}
}

func AppHandler(env *app.Env, h func(env *app.Env) AppHandlerFunc) appHandler {
	return appHandler{env: env, h: h}
}

type errorResponse struct {
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors,omitempty"`
}

func newErrorResponse(err Error) *errorResponse {
	return &errorResponse{Message: err.Error(), Errors: err.ValidationErrors()}
}

type HandlerResponse struct {
	Status int
	Body   interface{}
}

func NewHandlerResponse(status int, body interface{}) HandlerResponse {
	return HandlerResponse{Status: status, Body: body}
}

func ParseRequestBody(r *http.Request, params interface{}) error {
	return json.NewDecoder(r.Body).Decode(params)
}
