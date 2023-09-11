package middlewares_test

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/api"
	"github.com/mjmoshiri/log-lyfe/gol/internal/api/middlewares"
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"net/http"
	"testing"
)

func TestRecoverMiddleware(t *testing.T) {
	config := &models.ServerConfig{
		Port: "8083",
	}
	app := api.New(config)
	app.Use(middlewares.RecoverMiddleware)
	app.AddRoute("GET", "/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	})
	app.AddRoute("GET", "/home", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	go app.ListenAndServe()
	<-app.Ready
	resp, err := http.Get("http://localhost:8083/panic")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, resp.StatusCode)
	}

	// Test no panic
	resp, err = http.Get("http://localhost:8083/home")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
