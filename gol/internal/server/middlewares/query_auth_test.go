package middlewares

import (
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"github.com/mjmoshiri/log-lyfe/gol/internal/server"
	"net/http"
	"testing"
)

func TestQueryAuthMiddleware(t *testing.T) {
	config := &models.ServerConfig{
		Port:     ":8081",
		QueryKey: "correct-key",
	}
	app := server.New(config)
	app.Use(QueryAuthMiddleware(config.QueryKey))
	app.AddRoute("GET", "/Query", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	app.AddRoute("GET", "/home", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Home"))
	})
	go app.ListenAndServe()
	<-app.Ready

	tests := []struct {
		method         string
		path           string
		headerValue    string
		expectedStatus int
	}{
		{http.MethodPost, "/Query", "", http.StatusNotFound},
		{http.MethodGet, "/WrongQuery", "", http.StatusNotFound},
		{http.MethodGet, "/Query", "wrong-key", http.StatusUnauthorized},
		{http.MethodGet, "/Query", "correct-key", http.StatusOK},
		{http.MethodGet, "/home", "", http.StatusOK},
	}

	for _, tc := range tests {
		req, err := http.NewRequest(tc.method, "http://localhost:8081"+tc.path, nil)
		if err != nil {
			t.Fatal(err)
		}

		if tc.headerValue != "" {
			req.Header.Set("Query-Key", tc.headerValue)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != tc.expectedStatus {
			t.Errorf("For %s %s with header %s, expected status code %d, got %d",
				tc.method, tc.path, tc.headerValue, tc.expectedStatus, resp.StatusCode)
		}
	}
}
