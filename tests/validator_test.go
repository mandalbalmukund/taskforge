package tests

import (
	"testing"

	"taskforge/internal/model"
	"taskforge/internal/registry"
	"taskforge/internal/validator"
)

func TestCycleDetection(t *testing.T) {

	r := registry.New()
	v := validator.New(r)

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

	errs := v.Validate(wf)

	if len(errs) == 0 {
		t.Fatal("expected cycle error")
	}
}