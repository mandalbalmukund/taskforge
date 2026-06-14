package registry

import (
	"context"

	"taskforge/internal/model"
)

type TaskResult struct {
	Output map[string]interface{}
}

type TaskHandler interface {
	Type() string

	Execute(
		ctx context.Context,
		task model.TaskDefinition,
		state map[string]interface{},
	) (TaskResult, error)
}