// Package api provides core functionalities for setting up and managing the application's HTTP server.
// It includes utilities for route management, middleware application, and server configuration.
package api

import (
	"context"
	"errors"
	"github.com/mjmoshiri/log-lyfe/gol/internal/api/middlewares"
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// New initializes a new App instance with the provided server configuration.
func New(config *models.ServerConfig) App {
	return App{
		config: config,
		router: map[string]http.HandlerFunc{},
		Ready:  make(chan struct{}),
	}
}

// Use appends a middleware to the App's middleware stack.
func (a *App) Use(m middlewares.MiddlewareFunc) {
	a.middlewares = append(a.middlewares, m)
}

// AddRoute associates a given HTTP method and route with a handler in the App's router.
func (a *App) AddRoute(method string, route string, handler http.HandlerFunc) {
	key := method + ":" + route
	a.router[key] = handler
}

// ServeHTTP processes incoming HTTP requests, applying middlewares and routing to the appropriate handler.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + ":" + r.URL.Path
	if handler, exists := a.router[key]; exists {
		finalHandler := handler
		for i := len(a.middlewares) - 1; i >= 0; i-- {
			finalHandler = a.middlewares[i](finalHandler)
		}
		finalHandler(w, r)
	} else {
		http.NotFound(w, r)
	}
}

// ListenAndServe starts the App's HTTP server, listening for incoming requests and handling graceful shutdown.
func (a *App) ListenAndServe() {
	server := &http.Server{
		Addr:    "0.0.0.0:" + a.config.Port,
		Handler: a,
	}

	// Start the server asynchronously
	go func() {
		log.Printf("Server started on %s", a.config.Port)
		go func() {
			time.Sleep(100 * time.Millisecond)
			close(a.Ready)
		}()
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on %s: %v\n", a.config.Port, err)
		}
	}()

	// Await OS signals for termination
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan
	log.Println("Shutting down server...")

	// Context for graceful server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server gracefully stopped")
}
