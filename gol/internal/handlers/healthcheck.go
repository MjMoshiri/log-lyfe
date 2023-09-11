package handlers

import (
	"net/http"
)

// HealthCheckHandler handles requests to /ok
// returns 200 OK, in production this could be used for health checks with more information e.g. database connection status
func (h AppHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		// TODO: Log error
		return
	}
}
