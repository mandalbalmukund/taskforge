package extensibility

import (
	"context"
	"testing"

	"taskforge/internal/model"
	"taskforge/internal/registry"
)

type SlackTask struct{}

func (s SlackTask) Type() string {
	return "slack"
}

func (s SlackTask) Execute(
	ctx context.Context,
	task model.TaskDefinition,
	state map[string]interface{},
) (
	registry.TaskResult,
	error,
) {

	return registry.TaskResult{
		Output: map[string]interface{}{
			"sent": true,
		},
	}, nil
}

func TestPluginArchitecture(
	t *testing.T,
) {

	r := registry.New()

	r.Register(
		SlackTask{},
	)

	if !r.Exists("slack") {

		t.Fatal(
			"plugin registration failed",
		)
	}
}