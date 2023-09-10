package utils

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"sync"
)

// EventPool is a sync.Pool that contains events.
var EventPool = sync.Pool{
	New: func() interface{} {
		return &models.Event{}
	},
}
