package auth

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/gwkline/full-stack-infra/backend/internal/graph/model"
	"github.com/gwkline/full-stack-infra/backend/internal/helpers"
	"github.com/pquerna/otp/totp"
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

func TestGenerateToken(t *testing.T) {
	email := "test@example.com"
	token, err := generateToken(email, AccessTokenDuration)
	assert.Nil(t, err)
	claims, err := validateToken(token)
	assert.Nil(t, err)
	assert.Equal(t, email, claims.Email)
}

func TestValidateToken(t *testing.T) {
	// Valid token
	email := "test@example.com"
	token, _ := generateToken(email, AccessTokenDuration)
	claims, err := validateToken(token)
	assert.Nil(t, err)
	assert.Equal(t, email, claims.Email)

	// Invalid token
	invalidToken := "invalid.token.here"
	_, err = validateToken(invalidToken)
	assert.NotNil(t, err)
}

func TestGenerateTOTPKey(t *testing.T) {
	email := "test@example.com"
	key, err := generateTOTPKey(email)
	assert.Nil(t, err)
	assert.NotNil(t, key.URL())
}

func TestValidOtpCode(t *testing.T) {
	email := "test@example.com"
	key, _ := generateTOTPKey(email)
	validToken, _ := totp.GenerateCode(key.Secret(), time.Now())
	assert.True(t, validOtpCode(key.Secret(), validToken))

	// Invalid OTP
	assert.False(t, validOtpCode(key.Secret(), "123456"))
}

func TestHashPassword(t *testing.T) {
	password := "securepassword"
	hash, err := hashPassword(password)
	assert.Nil(t, err)
	assert.True(t, validPassword(password, hash))
}

func TestValidPassword(t *testing.T) {
	password := "securepassword"
	hash, _ := hashPassword(password)
	assert.True(t, validPassword(password, hash))
	assert.False(t, validPassword("wrongpassword", hash))
}
