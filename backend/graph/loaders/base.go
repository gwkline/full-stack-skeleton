package loaders

import (
	"context"

	"github.com/gwkline/full-stack-skeleton/backend/pkg/repo"
	"github.com/gwkline/full-stack-skeleton/backend/types"
)

type Loaders struct {
	AppleByID  *AppleByIDLoader
	Repository repo.Repository
}

func (l *Loaders) For(ctx context.Context) *Loaders {
	value := ctx.Value(types.LoadersKey)
	if value == nil {
		return nil
	}
	return value.(*Loaders)
}
