package handlers

import "github.com/mjmoshiri/log-lyfe/gol/storage"

// AppHandler is the custom handler type for the application
// that has access to the database interface.
type AppHandler struct {
	DB storage.DB
}
