package tests

import (
	"context"
	"testing"

	"taskforge/internal/cancellation"
)

func TestCancellation(
	t *testing.T,
) {

	manager :=
		cancellation.New()

	ctx, cancel :=
		context.WithCancel(
			context.Background(),
		)

	manager.Register(
		"exec1",
		cancel,
	)

	ok :=
		manager.Cancel(
			"exec1",
		)

	if !ok {
		t.Fatal(
			"cancel failed",
		)
	}

	select {

	case <-ctx.Done():

	default:

		t.Fatal(
			"context should be cancelled",
		)
	}
}