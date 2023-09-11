package api

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/api/middlewares"
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"net/http"
)

// App represents the main application structure, encapsulating middleware, configuration, routing, and server readiness.
type App struct {
	middlewares []middlewares.MiddlewareFunc
	config      *models.ServerConfig
	router      map[string]http.HandlerFunc
	Ready       chan struct{}
}
