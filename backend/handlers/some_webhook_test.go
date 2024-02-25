package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gwkline/full-stack-skeleton/backend/mocks"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/stretchr/testify/require"
)

func TestSomeWebhookHandler(t *testing.T) {

	tests := []struct {
		name           string
		messageURL     string
		expectedStatus int
		mockFuncs      func(workers *mocks.IWorker)
	}{
		{
			name:           "url not found",
			messageURL:     "",
			expectedStatus: http.StatusOK,
			mockFuncs:      func(workers *mocks.IWorker) {},
		},
		{
			name:           "happy path",
			messageURL:     "http://example.com",
			expectedStatus: http.StatusOK,
			mockFuncs:      func(workers *mocks.IWorker) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()

			workers := mocks.NewIWorker(t)

			tt.mockFuncs(workers)

			w := httptest.NewRecorder()
			form := url.Values{}
			form.Add("message-url", tt.messageURL)
			reqBody := strings.NewReader(form.Encode())

			services := types.Services{
				Workers: workers,
			}
			h := Handler{
				Service: &services,
			}

			var wg sync.WaitGroup
			wg.Add(1)

			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.POST("/someWebhookEndpoint", h.SomeWebhookHandler(&wg))

			req, _ := http.NewRequest("POST", "/someWebhookEndpoint", reqBody)
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			r.ServeHTTP(w, req)

			wg.Wait()

			require.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
