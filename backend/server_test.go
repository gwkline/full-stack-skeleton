package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

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
		log.Println(w.Body.String())
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
