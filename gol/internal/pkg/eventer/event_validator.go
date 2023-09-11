package eventer

import (
	"errors"
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
)

// ValidateEvent checks the event's metadata for completeness.
// In a production setting, this validation would likely involve a comprehensive set of rules and checks.
func ValidateEvent(event *models.Event) error {
	if event.Action == "" || event.Actor == "" || event.Timestamp.IsZero() || event.EventID == "" {
		return errors.New("missing required fields")
	}
	return nil
}
