package middlewares

import (
	"net/http"
)

// AuthMiddleware is a middleware that checks for a valid Authorization header
// this is a very simple implementation, as a proof of concept.
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
