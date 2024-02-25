package cron

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gdgvda/cron"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/workers"
	"github.com/gwkline/full-stack-skeleton/backend/types"
)

type Service struct {
	Cron  *cron.Cron
	Queue types.IQueue
}

func Init(q types.IQueue) *Service {
	nyc, _ := time.LoadLocation("America/New_York")
	c := cron.New(cron.WithLocation(nyc))
	c.Start()
	return &Service{
		Cron:  c,
		Queue: q,
	}
}

func (s *Service) SetupJobs(services *types.Services) {
	if os.Getenv("ENV") != "production" {
		return
	}

	// Monday-Friday (1-5) at 8am ET, queue a worker to do something
	s.Cron.Add("0 8 * * 1-5", func() {
		s.queueSomeJob()
	})
}

func (s *Service) Stop() context.Context {
	return s.Cron.Stop()
}

func (s *Service) queueSomeJob() {
	delay := rand.Intn(1 * 60 * 60) // generate a random delay under 1 hours

	t, err := workers.NewProcessSomeDataTask("passing some data in a worker", delay)
	if err != nil {
		fmt.Printf("failure making new process some data task: %s", err)
	}

	_, err = s.Queue.Enqueue(context.Background(), t)
	if err != nil {
		fmt.Printf("failure queueing task: %s", err)
	}
}
