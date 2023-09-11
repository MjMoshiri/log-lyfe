package models

import (
	"sync"
)

// EventPool is a sync.Pool that contains events.
var EventPool = sync.Pool{
	New: func() interface{} {
		return &Event{}
	},
}
