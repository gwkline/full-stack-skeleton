package auth

import (
	"context"
	"fmt"

	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (s *Service) Login(ctx context.Context, login types.LoginInput) (*types.JWT, error) {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("Login").End()

	user, err := s.Repository.User.FindBy([]types.Filter{{Key: "email", Value: login.Email}})
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	if !s.validatePassword(ctx, login.Password, user.PasswordHash) {
		return nil, fmt.Errorf("invalid password")
	}

	accessToken, err := s.GenerateToken(ctx, login.Email, AccessTokenDuration)
	if err != nil {
		return nil, fmt.Errorf("failed generating Access Token: %w", err)
	}

	refreshToken, err := s.GenerateToken(ctx, login.Email, RefreshTokenDuration)
	if err != nil {
		return nil, fmt.Errorf("failed generating Refresh Token: %w", err)
	}

	return &types.JWT{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}
