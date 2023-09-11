package handlers

import "github.com/mjmoshiri/log-lyfe/gol/storage"

// AppHandler represents the application's custom handler.
// It provides access to the underlying database interface.
type AppHandler struct {
	DB storage.DB
}
