package middlewares_test

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"github.com/mjmoshiri/log-lyfe/gol/internal/server"
	"github.com/mjmoshiri/log-lyfe/gol/internal/server/middlewares"
	"net/http"
	"testing"
)

func TestAuthMiddleware(t *testing.T) {
	const correctKey = "correct-auth-key"
	config := &models.ServerConfig{
		Port:      ":8081",
		SecretKey: correctKey,
	}
	app := server.New(config)
	app.Use(middlewares.AuthMiddleware(config.SecretKey))
	app.AddRoute("GET", "/protected", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	go app.ListenAndServe()
	<-app.Ready

	// Test without Authorization header
	resp, err := http.Get("http://localhost:8081/protected")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, resp.StatusCode)
	}

	// Test with incorrect Authorization value
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8081/protected", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", "wrong-key")
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, resp.StatusCode)
	}

	// Test with correct Authorization value
	req, err = http.NewRequest("GET", "http://localhost:8081/protected", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", correctKey)
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
