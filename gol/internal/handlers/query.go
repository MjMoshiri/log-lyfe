package handlers

import (
	json "github.com/json-iterator/go"
	"github.com/mjmoshiri/log-lyfe/gol/internal/pkg"
	"net/http"
	"strconv"
)

// HandleQueryRequest handles a query request
// TODO: Add timeout either by server or by request
func (h *AppHandler) HandleQueryRequest(w http.ResponseWriter, r *http.Request) {
	// Check for JSON content type
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid Content-Type, expected 'application/json'", http.StatusUnsupportedMediaType)
		return
	}

	// Convert body to map literal
	filters, err := pkg.ConvertToMapLiteral(r.Body)
	if err != nil {
		http.Error(w, "Error converting body to map literal", http.StatusBadRequest)
		return
	}

	fetchSizeString := r.Header.Get("Fetch-Size")
	// convert to int
	fetchSize := 0
	if fetchSizeString != "" {
		fetchSize, err = strconv.Atoi(fetchSizeString)
		if err != nil {
			http.Error(w, "Error converting fetch size to integer", http.StatusBadRequest)
			return
		}
		if fetchSize < 0 {
			http.Error(w, "Fetch size must be greater than or equal to 0", http.StatusBadRequest)
			return
		}
	}
	// Find events with the given filters
	events, err := h.DB.Find(filters, uint(fetchSize))
	if err != nil {
		http.Error(w, "Error fetching events from database", http.StatusInternalServerError)
		return
	}

	// Convert results to JSON and write to response body
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		http.Error(w, "Error encoding results to JSON", http.StatusInternalServerError)
		return
	}
}
