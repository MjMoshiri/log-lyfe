package middlewares

import (
	"net/http"
)

// QueryAuthMiddleware is a middleware that checks authentication of a query request
// TODO: Implement a generic middleware that take a AuthObject interface (e.g. the endpoint, key, value, method)
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
