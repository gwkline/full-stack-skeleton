package handlers

import (
	"github.com/gwkline/full-stack-skeleton/backend/graph/loaders"
	"github.com/gwkline/full-stack-skeleton/backend/graph/resolver"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/repo"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	Repository *repo.Repository
	Redis      *redis.Client
	Loader     loaders.Loaders
	Service    *types.Services
}

func New(db *repo.Repository, redis *redis.Client, services *types.Services) *Handler {
	return &Handler{
		Repository: db,
		Redis:      redis,
		Loader:     *resolver.NewLoaders(db),
		Service:    services,
	}
}
