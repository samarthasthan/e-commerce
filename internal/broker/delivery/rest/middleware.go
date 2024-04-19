package rest

import (
	"net/http"
	"time"

	"github.com/samarthasthan/e-commerce/pkg/logger"
)

// LoggingMiddleware logs details of incoming HTTP requests and responses.
func LoggingMiddleware(log *logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Log request details
		log.Infof("Received request: method=%s, url=%s, headers=%v", r.Method, r.URL, r.Header)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Calculate the duration
		duration := time.Since(startTime)

		// Log response details
		log.Infof("Completed request: method=%s, url=%s, status=%d, duration=%s",
			r.Method, r.URL, w.WriteHeader, duration)
	})
}
