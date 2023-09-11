package handlers

import (
	json "github.com/json-iterator/go"
	"github.com/mjmoshiri/log-lyfe/gol/internal/pkg"
	"net/http"
	"strconv"
)

// HandleQueryRequest processes query requests, expecting JSON content and optional fetch size.
// TODO: Implement a timeout mechanism for the request.
func (h *AppHandler) HandleQueryRequest(w http.ResponseWriter, r *http.Request) {
	// Ensure JSON content type
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid Content-Type, expected 'application/json'", http.StatusUnsupportedMediaType)
		return
	}

	// Convert request body to map literal for filtering
	filters, err := pkg.ConvertToMapLiteral(r.Body)
	if err != nil {
		http.Error(w, "Error converting body to map literal", http.StatusBadRequest)
		return
	}

	// Extract and validate fetch size from headers
	fetchSizeString := r.Header.Get("Fetch-Size")
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
	// Retrieve events based on filters and fetch size
	events, err := h.DB.Find(filters, uint(fetchSize))
	if err != nil {
		http.Error(w, "Error fetching events from database", http.StatusInternalServerError)
		return
	}
	// Encode and return the retrieved events as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		http.Error(w, "Error encoding results to JSON", http.StatusInternalServerError)
		return
	}
}
