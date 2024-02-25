package resolver

import (
	"context"

	"github.com/gwkline/full-stack-skeleton/backend/graph/generated"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Heartbeat is the resolver for the heartbeat field.
func (r *queryResolver) Heartbeat(ctx context.Context) (bool, error) {
	return true, nil
}

// Viewer is the resolver for the viewer field.
func (r *queryResolver) Viewer(ctx context.Context) (*types.User, error) {
	nrTx := newrelic.FromContext(ctx)
	defer nrTx.StartSegment("Viewer").End()

	return r.Service.Auth.CurrentUser(ctx)
}

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
