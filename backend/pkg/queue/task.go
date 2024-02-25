package queue

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func (s *Service) RunTask(ctx context.Context, queueName string, taskId string) error {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("RunTask").End()

	return s.Inspector.RunTask(queueName, taskId)
}

func (s *Service) DeleteTask(ctx context.Context, queueName string, taskId string) error {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("DeleteTask").End()

	return s.Inspector.DeleteTask(queueName, taskId)
}

func (s *Service) ArchiveTask(ctx context.Context, queueName string, taskId string) error {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("ArchiveTask").End()

	return s.Inspector.ArchiveTask(queueName, taskId)
}
