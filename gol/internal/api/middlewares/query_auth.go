package middlewares

import (
	"net/http"
)

// QueryAuthMiddleware returns a middleware function that authenticates
// GET requests to the "/query" endpoint using the provided key.
//
// TODO:
// Consider implementing a generic middleware that accepts an AuthObject interface
// to specify authentication details like endpoint, key, value, and method.
func QueryAuthMiddleware(key string) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet || r.URL.Path != "/query" {
				next(w, r)
				return
			}
			queryKeyHeader := r.Header.Get("Query-Key")
			if queryKeyHeader != key {
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
