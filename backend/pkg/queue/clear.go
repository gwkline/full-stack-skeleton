package queue

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (s *Service) Clear(ctx context.Context, name string) ([]*asynq.TaskInfo, error) {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("Clear").End()

	// Get information about all queues
	queues, err := s.Inspector.Queues()
	if err != nil {
		return nil, fmt.Errorf("failed fetching queue info: %w", err)
	}

	for _, queue := range queues {
		if queue != name {
			continue
		}

		scheduledTasks, err := s.Inspector.ListScheduledTasks(name)
		if err != nil {
			return nil, fmt.Errorf("failed listing scheduled tasks: %w", err)
		}

		pendingTasks, err := s.Inspector.ListPendingTasks(name)
		if err != nil {
			return nil, fmt.Errorf("failed listing pending tasks: %w", err)
		}

		_, err = s.Inspector.DeleteAllScheduledTasks(name)
		if err != nil {
			return nil, fmt.Errorf("failed archiving scheduled tasks: %w", err)
		}

		_, err = s.Inspector.DeleteAllPendingTasks(name)
		if err != nil {
			return nil, fmt.Errorf("failed archiving scheduled tasks: %w", err)
		}

		combinedTasks := append(scheduledTasks, pendingTasks...)
		return combinedTasks, nil
	}

	return nil, fmt.Errorf("queue %s not found", name)
}
