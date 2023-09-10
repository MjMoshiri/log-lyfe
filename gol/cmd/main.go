package main

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"github.com/mjmoshiri/log-lyfe/gol/internal/server"
	"github.com/mjmoshiri/log-lyfe/gol/internal/server/middlewares"
	"net/http"
)

func main() {
	config := &models.ServerConfig{
		Port:      ":8080",
		SecretKey: "test",
		QueryKey:  "test",
	}
	app := server.New(config)
	app.Use(middlewares.QueryAuthMiddleware(config.QueryKey))
	app.Use(middlewares.AuthMiddleware(config.SecretKey))
	app.Use(middlewares.RecoverMiddleware)
	app.AddRoute("GET", "/home", func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	})
	app.ListenAndServe()
}
