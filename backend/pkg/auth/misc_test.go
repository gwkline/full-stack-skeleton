package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	s := Service{}
	email := "test@example.com"
	token, err := s.GenerateToken(context.Background(), email, AccessTokenDuration)
	assert.Nil(t, err)
	claims, err := s.ValidateToken(context.Background(), token)
	assert.Nil(t, err)
	assert.Equal(t, email, claims.Email)
}

func TestHashPassword(t *testing.T) {
	s := Service{}
	password := "securepassword"
	hash, err := s.generatePasswordHash(context.Background(), password)
	assert.Nil(t, err)
	assert.True(t, s.validatePassword(context.Background(), password, hash))
}

func TestValidPassword(t *testing.T) {
	s := Service{}
	password := "securepassword"
	hash, _ := s.generatePasswordHash(context.Background(), password)
	assert.True(t, s.validatePassword(context.Background(), password, hash))
	assert.False(t, s.validatePassword(context.Background(), "wrongpassword", hash))
}
