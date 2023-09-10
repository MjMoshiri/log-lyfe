package server

import (
	"context"
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// New creates a new App instance.
func New(config *models.ServerConfig) *App {
	return &App{
		config: config,
		router: map[string]http.HandlerFunc{},
		Ready:  make(chan struct{}),
	}
}

// Use adds a middleware to the App instance.
func (a *App) Use(m MiddlewareFunc) {
	a.middlewares = append(a.middlewares, m)
}

// AddRoute adds a route to the App instance.
func (a *App) AddRoute(method string, route string, handler http.HandlerFunc) {
	key := method + ":" + route
	a.router[key] = handler
}

// ServeHTTP is the main entry point for the App instance.
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

// ListenAndServe starts the App instance.
func (a *App) ListenAndServe() {
	server := &http.Server{
		Addr:    a.config.Port,
		Handler: a,
	}

	// Start the server
	go func() {
		log.Printf("Server started on %s", a.config.Port)
		close(a.Ready)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", a.config.Port, err)
		}
	}()

	// Create a channel to listen for OS signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// Block until we receive a signal
	<-stopChan
	log.Println("Shutting down server...")

	// Create a context with timeout for the server to shut down
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server gracefully stopped")
}
