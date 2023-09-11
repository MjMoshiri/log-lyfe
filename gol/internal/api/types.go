package api

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/api/middlewares"
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"net/http"
)

// App is the main application struct.
type App struct {
	middlewares []middlewares.MiddlewareFunc
	config      *models.ServerConfig
	router      map[string]http.HandlerFunc
	Ready       chan struct{}
}
