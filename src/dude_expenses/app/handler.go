package app

import (
	"encoding/json"
	"net/http"
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request) Response
}

func Handle(next Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)

		response := CommonHeaders(next).ServeHTTP(w, r)
		w.WriteHeader(response.Status())
		encoder.Encode(response.Body())
	})
}

func ParseRequestBody(r *http.Request, params interface{}) error {
	return json.NewDecoder(r.Body).Decode(params)
}

type commonHeadersHandler struct {
	next Handler
}

func CommonHeaders(next Handler) Handler {
	return &commonHeadersHandler{next: next}
}

func (h *commonHeadersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) Response {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	return h.next.ServeHTTP(w, r)
}
