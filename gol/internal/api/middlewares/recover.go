package middlewares

import (
	"log"
	"net/http"
)

// RecoverMiddleware returns a middleware function that recovers from panics
// during request processing, logging the error and returning a 500 status.
// For production use, consider pairing this with a logging middleware.
func RecoverMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover from panic: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		next(w, r)
	}
}
