package queue

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (s *Service) Enqueue(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("Enqueue").End()

	return s.Client.Enqueue(task, opts...)
}
