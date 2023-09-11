package models

import "time"

// Event represents an event stored in the database.
// The structure is kept flat for efficient serialization, though in production, nested structures may be more appropriate.
type Event struct {
	// Bucket serves as the partition key for the event table.
	Bucket int64 `json:"-" db:"bucket"`

	// EventMetadata holds universally relevant fields for any audit log entry.
	Action    string    `json:"action" db:"action"`
	Actor     string    `json:"actor" db:"actor"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	EventID   string    `json:"event_id" db:"event_id"`
	// Consider adding fields like SessionID, ReferenceID, etc., in production for detailed tracking.

	// AppInfo provides details about the originating application or system.
	Version string `json:"version" db:"version"`
	// Fields like InstanceID, Region, ServiceName, and Host can be added in production for comprehensive context.

	// ActionMetadata contains specific details about the action.
	ActionMeta map[string]string `json:"action_metadata"    db:"action_metadata"`
	// ActorMetadata holds additional information about the actor.
	ActorMeta map[string]string `json:"actor_metadata"     db:"actor_metadata"`
}

// Clear resets the event's fields, preparing it for reuse.
func (e *Event) Clear() {
	e.Action = ""
	e.Actor = ""
	e.Timestamp = time.Time{}
	e.EventID = ""
	e.Version = ""
	e.ActionMeta = map[string]string{}
	e.ActorMeta = map[string]string{}
}
