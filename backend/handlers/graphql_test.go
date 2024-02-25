package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/gwkline/full-stack-skeleton/backend/mocks"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/util"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGraphQL(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		query          string
		setupMock      func(mock *sqlmock.Sqlmock, auth *mocks.IAuth)
		expectedBody   string
		expectedStatus int
	}{
		{
			name:  "heartbeat query success",
			query: `{"query": "{ heartbeat }"}`,
			setupMock: func(db *sqlmock.Sqlmock, auth *mocks.IAuth) {
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:  "unauthorized due to missing token",
			query: `{"query": "{ viewer { apples { id } } }"}`,
			setupMock: func(db *sqlmock.Sqlmock, auth *mocks.IAuth) {
				auth.On("CurrentUser", mock.Anything).Return(nil, fmt.Errorf("failed retrieving user from context"))
			},
			expectedBody:   `{"errors":[{"message":"user not found in context","path":["viewer"]}],"data":null}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:  "authorized access with admin role",
			query: `{"query": "{ viewer { apples { id } } }"}`,
			setupMock: func(db *sqlmock.Sqlmock, auth *mocks.IAuth) {
				auth.On("CurrentUser", mock.Anything).Return(&types.User{Role: "Admin", Email: "nice@gmail.com"})
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := util.MockDB()
			auth := mocks.NewIAuth(t)
			h := Handler{
				Repository: db,
				Service: &types.Services{
					Auth: auth,
				},
			}

			tt.setupMock(&mock, auth)

			r := gin.Default()
			r.POST("/graphql", h.Graphql())

			req, _ := http.NewRequest("POST", "/graphql", bytes.NewBufferString(tt.query))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			fmt.Println(w.Body)

			assert.Equal(t, tt.expectedStatus, w.Code, "GraphQL endpoint returned wrong status code")
			if tt.expectedBody != "" {
				assert.Equal(t, tt.expectedBody, w.Body.String(), "GraphQL endpoint returned wrong body")
			}

		})
	}
}
