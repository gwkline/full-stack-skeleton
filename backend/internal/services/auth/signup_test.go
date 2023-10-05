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

func TestSignupHandler(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)
	db, mock := helpers.MockDB()

	var body model.NewUser
	body.Email = "new_email@gmail.com"
	body.Password = "new_password"

	helpers.MockJsonPost(ctx, body)

	findUserSQL := regexp.QuoteMeta(`SELECT id, email, passwordHash, otpSecret, phone,
    EXTRACT(EPOCH FROM createdAt)::bigint,
    EXTRACT(EPOCH FROM updatedAt)::bigint
    FROM users WHERE email = $1`)
	mock.ExpectQuery(findUserSQL).
		WithArgs("new_email@gmail.com").
		WillReturnError(sql.ErrNoRows) // simulating that user doesn't exist

	insertUserSQL := regexp.QuoteMeta(`INSERT INTO users (email, passwordHash, otpSecret, phone, createdAt, updatedAt) VALUES($1, $2, $3, $4, TIMESTAMP 'epoch' + $5 * INTERVAL '1 second', TIMESTAMP 'epoch' + $6 * INTERVAL '1 second') RETURNING id;`)

	mock.ExpectPrepare(insertUserSQL).ExpectQuery().
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	SignupHandler(ctx, db)
	assert.EqualValues(t, http.StatusOK, w.Code)
}

func TestSignupHandler_JSONBindingError(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)

	helpers.MockJsonPost(ctx, "{this_is_invalid_json}")

	SignupHandler(ctx, nil)
	assert.EqualValues(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Bad request")
}

func TestSignupHandler_EmptyEmailPassword(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)

	body := model.NewUser{Email: "", Password: ""}
	helpers.MockJsonPost(ctx, body)

	SignupHandler(ctx, nil)
	assert.EqualValues(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Bad request")
}

func TestSignupHandler_DatabaseInsertionError(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)

	db, mock := helpers.MockDB()
	findUserSQL := regexp.QuoteMeta(`SELECT id, email, passwordHash, otpSecret, phone,
    EXTRACT(EPOCH FROM createdAt)::bigint,
    EXTRACT(EPOCH FROM updatedAt)::bigint
    FROM users WHERE email = $1`)
	mock.ExpectQuery(findUserSQL).
		WithArgs("email@gmail.com").
		WillReturnError(sql.ErrNoRows)

	mock.ExpectExec(`INSERT`).WillReturnError(sql.ErrConnDone)

	body := model.NewUser{Email: "email@gmail.com", Password: "password123"}
	helpers.MockJsonPost(ctx, body)

	SignupHandler(ctx, db)
	assert.EqualValues(t, http.StatusInternalServerError, w.Code)
}

func TestSignupHandler_AlreadExists(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)
	db, mock := helpers.MockDB()

	var body model.NewUser
	body.Email = "new_email@gmail.com"
	body.Password = "new_password"

	helpers.MockJsonPost(ctx, body)

	findUserSQL := regexp.QuoteMeta(`SELECT id, email, passwordHash, otpSecret, phone,
    EXTRACT(EPOCH FROM createdAt)::bigint,
    EXTRACT(EPOCH FROM updatedAt)::bigint
    FROM users WHERE email = $1`)
	mock.ExpectQuery(findUserSQL).
		WithArgs("new_email@gmail.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "passwordHash", "otpSecret", "phone", "createdAt", "updatedAt"}).
			AddRow(1, "e", "hashedPW", nil, "phone", 1633429591, 1633429591))

	SignupHandler(ctx, db)
	assert.EqualValues(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "User already exists")
}

func TestSignupHandler_BadPW(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)
	db, mock := helpers.MockDB()

	var body model.NewUser
	body.Email = "new_email@gmail.com"
	body.Password = "234567098765432234567234567098765432234567234567098765432234567234567098765432234567234567098765432234567234567098765432234567234"

	helpers.MockJsonPost(ctx, body)

	findUserSQL := regexp.QuoteMeta(`SELECT id, email, passwordHash, otpSecret, phone,
    EXTRACT(EPOCH FROM createdAt)::bigint,
    EXTRACT(EPOCH FROM updatedAt)::bigint
    FROM users WHERE email = $1`)
	mock.ExpectQuery(findUserSQL).
		WithArgs("new_email@gmail.com").
		WillReturnError(sql.ErrNoRows) // simulating that user doesn't exist

	insertUserSQL := regexp.QuoteMeta(`INSERT INTO users (email, passwordHash, otpSecret, phone, createdAt, updatedAt) VALUES($1, $2, $3, $4, TIMESTAMP 'epoch' + $5 * INTERVAL '1 second', TIMESTAMP 'epoch' + $6 * INTERVAL '1 second') RETURNING id;`)

	mock.ExpectPrepare(insertUserSQL).ExpectQuery().
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	SignupHandler(ctx, db)
	assert.EqualValues(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Error hashing password")
}
