package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPlayground(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := Handler{}

	r := gin.Default()
	r.GET("/", h.Playground())

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Playground endpoint returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}
}
