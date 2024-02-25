package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateToken(t *testing.T) {
	s := Service{}
	email := "test@example.com"
	token, _ := s.GenerateToken(context.Background(), email, AccessTokenDuration)
	claims, err := s.ValidateToken(context.Background(), token)
	assert.Nil(t, err)
	assert.Equal(t, email, claims.Email)

	invalidToken := "invalid.token.here"
	_, err = s.ValidateToken(context.Background(), invalidToken)
	assert.NotNil(t, err)
}
