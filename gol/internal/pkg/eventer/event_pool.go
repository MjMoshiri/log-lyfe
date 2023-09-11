// Package eventer provides utilities for working with events.
package eventer

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"sync"
)

// EventPool is a pool of reusable Event objects to optimize memory allocation.
var EventPool = sync.Pool{
	New: func() interface{} {
		return &models.Event{}
	},
}
