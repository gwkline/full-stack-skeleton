package auth

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (s *Service) ValidateToken(ctx context.Context, tokenStr string) (*types.Claims, error) {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("ValidateToken").End()

	claims := &types.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	return claims, nil
}
