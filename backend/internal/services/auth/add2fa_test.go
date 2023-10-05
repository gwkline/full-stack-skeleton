package auth

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gwkline/full-stack-infra/backend/internal/graph/model"
	"github.com/gwkline/full-stack-infra/backend/internal/services/helpers"
	"github.com/stretchr/testify/assert"
)

func Test2FAHandler(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)
	db, mock := helpers.MockDB()

	var body model.NewUser
	body.Email = "new_email@gmail.com"
	body.Password = "new_password"

	hashedPW, _ := hashPassword("password123")

	helpers.MockJsonPost(ctx, body)

	findUserSQL := regexp.QuoteMeta(`SELECT id, email, passwordHash, otpSecret, phone,
    EXTRACT(EPOCH FROM createdAt)::bigint,
    EXTRACT(EPOCH FROM updatedAt)::bigint
    FROM users WHERE email = $1`)
	mock.ExpectQuery(findUserSQL).
		WithArgs("new_email@gmail.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "passwordHash", "otpSecret", "phone", "createdAt", "updatedAt"}).
			AddRow(1, "email@gmail.com", hashedPW, nil, "phone", 1633429591, 1633429591))

	findUserSQL2 := regexp.QuoteMeta(`SELECT id, email, passwordHash, otpSecret, phone,
			EXTRACT(EPOCH FROM createdAt)::bigint,
			EXTRACT(EPOCH FROM updatedAt)::bigint
			FROM users WHERE id = $1`)
	mock.ExpectQuery(findUserSQL2).
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "passwordHash", "otpSecret", "phone", "createdAt", "updatedAt"}).
			AddRow(1, "email@gmail.com", hashedPW, nil, "phone", 1633429591, 1633429591))

	insertUserSQL := regexp.QuoteMeta(`UPDATE users SET email = $1, phone = $2, otpSecret = $3, updatedAt = NOW() WHERE id = $4 RETURNING createdAt;`)

	mock.ExpectPrepare(insertUserSQL).ExpectQuery().
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"createdAt"}).
			AddRow(time.Now()))

	Add2FA(ctx, db)
	assert.EqualValues(t, http.StatusOK, w.Code)
}

func TestAdd2FA_JSONBindingError(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)

	helpers.MockJsonPost(ctx, "{this_is_invalid_json}")

	Add2FA(ctx, nil)
	assert.EqualValues(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Bad request")
}

func TestAdd2FA_EmptyEmailPassword(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)

	body := Login{Email: "", Password: ""}
	helpers.MockJsonPost(ctx, body)

	Add2FA(ctx, nil)
	assert.EqualValues(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Bad request")
}

func TestAdd2FA_UserNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)

	db, mock := helpers.MockDB()
	mock.ExpectQuery(`SELECT`).WillReturnError(sql.ErrNoRows)

	body := Login{Email: "email@gmail.com", Password: "password123"}
	helpers.MockJsonPost(ctx, body)

	Add2FA(ctx, db)
	assert.EqualValues(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "User not found")
}

func Test2FAHandler_TOTPError(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)
	db, mock := helpers.MockDB()

	var body model.NewUser
	body.Email = "new_email@gmail.com"
	body.Password = "new_password"

	hashedPW, _ := hashPassword("password123")

	helpers.MockJsonPost(ctx, body)

	findUserSQL := regexp.QuoteMeta(`SELECT id, email, passwordHash, otpSecret, phone,
    EXTRACT(EPOCH FROM createdAt)::bigint,
    EXTRACT(EPOCH FROM updatedAt)::bigint
    FROM users WHERE email = $1`)
	mock.ExpectQuery(findUserSQL).
		WithArgs("new_email@gmail.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "passwordHash", "otpSecret", "phone", "createdAt", "updatedAt"}).
			AddRow(1, "e", hashedPW, nil, "phone", 1633429591, 1633429591))

	findUserSQL2 := regexp.QuoteMeta(`SELECT id, email, passwordHash, otpSecret, phone,
			EXTRACT(EPOCH FROM createdAt)::bigint,
			EXTRACT(EPOCH FROM updatedAt)::bigint
			FROM users WHERE id = $1`)
	mock.ExpectQuery(findUserSQL2).
		WithArgs("1").
		WillReturnError(sql.ErrNoRows)

	Add2FA(ctx, db)
	assert.EqualValues(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Failed to set key on user")

}

func Test2FAHandler_AlreadyEnabled(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := helpers.TestGinContext(w)
	db, mock := helpers.MockDB()

	var body model.NewUser
	body.Email = "new_email@gmail.com"
	body.Password = "new_password"

	hashedPW, _ := hashPassword("password123")

	helpers.MockJsonPost(ctx, body)

	findUserSQL := regexp.QuoteMeta(`SELECT id, email, passwordHash, otpSecret, phone,
    EXTRACT(EPOCH FROM createdAt)::bigint,
    EXTRACT(EPOCH FROM updatedAt)::bigint
    FROM users WHERE email = $1`)
	mock.ExpectQuery(findUserSQL).
		WithArgs("new_email@gmail.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "passwordHash", "otpSecret", "phone", "createdAt", "updatedAt"}).
			AddRow(1, "email@gmail.com", hashedPW, "nil", "phone", 1633429591, 1633429591))

	Add2FA(ctx, db)
	assert.EqualValues(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "2FA already enabled")
}
