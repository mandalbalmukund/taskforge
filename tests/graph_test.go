package tests

import (
	"testing"

	"taskforge/internal/engine"
	"taskforge/internal/model"
)

func TestGraphBuilder(t *testing.T) {

	tasks := []model.TaskDefinition{
		{
			ID: "A",
		},
		{
			ID: "B",
			DependsOn: []string{
				"A",
			},
		},
	}

	graph :=
		engine.BuildGraph(tasks)

	if graph.Dependencies["B"] != 1 {
		t.Fatal(
			"dependency count mismatch",
		)
	}
}