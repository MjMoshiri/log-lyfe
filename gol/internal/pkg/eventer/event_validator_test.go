package eventer

import (
	"errors"
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidateEvent(t *testing.T) {
	// 1. Action is empty.
	t.Run("Empty Action", func(t *testing.T) {
		event := &models.Event{
			Actor:     "john_doe",
			Timestamp: time.Now(),
			EventID:   "12345",
		}
		err := ValidateEvent(event)
		assert.Equal(t, errors.New("missing required fields"), err)
	})

	// 2. Actor is empty.
	t.Run("Empty Actor", func(t *testing.T) {
		event := &models.Event{
			Action:    "login",
			Timestamp: time.Now(),
			EventID:   "12345",
		}
		err := ValidateEvent(event)
		assert.Equal(t, errors.New("missing required fields"), err)
	})

	// 3. Timestamp is zero.
	t.Run("Zero Timestamp", func(t *testing.T) {
		event := &models.Event{
			Action:  "login",
			Actor:   "john_doe",
			EventID: "12345",
		}
		err := ValidateEvent(event)
		assert.Equal(t, errors.New("missing required fields"), err)
	})

	// 4. EventID is empty.
	t.Run("Empty EventID", func(t *testing.T) {
		event := &models.Event{
			Action:    "login",
			Actor:     "john_doe",
			Timestamp: time.Now(),
		}
		err := ValidateEvent(event)
		assert.Equal(t, errors.New("missing required fields"), err)
	})

	// 5. All fields are missing.
	t.Run("Multiple Fields Missing", func(t *testing.T) {
		event := &models.Event{
			Actor: "john_doe",
		}
		err := ValidateEvent(event)
		assert.Equal(t, errors.New("missing required fields"), err)
	})
}
