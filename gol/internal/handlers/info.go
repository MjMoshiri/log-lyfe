package handlers

import (
	json "github.com/json-iterator/go"
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"net/http"
	"os"
	"runtime"
)

// InfoHandler handles requests to /info
// returns system information, in production this could be used for versioning
func (h *AppHandler) InfoHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Failed to get hostname", http.StatusInternalServerError)
		return
	}
	info := models.SystemInfo{
		Hostname:  hostname,
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		CPUs:      runtime.NumCPU(),
		GoVersion: runtime.Version(),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(info)
	if err != nil {
		// TODO: Log error
		return
	}
}
