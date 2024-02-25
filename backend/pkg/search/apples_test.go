package search

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gwkline/full-stack-skeleton/backend/mocks"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/repo"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/util"
	"github.com/gwkline/full-stack-skeleton/backend/types"
)

func TestApples(t *testing.T) {
	tests := []struct {
		name    string
		input   *types.ConnectionInput
		mock    func(r *repo.Repository)
		wantErr bool
	}{
		{
			name: "successful query",
			input: &types.ConnectionInput{
				SortBy:    nil,
				Direction: nil,
				Query:     nil,
				After:     nil,
				Filters:   nil,
				Limit:     nil,
			},
			mock: func(r *repo.Repository) {
				r.Apple.(*mocks.IGenericRepo[types.Apple]).EXPECT().Connection(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]*types.Apple{}, &types.PageInfo{}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()

			mockRepo := util.NewMockRepository(t)
			s := Service{Repository: mockRepo}

			tt.mock(mockRepo)

			_, err := s.Apples(context.Background(), tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
