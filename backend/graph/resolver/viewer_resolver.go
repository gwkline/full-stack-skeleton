package resolver

import (
	"context"
	"fmt"

	"github.com/gwkline/full-stack-skeleton/backend/graph/generated"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// UserConnection is the resolver for the userConnection field.
func (r *viewerResolver) AppleConnection(ctx context.Context, obj *types.User, input types.ConnectionInput) (*types.AppleConnection, error) {
	nrTx := newrelic.FromContext(ctx)
	defer nrTx.StartSegment("UserConnection").End()

	return r.Service.Search.Apples(ctx, &input)
}

// Queue is the resolver for the queue field.
func (r *viewerResolver) Queue(ctx context.Context, obj *types.User) (*generated.Queue, error) {
	nrTx := newrelic.FromContext(ctx)
	defer nrTx.StartSegment("Queue").End()

	thisQueue, err := r.Service.Queue.Get(ctx, "queue")
	if err != nil {
		return nil, fmt.Errorf("failed to get queue info: %w", err)
	}

	return &generated.Queue{
		Aggregating:      thisQueue.Aggregating,
		Name:             thisQueue.Queue,
		MemoryUsageBytes: int(thisQueue.MemoryUsage),
		Size:             thisQueue.Size,
		Groups:           thisQueue.Groups,
		LatencyMsec:      int(thisQueue.Latency),
		DisplayLatency:   fmt.Sprintf("%v", thisQueue.Latency),
		Active:           thisQueue.Active,
		Pending:          thisQueue.Pending,
		Scheduled:        thisQueue.Scheduled,
		Retry:            thisQueue.Retry,
		Archived:         thisQueue.Archived,
		Completed:        thisQueue.Completed,
		Processed:        thisQueue.Processed,
		Failed:           thisQueue.Failed,
		Paused:           thisQueue.Paused,
		Timestamp:        thisQueue.Timestamp.String(),
	}, nil
}

func (r *Resolver) Viewer() generated.ViewerResolver { return &viewerResolver{r} }

type viewerResolver struct{ *Resolver }
