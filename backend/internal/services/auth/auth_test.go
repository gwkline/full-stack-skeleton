package auth

import (
	"testing"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
)

type JWT struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
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
