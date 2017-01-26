package log

import (
	"dude_expenses/app"
	"net/http"
)

type loggingHandler struct {
	env  *app.Env
	next http.Handler
}

func WithLogging(env *app.Env, next http.Handler) loggingHandler {
	return loggingHandler{env: env, next: next}
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := HttpLogger(h.env, r)
	lrw := &loggingResponseWriter{w: w}
	h.next.ServeHTTP(lrw, r)
	logger.SetStatus(lrw.Status)
	defer logger.Log()
}

type loggingResponseWriter struct {
	w      http.ResponseWriter
	Status int
}

func (lrw *loggingResponseWriter) WriteHeader(status int) {
	lrw.Status = status
	lrw.w.WriteHeader(status)
}

func (lrw loggingResponseWriter) Header() http.Header {
	return lrw.w.Header()
}

func (lrw loggingResponseWriter) Write(b []byte) (int, error) {
	return lrw.w.Write(b)
}
