package workers

import (
	"context"
	"fmt"

	"github.com/gwkline/full-stack-skeleton/backend/pkg/repo"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Service struct {
	Repository  *repo.Repository
	NewRelic    *newrelic.Application
	SomeService types.ISomeService
}

func Init(repo *repo.Repository, nrApp *newrelic.Application, someService types.ISomeService) *Service {
	return &Service{
		Repository:  repo,
		NewRelic:    nrApp,
		SomeService: someService,
	}
}

func (s *Service) setupWorkerContext(ctx context.Context, name string) (context.Context, *newrelic.Transaction) {
	tx := s.NewRelic.StartTransaction(fmt.Sprintf("worker::%s", name))
	ctx = newrelic.NewContext(ctx, tx)

	return ctx, tx
}
