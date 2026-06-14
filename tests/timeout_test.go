package tests

import (
	"context"
	"testing"
	"time"

	"taskforge/internal/timeout"
)

func TestTimeout(
	t *testing.T,
) {

	ctx, cancel :=
		timeout.WithTimeout(
			context.Background(),
			time.Millisecond,
		)

	defer cancel()

	time.Sleep(
		10 * time.Millisecond,
	)

	if ctx.Err() == nil {
		t.Fatal(
			"expected timeout",
		)
	}
}