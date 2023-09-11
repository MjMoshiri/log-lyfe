package handlers

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"github.com/mjmoshiri/log-lyfe/gol/internal/pkg"
	"github.com/mjmoshiri/log-lyfe/gol/internal/pkg/parser"
	"net/http"
	"time"
)

// HandleEventRequest handles Post requests to /event with Timeout
func (h *AppHandler) HandleEventRequest(w http.ResponseWriter, r *http.Request) {
	// Timeout for the whole process
	timeoutChannel := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeoutChannel <- true
	}()

	select {
	case <-h.processEventRequest(r, w):
		return
	case <-timeoutChannel:
		http.Error(w, "Request timed out", http.StatusRequestTimeout)
		return
	}
}

// processEventRequest the actual processing of the event request happens here
func (h *AppHandler) processEventRequest(r *http.Request, w http.ResponseWriter) chan bool {
	doneChannel := make(chan bool, 1)

	go func() {
		defer close(doneChannel)

		// Check for JSON content type
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			http.Error(w, "Invalid Content-Type, expected 'application/json'", http.StatusUnsupportedMediaType)
			return
		}

		// Get an event from the pool
		event := models.EventPool.Get().(*models.Event)
		defer models.EventPool.Put(event)

		// Decode the body
		err := parser.FromJSON(r.Body, event)
		if err != nil {
			http.Error(w, "Error decoding the event", http.StatusBadRequest)
			return
		}

		// Validate the event
		err = pkg.ValidateEvent(event)
		if err != nil {
			http.Error(w, "Invalid event data", http.StatusBadRequest)
			return
		}

		// Insert into the database
		err = h.DB.Insert(event)
		if err != nil {
			http.Error(w, "Failed to insert event into database", http.StatusInternalServerError)
			return
		}

		// If everything worked, return 200
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("Event processed successfully"))
		if err != nil {
			// TODO: log error
			return
		}
	}()

	return doneChannel
}
