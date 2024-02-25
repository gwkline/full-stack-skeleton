package resolver

import (
	"context"
	"fmt"
	"time"

	"github.com/gwkline/full-stack-skeleton/backend/graph/generated"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Queue is the resolver for the queue field.
func (r *subscriptionResolver) Queue(ctx context.Context) (<-chan *generated.Queue, error) {
	tx := newrelic.FromContext(ctx)
	tx.Ignore()

	ch := make(chan *generated.Queue)
	go func() {
		for {
			thisQueue, err := r.Service.Queue.Get(ctx, "queue")
			if err != nil {
				fmt.Printf("failed to get queue info: %s\n", err)
				return
			}

			select {
			case <-ctx.Done():
				close(ch)
				return
			case ch <- &generated.Queue{
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
			}:
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return ch, nil
}

func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
