package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetEnv(t *testing.T) {
	key := "TEST_ENV_VAR"
	fallback := "default"

	// Case 1: Env var is set
	os.Setenv(key, "test_value")
	got := getEnv(key, fallback)
	if got != "test_value" {
		t.Errorf("Expected test_value, got %s", got)
	}

	// Case 2: Env var is not set
	os.Unsetenv(key)
	got = getEnv(key, fallback)
	if got != fallback {
		t.Errorf("Expected %s, got %s", fallback, got)
	}
}

func TestInitSentry(t *testing.T) {
	// Init Sentry
	err := InitSentry("https://sentry.io/")
	if err != true {
		t.Errorf("Expected true, got false")
	}
}

func TestMainExecution(t *testing.T) {
	go main()

	req, _ := http.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()
	r := gin.Default()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

}

func TestLoadConfig(t *testing.T) {
	// Mock environment variables
	os.Setenv("PORT", "8081")
	os.Setenv("ENV", "production")
	os.Setenv("POSTGRES_USER", "testUser")
	os.Setenv("POSTGRES_PASSWORD", "testPass")
	os.Setenv("POSTGRES_DB", "testDB")
	os.Setenv("SENTRY_BE_DSN", "https://sentry.io/")

	expectedConfig := Config{
		Port:             "8081",
		Env:              "production",
		DatabaseUser:     "testUser",
		DatabasePassword: "testPass",
		DatabaseName:     "testDB",
		SentryDSN:        "https://sentry.io/",
	}

	config := LoadConfig()
	if config != expectedConfig {
		t.Errorf("Expected %+v, got %+v", expectedConfig, config)
	}
}

func TestSetupRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		env             string
		expectedStatus  int
		expectedHeaders http.Header
	}{
		{"production", http.StatusOK, http.Header{"Allow": []string{"OPTIONS, POST, GET, PUT, DELETE"}, "Access-Control-Allow-Origin": []string{"http://localhost:8080"}, "Access-Control-Allow-Headers": []string{"Origin, Content-Type, Accept, Authorization"}}},
		{"development", http.StatusOK, http.Header{"Allow": []string{"POST, GET, OPTIONS, PUT, DELETE"}}},
		{"", http.StatusOK, http.Header{"Allow": []string{"POST, GET, OPTIONS, PUT, DELETE"}}},
	}

	for _, tt := range tests {
		r := SetupRouter(tt.env)

		req, _ := http.NewRequest("OPTIONS", "/graphql", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != tt.expectedStatus {
			t.Errorf("For env %s: expected status %d, got %d", tt.env, tt.expectedStatus, w.Code)
		}
	}
}

func TestGraphQLHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/graphql", graphqlHandler())

	query := `{
		"query": "{ todos { id } }"
	}`

	req, _ := http.NewRequest("POST", "/graphql", bytes.NewBufferString(query))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("GraphQL endpoint returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}
}

func TestPlaygroundHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/", playgroundHandler())

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Playground endpoint returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}
}
