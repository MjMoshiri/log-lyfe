package server

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"net/http"
)

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// App is the main application struct.
type App struct {
	middlewares []MiddlewareFunc
	config      *models.ServerConfig
	router      map[string]http.HandlerFunc
	Ready       chan struct{}
}
