package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/newrelic/go-agent/v3/newrelic"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) GenerateToken(ctx context.Context, email string, duration time.Duration) (string, error) {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("GenerateToken").End()

	claims := &jwt.MapClaims{
		"exp":   jwt.NewNumericDate(time.Now().Add(duration)),
		"Email": email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (s *Service) generatePasswordHash(ctx context.Context, password string) (string, error) {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("generatePasswordHash").End()

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 11)
	return string(bytes), err
}

func (s *Service) validatePassword(ctx context.Context, password, hash string) bool {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("validatePassword").End()

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type key int

const (
	UserKey key = iota
)

func (s *Service) CurrentUser(ctx context.Context) (*types.User, error) {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("CurrentUser").End()

	raw, ok := ctx.Value(UserKey).(*types.User)
	if !ok {
		return nil, fmt.Errorf("failed retrieving user from context")
	}

	return raw, nil
}

func (s *Service) SetCurrentUser(ctx context.Context, user *types.User) context.Context {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("SetCurrentUser").End()

	return context.WithValue(ctx, UserKey, user)
}
