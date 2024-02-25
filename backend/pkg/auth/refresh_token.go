package auth

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (s *Service) RefreshToken(ctx context.Context, data types.JWT) (*types.JWT, error) {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("RefreshToken").End()

	claims, err := s.ValidateToken(ctx, data.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	claims2, err := s.ValidateToken(ctx, data.AccessToken)
	if err != nil && !strings.Contains(err.Error(), "token is expired") {
		return nil, fmt.Errorf("invalid access token: %w", err)
	}

	if claims.Email != claims2.Email {
		return nil, fmt.Errorf("invalid refresh token")
	}

	if time.Unix(claims.ExpiresAt.Unix(), 0).Before(time.Now()) {
		return nil, fmt.Errorf("refresh token expired")
	}

	newAccessToken, err := s.GenerateToken(ctx, claims.Email, AccessTokenDuration)
	if err != nil {
		return nil, fmt.Errorf("failed generating Access Token: %w", err)
	}

	return &types.JWT{
		RefreshToken: data.RefreshToken,
		AccessToken:  newAccessToken,
	}, nil
}
