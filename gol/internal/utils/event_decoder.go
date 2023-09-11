package utils

import (
	json "github.com/json-iterator/go"
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"io"
)

// DecodeEvent decodes an event from an io.Reader, returning a pointer to the event
// or an error in case of failure.
func DecodeEvent(r io.Reader, event *models.Event) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	return decoder.Decode(event)
}
