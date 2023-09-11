package parser

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"github.com/mjmoshiri/log-lyfe/gol/internal/pkg/eventer"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDecodeEvent(t *testing.T) {
	// 1. Successful Decoding Test
	t.Run("Successful Decoding", func(t *testing.T) {
		validJSON := `{
			"action": "login",
			"actor": "john_doe",
			"timestamp": "2023-09-10T10:00:00Z",
			"event_id": "12345",
			"version": "1.0",
			"action_metadata": {"ip": "127.0.0.1"},
			"actor_metadata": {"role": "admin"}
		}`
		event := eventer.EventPool.Get().(*models.Event)
		defer eventer.EventPool.Put(event)
		err := FromJSON(strings.NewReader(validJSON), event)
		assert.NoError(t, err)
		assert.Equal(t, "login", event.Action)
	})

	// 2. Unknown Fields Test
	t.Run("Unknown Fields", func(t *testing.T) {
		unknownFieldJSON := `{
			"unknown_field": "unknown",
			"action": "login",
			"actor": "john_doe",
			"timestamp": "2023-09-10T10:00:00Z",
			"event_id": "12345",
			"version": "1.0"
		}`
		event := eventer.EventPool.Get().(*models.Event)
		defer eventer.EventPool.Put(event)
		err := FromJSON(strings.NewReader(unknownFieldJSON), event)
		assert.Error(t, err)
	})

	// 3. Invalid JSON Test
	t.Run("Invalid JSON", func(t *testing.T) {
		invalidJSON := `{"action": "login", "actor": "john_doe"`
		event := eventer.EventPool.Get().(*models.Event)
		defer eventer.EventPool.Put(event)
		err := FromJSON(strings.NewReader(invalidJSON), event)
		assert.Error(t, err)
	})

	// 4. Metadata Type Mismatch Test
	t.Run("Metadata Type Mismatch", func(t *testing.T) {
		wrongTypeJSON := `{
			"action": "login",
			"actor": "john_doe",
			"timestamp": "2023-09-10T10:00:00Z",
			"event_id": "12345",
			"version": "1.0",
			"action_metadata": []
		}`
		event := eventer.EventPool.Get().(*models.Event)
		defer eventer.EventPool.Put(event)
		err := FromJSON(strings.NewReader(wrongTypeJSON), event)
		assert.Error(t, err)
	})

	// 5. Invalid Date Format Test
	t.Run("Invalid Date Format", func(t *testing.T) {
		invalidDateJSON := `{
			"action": "login",
			"actor": "john_doe",
			"timestamp": "10-09-2023",
			"event_id": "12345",
			"version": "1.0"
		}`
		event := eventer.EventPool.Get().(*models.Event)
		defer eventer.EventPool.Put(event)
		err := FromJSON(strings.NewReader(invalidDateJSON), event)
		assert.Error(t, err)
	})

	// 6. Empty Input Test
	t.Run("Empty Input", func(t *testing.T) {
		event := eventer.EventPool.Get().(*models.Event)
		defer eventer.EventPool.Put(event)
		err := FromJSON(strings.NewReader(""), event)
		assert.Error(t, err)
	})

	// 7. Metadata As String Test
	t.Run("Metadata As String", func(t *testing.T) {
		metadataStringJSON := `{
			"action": "login",
			"actor": "john_doe",
			"timestamp": "2023-09-10T10:00:00Z",
			"event_id": "12345",
			"version": "1.0",
			"action_metadata": "string_data"
		}`
		event := eventer.EventPool.Get().(*models.Event)
		defer eventer.EventPool.Put(event)
		err := FromJSON(strings.NewReader(metadataStringJSON), event)
		assert.Error(t, err)
	})
}
