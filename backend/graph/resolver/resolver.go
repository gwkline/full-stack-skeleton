package resolver

import (
	"github.com/gwkline/full-stack-skeleton/backend/graph/loaders"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/repo"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/redis/go-redis/v9"
)

//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repository *repo.Repository
	Redis      *redis.Client
	Loader     loaders.Loaders
	Service    *types.Services
}
