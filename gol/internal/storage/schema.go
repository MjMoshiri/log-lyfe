package storage

import (
	"github.com/scylladb/gocqlx/v2/table"
)

var EventTable = table.New(table.Metadata{
	Name:    "events",
	Columns: []string{"bucket", "action", "actor", "timestamp", "event_id", "version", "action_metadata", "actor_metadata"},
	PartKey: []string{"bucket"},
	SortKey: []string{"timestamp"},
})
