package utils

import (
	json "github.com/json-iterator/go"
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"io"
)

// DecodeEvent decodes an event from an io.Reader, returning a pointer to the event
// or an error in case of failure.
func DecodeEvent(r io.Reader) (*models.Event, error) {
	event := EventPool.Get().(*models.Event)
	event.Clear()
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(event)
	if err != nil {
		return nil, err
	}
	return event, nil
}
