package middlewares

import (
	"log"
	"net/http"
)

// RecoverMiddleware is a middleware that recovers from panics
// In production it's better to add a logging middleware before this one
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
