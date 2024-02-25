package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

const (
	TypeProcessSomeData = "random_pkg::process_some_data"
)

type someTaskPayload struct {
	SomeInfo string `json:"some_info"`
}

func NewProcessSomeDataTask(someinfo string, processIn int) (*asynq.Task, error) {
	payload, err := json.Marshal(someTaskPayload{SomeInfo: someinfo})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeProcessSomeData, payload, asynq.ProcessIn(time.Second*time.Duration(processIn))), nil
}

func (s *Service) HandleProcessSomeDataTask(ctx context.Context, t *asynq.Task) error {
	var p someTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	fmt.Println(" [*] Processing Teddy Email")
	ctx, tx := s.setupWorkerContext(ctx, "ProcessSomeData")
	defer tx.End()

	return s.SomeService.SomeMethod(ctx)
}
