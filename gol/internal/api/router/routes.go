package router

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/api"
	"github.com/mjmoshiri/log-lyfe/gol/internal/api/middlewares"
	"github.com/mjmoshiri/log-lyfe/gol/internal/handlers"
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"github.com/mjmoshiri/log-lyfe/gol/storage"
)

// SetupRoutes initializes the application routes and applies necessary middlewares.
// Consider using OpenAPI for route generation in production environments.
func SetupRoutes(app *api.App, db storage.DB, serverConfig *models.ServerConfig) {
	app.Use(middlewares.QueryAuthMiddleware(serverConfig.QueryKey))
	app.Use(middlewares.AuthMiddleware(serverConfig.SecretKey))
	app.Use(middlewares.RecoverMiddleware)
	h := &handlers.AppHandler{
		DB: db,
	}
	app.AddRoute("GET", "/info", h.InfoHandler)
	app.AddRoute("GET", "/ok", h.HealthCheckHandler)
	app.AddRoute("GET", "/query", h.HandleQueryRequest)
	app.AddRoute("POST", "/insert", h.HandleEventRequest)
}
