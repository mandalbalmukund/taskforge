package tests

import (
	"context"
	"testing"

	"taskforge/internal/model"
	"taskforge/internal/registry"
)

type MockTask struct{}

func (m MockTask) Type() string {
	return "mock"
}

func (m MockTask) Execute(
	ctx context.Context,
	task model.TaskDefinition,
	state map[string]interface{},
) (registry.TaskResult, error) {

	return registry.TaskResult{
		Output: map[string]interface{}{
			"done": true,
		},
	}, nil
}