package tests

import (
	"testing"

	httptask "taskforge/internal/tasks/http"
)

func TestRetryableStatus(
	t *testing.T,
) {

	if !httptask.RetryableStatus(
		500,
	) {

		t.Fatal(
			"500 should retry",
		)
	}
}