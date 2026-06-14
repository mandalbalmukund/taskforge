package tests

import (
	"testing"

	approval "taskforge/internal/tasks/approval"
)

func TestApprovalStore(
	t *testing.T,
) {

	store :=
		approval.NewStore()

	store.Approve("task1")

	if store.Get("task1") !=
		approval.Approved {

		t.Fatal(
			"approval failed",
		)
	}
}