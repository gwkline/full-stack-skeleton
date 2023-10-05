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
