package parser

import (
	json "github.com/json-iterator/go"
	"io"
)

// FromJSON decodes an event from an io.Reader, returning a pointer to the event
// or an error in case of failure.
func FromJSON(r io.Reader, event interface{}) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	return decoder.Decode(event)
}
