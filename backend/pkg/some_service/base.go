package some_service

import (
	"github.com/gwkline/full-stack-skeleton/backend/pkg/repo"
)

type Service struct {
	Repository *repo.Repository
}

func Init(repo *repo.Repository) *Service {
	return &Service{Repository: repo}
}
