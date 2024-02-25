package auth

import (
	"context"
	"testing"
	"time"

	"github.com/gwkline/full-stack-skeleton/backend/types"
)

func TestRefreshToken(t *testing.T) {
	s := &Service{}
	tests := []struct {
		name        string
		body        types.JWT
		expectedErr string
	}{
		{
			name: "Valid tokens",
			body: types.JWT{
				AccessToken: func() string {
					token, _ := s.GenerateToken(context.Background(), "email@gmail.com", time.Second*1)
					return token
				}(),
				RefreshToken: func() string {
					token, _ := s.GenerateToken(context.Background(), "email@gmail.com", time.Hour*1)
					return token
				}(),
			},
			expectedErr: "",
		},
		{
			name: "Invalid email",
			body: types.JWT{
				AccessToken: func() string {
					token, _ := s.GenerateToken(context.Background(), "email@gmail.com", time.Second*1)
					return token
				}(),
				RefreshToken: func() string {
					token, _ := s.GenerateToken(context.Background(), "email123@gmail.com", time.Hour*1)
					return token
				}(),
			},
			expectedErr: "invalid refresh token",
		},
		{
			name: "Expired token",
			body: types.JWT{
				AccessToken: func() string {
					token, _ := s.GenerateToken(context.Background(), "email@gmail.com", time.Second*1)
					return token
				}(),
				RefreshToken: func() string {
					token, _ := s.GenerateToken(context.Background(), "email@gmail.com", time.Hour*0)
					return token
				}(),
			},
			expectedErr: "invalid refresh token: invalid token: token has invalid claims: token is expired",
		},
		{
			name: "Malformed token",
			body: types.JWT{
				AccessToken: func() string {
					token, _ := s.GenerateToken(context.Background(), "email@gmail.com", time.Second*1)
					return token
				}(),
				RefreshToken: "6969",
			},
			expectedErr: "invalid refresh token: invalid token: token is malformed: token contains an invalid number of segments",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()

			_, err := s.RefreshToken(context.Background(), tt.body)

			if err != nil {
				if err.Error() != tt.expectedErr {
					t.Errorf("expected error: %v, got: %v", tt.expectedErr, err)
				}
			} else if tt.expectedErr != "" {
				t.Errorf("expected error: %v, got: nil", tt.expectedErr)
			}
		})
	}
}
