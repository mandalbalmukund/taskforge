package tests

import (
	"context"
	"testing"
	"time"

	"taskforge/internal/model"
	"taskforge/internal/retry"
)

func TestRetrySuccess(
	t *testing.T,
) {

	count := 0

	err :=
		retry.ExecuteWithRetry(
			context.Background(),
			model.RetryPolicy{
				MaxRetries: 3,
				InitialWait: time.Millisecond,
				Multiplier: 2,
			},
			func() error {

				count++

				if count < 3 {
					return retry.ErrRetryable
				}

				return nil
			},
		)

	if err != nil {
		t.Fatal(err)
	}

	if count != 3 {
		t.Fatal(
			"expected 3 attempts",
		)
	}
}