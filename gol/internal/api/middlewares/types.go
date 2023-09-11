package middlewares

import "net/http"

// MiddlewareFunc defines a function type that wraps an HTTP handler with middleware logic.
// It accepts a handler and returns a wrapped handler.
//
// Example:
//
//	func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
//	    return func(w http.ResponseWriter, r *http.Request) {
//	        log.Println(r.URL.Path)
//	        next(w, r)
//	    }
//	}
type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc
