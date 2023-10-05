package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gwkline/full-stack-infra/backend/internal/helpers"
	"github.com/stretchr/testify/assert"
)

func TestRefreshTokenHandler(t *testing.T) {

	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)
	db, _ := helpers.MockDB()

	var body JWT
	body.AccessToken, _ = generateToken("email@gmail.com", time.Second*1)
	body.RefreshToken, _ = generateToken("email@gmail.com", time.Hour*1)

	helpers.MockJsonPost(ctx, body)

	RefreshTokenHandler(ctx, db)
	assert.EqualValues(t, http.StatusOK, w.Code)
}

func TestRefreshTokenHandler_JSONBindingError(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)

	helpers.MockJsonPost(ctx, "{this_is_invalid_json}")

	RefreshTokenHandler(ctx, nil)
	assert.EqualValues(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Bad request")
}