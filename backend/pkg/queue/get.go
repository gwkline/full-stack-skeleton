package queue

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (s *Service) Get(ctx context.Context, name string) (*asynq.QueueInfo, error) {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("Get").End()

	// Get information about all queues
	queues, err := s.Inspector.Queues()
	if err != nil {
		return nil, fmt.Errorf("failed fetching queue info: %w", err)
	}

	for _, queue := range queues {
		if queue != name {
			continue
		}

		foundQueue, err := s.Inspector.GetQueueInfo(queue)
		if err != nil {
			return nil, fmt.Errorf("failed to get queue info: %w", err)
		}

		return foundQueue, nil
	}

	return nil, fmt.Errorf("queue %s not found", name)
}
