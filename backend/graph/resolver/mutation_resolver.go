package resolver

import (
	"context"
	"fmt"

	"github.com/gwkline/full-stack-skeleton/backend/graph/generated"
	"github.com/gwkline/full-stack-skeleton/backend/types"
)

func (r *mutationResolver) CreateApple(ctx context.Context, input generated.AppleCreateInput) ([]*types.Apple, error) {
	var apples []*types.Apple
	for range input.Quantity {
		apple, err := r.Repository.Apple.Create(&types.Apple{
			UserID:  input.UserID,
			Variety: input.Variety,
		})
		if err != nil {
			return nil, err
		}

		apples = append(apples, apple)
	}

	return apples, nil
}

func (r *mutationResolver) UpdateApple(ctx context.Context, input generated.AppleUpdateInput) (bool, error) {
	apple, err := r.Repository.Apple.FindBy([]types.Filter{{Key: "id", Value: input.AppleID}})
	if err != nil {
		return false, err
	}

	if input.Variety != nil {
		apple.Variety = *input.Variety
	}

	if input.UserID != nil {
		apple.UserID = *input.UserID
	}

	_, err = r.Repository.Apple.Update(apple)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) DeleteApple(ctx context.Context, input generated.DeleteInput) (bool, error) {
	apples, err := r.Repository.Apple.ListBy([]types.Filter{{Key: "id", Value: input.Ids}})
	if err != nil {
		return false, err
	}

	for _, apple := range apples {
		err := r.Repository.Apple.Archive(apple)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input types.LoginInput) (*types.JWT, error) {
	res, err := r.Service.Auth.Login(ctx, input)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Signup is the resolver for the signup field.
func (r *mutationResolver) Signup(ctx context.Context, input generated.UserInput) (*types.JWT, error) {
	user := types.User{
		Email: input.Email,
	}
	res, err := r.Service.Auth.Signup(ctx, user, input.Password)
	if err != nil {
		return res, err
	}

	return res, nil
}

// RefreshToken is the resolver for the RefreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input generated.RefreshInput) (*types.JWT, error) {
	jwt := types.JWT{
		RefreshToken: input.RefreshToken,
		AccessToken:  input.AccessToken,
	}
	res, err := r.Service.Auth.RefreshToken(ctx, jwt)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ToggleQueuePaused is the resolver for the toggleQueuePaused field.
func (r *mutationResolver) ToggleQueuePaused(ctx context.Context) (bool, error) {
	err := r.Service.Queue.TogglePause(ctx, "queue")
	if err != nil {
		return false, fmt.Errorf("failed toggling queue pause: %w", err)
	}

	return true, nil
}

// ClearQueue is the resolver for the clearQueue field.
func (r *mutationResolver) ClearQueue(ctx context.Context, input generated.ClearQueueInput) (bool, error) {
	_, err := r.Service.Queue.Clear(ctx, "queue")
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
