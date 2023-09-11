// Package handlers provides HTTP request handlers for various application endpoints.
package handlers

import (
	"net/http"
)

// HealthCheckHandler responds to /ok requests.
// While it currently returns a simple 200 OK, in production, it can be enhanced
// to provide more detailed health information, such as database connection status.
func (h AppHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		// TODO: Log error
		return
	}
}
