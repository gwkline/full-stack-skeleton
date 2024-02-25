package queue

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gwkline/full-stack-skeleton/backend/pkg/workers"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/hibiken/asynq"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (s *Service) StartServer(ctx context.Context, workPkg types.IWorker) {
	tx := newrelic.FromContext(ctx)
	defer tx.StartSegment("StartServer").End()

	fmt.Println("Initializing Asynq")

	st := workPkg
	mux := asynq.NewServeMux()
	mux.HandleFunc(workers.TypeProcessSomeData, st.HandleProcessSomeDataTask)

	if err := s.Server.Run(mux); err != nil {
		fmt.Printf("oops, mux err: %v\n", err)
	}

	// Listen for SIGTERM signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Wait for the signal
	<-stop

	// Shutdown the server
	s.Server.Shutdown()
}
