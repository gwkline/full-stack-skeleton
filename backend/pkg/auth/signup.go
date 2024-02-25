package auth

import (
	"context"
	"fmt"
	"strings"

	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (s *Service) Signup(ctx context.Context, newUser types.User, password string) (*types.JWT, error) {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("Signup").End()

	if newUser.Email == "" || password == "" {
		return nil, fmt.Errorf("invalid login")
	}

	if !strings.Contains(newUser.Email, "@gmail.com") {
		return nil, fmt.Errorf("please use your work email to sign in")
	}

	_, err := s.Repository.User.FindBy([]types.Filter{{Key: "email", Value: newUser.Email}})
	if err == nil {
		return nil, fmt.Errorf("user already exists")
	}

	var user types.User
	user.Email = newUser.Email
	user.PasswordHash, err = s.generatePasswordHash(ctx, password)
	if err != nil {
		return nil, fmt.Errorf("failed hashing password: %w", err)
	}

	finalUser, err := s.Repository.User.Create(&user)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	accessToken, err := s.GenerateToken(ctx, finalUser.Email, AccessTokenDuration)
	if err != nil {
		return nil, fmt.Errorf("failed generating Access Token: %w", err)
	}

	refreshToken, err := s.GenerateToken(ctx, finalUser.Email, RefreshTokenDuration)
	if err != nil {
		return nil, fmt.Errorf("failed generating Refresh Token: %w", err)
	}

	return &types.JWT{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
