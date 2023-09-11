// Package middlewares provides middleware functions for the API.
package middlewares

import (
	"net/http"
)

// AuthMiddleware returns a middleware function that checks for a valid
// Authorization header against the provided key. It's a basic implementation
// primarily for demonstration purposes.
func AuthMiddleware(key string) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader != key {
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte("Unauthorized"))
				if err != nil {
					// TODO: log error
				}
				return
			}
			next(w, r)
		}
	}
}
