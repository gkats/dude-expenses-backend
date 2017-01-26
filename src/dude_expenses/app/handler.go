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
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)

		response := next.ServeHTTP(w, r)
		w.WriteHeader(response.Status())
		encoder.Encode(response.Body())
	})
}

func ParseRequestBody(r *http.Request, params interface{}) error {
	return json.NewDecoder(r.Body).Decode(params)
}
