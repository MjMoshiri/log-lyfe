package pkg

import (
	"errors"
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
)

// ValidateEvent validates whether the event metadata of an event is valid.
// In production, this method would be more complex, have a set of rules, and be more granular.
func ValidateEvent(event *models.Event) error {
	if event.Action == "" || event.Actor == "" || event.Timestamp.IsZero() || event.EventID == "" {
		return errors.New("missing required fields")
	}
	return nil
}
