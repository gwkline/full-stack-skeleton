package queue

import (
	"os"
	"time"

	"github.com/hibiken/asynq"
)

type Service struct {
	Client    *asynq.Client
	Inspector *asynq.Inspector
	Server    *asynq.Server
}

func Init(ops *asynq.RedisClientOpt) *Service {
	var logLevel asynq.LogLevel
	switch os.Getenv("ENV") {
	case "development":
		logLevel = asynq.InfoLevel
	default:
		logLevel = asynq.ErrorLevel
	}

	return &Service{
		Client:    asynq.NewClient(ops),
		Inspector: asynq.NewInspector(ops),
		Server: asynq.NewServer(
			ops,
			asynq.Config{
				Concurrency:     10,
				ShutdownTimeout: time.Second * 15,
				LogLevel:        logLevel,
				Queues: map[string]int{
					"another": 6, // processed 60% of the time
					"default": 3, // processed 30% of the time
					"low":     1, // processed 10% of the time
				},
			},
		),
	}
}
