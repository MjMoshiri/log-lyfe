package middlewares

import "net/http"

// MiddlewareFunc is a function that takes a handler and returns a handler.
// This is used to wrap handlers with middleware.
type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc
