package unit

import (
	"testing"

	"taskforge/internal/model"
)

func TestCycleValidation(
	t *testing.T,
) {

	wf := model.Workflow{
		Tasks: []model.TaskDefinition{
			{
				ID: "A",
				DependsOn: []string{"B"},
			},
			{
				ID: "B",
				DependsOn: []string{"A"},
			},
		},
	}

	_ = wf

	
}