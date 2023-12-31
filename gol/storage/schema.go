package storage

import (
	"github.com/scylladb/gocqlx/v2/table"
)

// EventTable defines the structure of the 'events' table in Cassandra.
var EventTable = table.New(table.Metadata{
	Name:    "events",
	Columns: []string{"bucket", "action", "actor", "timestamp", "event_id", "version", "action_metadata", "actor_metadata"},
	PartKey: []string{"bucket"},
	SortKey: []string{"timestamp"},
})
