package kubernetes

import (
	"context"
	"fmt"

	"taskforge/internal/model"
	"taskforge/internal/registry"
)

type Task struct{}

func New() *Task {
	return &Task{}
}

func (t *Task) Type() string {
	return "kubernetes"
}

func (t *Task) Execute(
	ctx context.Context,
	task model.TaskDefinition,
	state map[string]interface{},
) (registry.TaskResult, error) {

	action :=
		task.Config["action"].(string)

	deployment :=
		task.Config["deployment"].(string)

	switch action {

	case "rollout_restart":

		return registry.TaskResult{
			Output: map[string]interface{}{
				"deployment": deployment,
				"status":     "restarted",
			},
		}, nil

	case "scale":

		return registry.TaskResult{
			Output: map[string]interface{}{
				"deployment": deployment,
				"status":     "scaled",
			},
		}, nil

	default:

		return registry.TaskResult{},
			fmt.Errorf(
				"unknown action",
			)
	}
}