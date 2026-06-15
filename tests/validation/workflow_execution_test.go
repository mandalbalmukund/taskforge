package integration

import (
	"context"
	"testing"

	"taskforge/internal/engine"
	"taskforge/internal/scheduler"
	"taskforge/internal/state"
)

func TestWorkflowExecution(
	t *testing.T,
) {

	graph :=
		engine.BuildGraph(
			testWorkflow(),
		)

	shared :=
		state.New()

	s :=
		scheduler.New(
			testRegistry(),
			4,
		)

	err :=
		s.Run(
			context.Background(),
			graph,
			shared,
		)

	if err != nil {
		t.Fatal(err)
	}
}