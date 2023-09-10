package handlers

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/utils"
	"net/http"
	"time"
)

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

		// Decode the body
		event, err := utils.DecodeEvent(r.Body)
		if err != nil {
			http.Error(w, "Error decoding the event", http.StatusBadRequest)
			return
		}

		// Validate the event
		err = utils.ValidateEvent(*event)
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
		w.Write([]byte("Event processed successfully"))
	}()

	return doneChannel
}
