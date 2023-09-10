package models

import "time"

// Event is the representation of an event in the database.
// struct kept flat for simplicity and faster serialization. In production, this can be nested.
type Event struct {
	// Bucket is the partition key for the event table.
	Bucket int64 `json:"-" db:"bucket"`
	// ----------------------------------------
	// EventMetadata contains the universally relevant fields for any audit log entry.
	Action    string    `json:"action" db:"action"`
	Actor     string    `json:"actor" db:"actor"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	EventID   string    `json:"event_id" db:"event_id"`
	// In production, more fields like
	// SessionID, ReferenceID, etc., can be added for better granularity.
	// ----------------------------------------
	// AppInfo contains details about the application or system where the event originated.
	Version string `json:"version" db:"version"`
	// In production, additional fields like
	// InstanceID, Region, ServiceName, and Host can be added for richer context.
	// ----------------------------------------
	// ActionMetadata contains any additional details specific to the action taken.
	ActionMeta map[string]interface{} `json:"action_metadata"    db:"action_metadata"`
	// ActorMetadata contains additional details about the actor.
	ActorMeta map[string]interface{} `json:"actor_metadata"     db:"actor_metadata"`
}

// Clear clears the event struct for reuse.
func (e *Event) Clear() {
	e.Action = ""
	e.Actor = ""
	e.Timestamp = time.Time{}
	e.EventID = ""
	e.Version = ""
	e.ActionMeta = nil
	e.ActorMeta = nil
}
