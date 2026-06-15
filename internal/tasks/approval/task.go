package approval

import (
	"context"
	"fmt"
	"time"

	"taskforge/internal/model"
	"taskforge/internal/registry"
)

type Task struct {
	store *ApprovalStore
}

func New(
	store *ApprovalStore,
) *Task {

	return &Task{
		store: store,
	}
}

func (t *Task) Type() string {
	return "approval"
}

func (t *Task) Execute(
	ctx context.Context,
	task model.TaskDefinition,
	state map[string]interface{},
) (registry.TaskResult, error) {

	taskID := task.ID

	timeout :=
		task.Config["timeout"].
			(time.Duration)

	ticker :=
		time.NewTicker(
			time.Second,
		)

	defer ticker.Stop()

	timeoutChan :=
		time.After(timeout)

	for {

		select {

		case <-ctx.Done():
			return registry.TaskResult{},
				ctx.Err()

		case <-timeoutChan:
			return registry.TaskResult{},
				fmt.Errorf(
					"approval timeout",
				)

		case <-ticker.C:

			status :=
				t.store.Get(taskID)

			switch status {

			case Approved:

				return registry.TaskResult{
					Output: map[string]interface{}{
						"approval": "approved",
					},
				}, nil

			case Rejected:

				return registry.TaskResult{},
					fmt.Errorf(
						"approval rejected",
					)
			}
		}
	}
}