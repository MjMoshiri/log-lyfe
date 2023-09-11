package handlers

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"github.com/mjmoshiri/log-lyfe/gol/internal/pkg/eventer"
	"github.com/mjmoshiri/log-lyfe/gol/internal/pkg/parser"
	"net/http"
	"time"
)

// HandleEventRequest processes POST requests to /event, ensuring the operation completes within a set timeout.
func (h *AppHandler) HandleEventRequest(w http.ResponseWriter, r *http.Request) {
	// Timeout for the whole process
	timeoutChannel := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeoutChannel <- true
	}()

	select {
	case <-h.processEventRequest(w, r):
		return
	case <-timeoutChannel:
		http.Error(w, "Request timed out", http.StatusRequestTimeout)
		return
	}
}

// processEventRequest performs the actual processing of the event request. It validates the content type,
// decodes the event, validates the event data, and inserts it into the database.
func (h *AppHandler) processEventRequest(w http.ResponseWriter, r *http.Request) chan bool {
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
		event := eventer.EventPool.Get().(*models.Event)
		defer eventer.EventPool.Put(event)

		// Decode the body
		err := parser.FromJSON(r.Body, event)
		if err != nil {
			http.Error(w, "Error decoding the event", http.StatusBadRequest)
			return
		}

		// Validate the event
		err = eventer.ValidateEvent(event)
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
