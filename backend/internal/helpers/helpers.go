package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/gwkline/full-stack-infra/backend/internal/database"
)

func ValidateEmail(email string) error {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return err
	}

	err = checkmail.ValidateHost(email)
	if err != nil {
		return err
	}

	return nil
}

func TestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx
}

func MockJsonGet(c *gin.Context, params gin.Params, u url.Values) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")

	// set path params
	c.Params = params

	// set query params
	c.Request.URL.RawQuery = u.Encode()
}

func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func MockJsonPut(c *gin.Context, content interface{}, params gin.Params) {
	c.Request.Method = "PUT"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = params

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func MockJsonDelete(c *gin.Context, params gin.Params) {
	c.Request.Method = "DELETE"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = params
}

func MockDB() (*database.Database, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Failed to create sqlmock: %s", err)
	}

	// Create your Database instance with the mock connection
	mockDatabase := database.NewDatabase(db)
	return mockDatabase, mock
}
