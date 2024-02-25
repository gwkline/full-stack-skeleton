package queue

import (
	"context"
	"fmt"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func (s *Service) TogglePause(ctx context.Context, name string) error {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("TogglePause").End()

	queue, err := s.Get(ctx, name)
	if err != nil {
		return fmt.Errorf("failed to get queue info: %w", err)
	}

	if queue.Paused {
		err := s.Inspector.UnpauseQueue(name)
		if err != nil {
			return fmt.Errorf("failed unpausing queue: %w", err)
		}
		return nil
	}

	s.Inspector.PauseQueue(name)
	if err != nil {
		return fmt.Errorf("failed pausing queue: %w", err)
	}

	return nil
}
