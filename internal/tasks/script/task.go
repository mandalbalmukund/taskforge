package script

import (
	"bytes"
	"context"
	"os/exec"

	"taskforge/internal/model"
	"taskforge/internal/registry"
)

type Task struct{}

func New() *Task {
	return &Task{}
}

func (t *Task) Type() string {
	return "script"
}

func (t *Task) Execute(
	ctx context.Context,
	task model.TaskDefinition,
	state map[string]interface{},
) (registry.TaskResult, error) {

	command :=
		task.Config["command"].(string)

	cmd :=
		exec.CommandContext(
			ctx,
			"sh",
			"-c",
			command,
		)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	output := map[string]interface{}{
		"stdout": stdout.String(),
		"stderr": stderr.String(),
	}

	if err != nil {

		output["exitCode"] = -1

		return registry.TaskResult{
			Output: output,
		}, err
	}

	output["exitCode"] = 0

	return registry.TaskResult{
		Output: output,
	}, nil
}