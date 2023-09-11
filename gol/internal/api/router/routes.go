package router

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/api"
	"github.com/mjmoshiri/log-lyfe/gol/internal/api/middlewares"
	"github.com/mjmoshiri/log-lyfe/gol/internal/handlers"
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"github.com/mjmoshiri/log-lyfe/gol/storage"
)

// SetupRoutes sets up the routes for the application.
// In production, OpenAPI could be used to generate the routes.
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
