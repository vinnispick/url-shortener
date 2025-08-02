package middleware

import (
	"net/http"
	"time"
	"url-shortener/internal/logger"
)

func LoggingMiddleware(log logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			log.Info("Received request: " + r.Method + " " + r.URL.String())

			next.ServeHTTP(w, r)

			duration := time.Since(start)
			log.Info("Processed request: " + r.Method + " " + r.URL.String() + " in " + duration.String())
		})
	}
}
