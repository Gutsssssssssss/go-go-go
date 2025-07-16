package middleware

import (
	"log/slog"
	"net/http"
)

func Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Debug("Request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
