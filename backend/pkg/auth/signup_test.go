package auth

import (
	"context"
	"fmt"
	"testing"

	"github.com/gwkline/full-stack-skeleton/backend/graph/generated"
	"github.com/gwkline/full-stack-skeleton/backend/mocks"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/repo"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/util"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/stretchr/testify/mock"
)

func TestSignup(t *testing.T) {
	s := Service{}
	hashedPW, _ := s.generatePasswordHash(context.Background(), "new_password")

	tests := []struct {
		name        string
		body        generated.UserInput
		mockFunc    func(*repo.Repository)
		expectedErr string
	}{
		{
			name: "Successful signup",
			body: generated.UserInput{Email: "rnadom@gmail.com", Password: "new_password"},
			mockFunc: func(db *repo.Repository) {
				db.User.(*mocks.IGenericRepo[types.User]).On("FindBy", []types.Filter{{Key: "email", Value: "rnadom@gmail.com"}}).Return(nil, fmt.Errorf("user not found"))

				user := &types.User{Email: "rnadom@gmail.com", PasswordHash: hashedPW}
				db.User.(*mocks.IGenericRepo[types.User]).On("Create", mock.Anything).Return(user, nil)
			},
			expectedErr: "",
		},
		{
			name: "Empty email and password",
			body: generated.UserInput{Email: "", Password: ""},
			mockFunc: func(db *repo.Repository) {
			},
			expectedErr: "invalid login",
		},
		{
			name: "User already exists",
			body: generated.UserInput{Email: "rnadom@gmail.com", Password: "new_password"},
			mockFunc: func(db *repo.Repository) {
				db.User.(*mocks.IGenericRepo[types.User]).On("FindBy", []types.Filter{{Key: "email", Value: "rnadom@gmail.com"}}).Return(&types.User{Email: "rnadom@gmail.com", PasswordHash: hashedPW}, nil)

			},
			expectedErr: "user already exists",
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

			tt.mockFunc(db)

			user := types.User{
				Email: "rnadom@gmail.com",
			}

			_, err := s.Signup(context.Background(), user, tt.body.Password)

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
