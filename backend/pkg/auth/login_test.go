package auth

import (
	"context"
	"fmt"
	"testing"

	"github.com/gwkline/full-stack-skeleton/backend/mocks"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/util"
	"github.com/gwkline/full-stack-skeleton/backend/types"
)

func TestLogin(t *testing.T) {
	s := Service{}
	hashedPW, _ := s.generatePasswordHash(context.Background(), "password123")

	tests := []struct {
		name        string
		body        types.LoginInput
		mockUser    *types.User
		mockErr     error
		expectedErr string
	}{
		{
			name:        "Login",
			body:        types.LoginInput{Email: "email@gmail.com", Password: "password123"},
			mockUser:    &types.User{Email: "email@gmail.com", PasswordHash: hashedPW},
			mockErr:     nil,
			expectedErr: "",
		},
		{
			name:        "EmptyEmailPassword",
			body:        types.LoginInput{Email: "", Password: ""},
			mockUser:    nil,
			mockErr:     fmt.Errorf("some err"),
			expectedErr: "user not found: some err",
		},
		{
			name:        "UserNotFound",
			body:        types.LoginInput{Email: "email@gmail.com", Password: "password123"},
			mockUser:    nil,
			mockErr:     fmt.Errorf("some err"),
			expectedErr: "user not found: some err",
		},
		{
			name:        "InvalidatePassword",
			body:        types.LoginInput{Email: "email@gmail.com", Password: "password123"},
			mockUser:    &types.User{Email: "email@gmail.com", PasswordHash: "hashedPassword1234"},
			mockErr:     nil,
			expectedErr: "invalid password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()

			db := util.NewMockRepository(t)
			s := &Service{
				Repository: db,
			}

			db.User.(*mocks.IGenericRepo[types.User]).On("FindBy", []types.Filter{{Key: "email", Value: tt.body.Email}}).Return(tt.mockUser, tt.mockErr)

			_, err := s.Login(context.Background(), tt.body)

			if err != nil {
				if err.Error() != tt.expectedErr {
					t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				}
			} else if tt.expectedErr != "" {
				t.Errorf("expected error %v, got nil", tt.expectedErr)
			}
		})
	}
}
