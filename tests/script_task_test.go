package tests

import (
	"context"
	"testing"

	"taskforge/internal/model"
	scripttask "taskforge/internal/tasks/script"
)

func TestScriptTask(
	t *testing.T,
) {

	task := scripttask.New()

	result, err :=
		task.Execute(
			context.Background(),
			model.TaskDefinition{
				Config: map[string]interface{}{
					"command": "echo hello",
				},
			},
			nil,
		)

	if err != nil {
		t.Fatal(err)
	}

	if result.Output["exitCode"] != 0 {
		t.Fatal("expected success")
	}
}