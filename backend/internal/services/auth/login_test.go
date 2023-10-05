package auth

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/gwkline/full-stack-infra/backend/internal/graph/model"
	"github.com/gwkline/full-stack-infra/backend/internal/helpers"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)

	hashedPW, _ := hashPassword("password123")
	testUser := model.NewUser{
		Email:    "email@gmail.com",
		Password: hashedPW,
	}

	db, mock := helpers.MockDB()
	db.InsertUser(testUser)

	var body Login
	body.Email = "email@gmail.com"
	body.Password = "password123"
	body.OTP = ""

	helpers.MockJsonPost(ctx, body)

	findUserSQL := regexp.QuoteMeta(`SELECT id, email, passwordHash, otpSecret, phone, 
    EXTRACT(EPOCH FROM createdAt)::bigint, 
    EXTRACT(EPOCH FROM updatedAt)::bigint 
    FROM users WHERE email = $1`)

	mock.ExpectQuery(findUserSQL).
		WithArgs("email@gmail.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "passwordHash", "otpSecret", "phone", "createdAt", "updatedAt"}).
			AddRow(1, "email@gmail.com", hashedPW, nil, "phone", 1633429591, 1633429591))

	LoginHandler(ctx, db)
	assert.EqualValues(t, http.StatusOK, w.Code)
}

func TestLoginHandler_JSONBindingError(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)

	// Mock with invalid JSON data
	helpers.MockJsonPost(ctx, "{this_is_invalid_json}")

	LoginHandler(ctx, nil)
	assert.EqualValues(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Bad request")
}

func TestLoginHandler_EmptyEmailPassword(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)

	body := Login{Email: "", Password: ""}
	helpers.MockJsonPost(ctx, body)

	LoginHandler(ctx, nil)
	assert.EqualValues(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Bad request")
}

func TestLoginHandler_UserNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)

	db, mock := helpers.MockDB()
	mock.ExpectQuery(`SELECT`).WillReturnError(sql.ErrNoRows)

	body := Login{Email: "email@gmail.com", Password: "password123"}
	helpers.MockJsonPost(ctx, body)

	LoginHandler(ctx, db)
	assert.EqualValues(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "User not found")
}
