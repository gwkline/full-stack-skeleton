package resolver

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gwkline/full-stack-skeleton/backend/types"
)

func (r *Resolver) HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role types.Role) (res interface{}, err error) {
	user, err := r.Service.Auth.CurrentUser(ctx)
	if err != nil {
		return nil, errors.New("user not found in context")
	}

	if role == "ADMIN" && !user.IsAdmin() {
		return nil, errors.New("you're missing the required permissions to perform this operation")
	}

	return next(ctx)
}
