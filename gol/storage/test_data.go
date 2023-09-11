package storage

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"time"
)

// Events is a slice of models.Event for testing the storage package.
var Events = []models.Event{
	{
		Action:     "create",
		Actor:      "user1",
		Timestamp:  time.Now(),
		EventID:    "e1",
		Version:    "1.0",
		ActionMeta: map[string]string{"key": "value"},
		ActorMeta:  map[string]string{"role": "admin"},
	},
	{
		Action:     "update",
		Actor:      "user2",
		Timestamp:  time.Now(),
		EventID:    "e2",
		Version:    "1.0",
		ActionMeta: map[string]string{},
		ActorMeta:  map[string]string{"role": "member"},
	},
	{
		Action:     "delete",
		Actor:      "user3",
		Timestamp:  time.Now(),
		EventID:    "e3",
		Version:    "1.0",
		ActionMeta: map[string]string{"reason": "obsolete"},
		ActorMeta:  map[string]string{"role": "guest"},
	},
	{
		Action:     "create",
		Actor:      "user4",
		Timestamp:  time.Now(),
		EventID:    "e4",
		Version:    "1.0",
		ActionMeta: map[string]string{"key": "otherValue"},
		ActorMeta:  map[string]string{"role": "admin"},
	},
	{
		Action:     "update",
		Actor:      "user5",
		Timestamp:  time.Now(),
		EventID:    "e5",
		Version:    "1.0",
		ActionMeta: map[string]string{"property": "color"},
		ActorMeta:  map[string]string{"role": "member"},
	},
	{
		Action:     "create",
		Actor:      "user6",
		Timestamp:  time.Now().Add(time.Hour * -1),
		EventID:    "e6",
		Version:    "1.0",
		ActionMeta: map[string]string{"key": "someValue"},
		ActorMeta:  map[string]string{"role": "admin"},
	},
	{
		Action:     "create",
		Actor:      "user7",
		Timestamp:  time.Now().Add(time.Hour * -1),
		EventID:    "e7",
		Version:    "1.0",
		ActionMeta: map[string]string{},
		ActorMeta:  map[string]string{"role": "member"},
	},
	{
		Action:     "delete",
		Actor:      "user8",
		Timestamp:  time.Now().Add(time.Hour * -1),
		EventID:    "e8",
		Version:    "1.0",
		ActionMeta: map[string]string{"reason": "userRequest"},
		ActorMeta:  map[string]string{"role": "guest"},
	},
	{
		Action:     "update",
		Actor:      "user9",
		Timestamp:  time.Now().Add(time.Hour * -1),
		EventID:    "e9",
		Version:    "1.0",
		ActionMeta: map[string]string{"property": "size"},
		ActorMeta:  map[string]string{"role": "admin"},
	},
	{
		Action:     "create",
		Actor:      "user10",
		Timestamp:  time.Now().Add(time.Hour * -1),
		EventID:    "e10",
		Version:    "1.0",
		ActionMeta: map[string]string{"key": "lastValue"},
		ActorMeta:  map[string]string{"role": "guest"},
	},
}
