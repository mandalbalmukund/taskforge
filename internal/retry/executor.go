package retry

import (
	"context"
	"time"

	"taskforge/internal/model"
)

func ExecuteWithRetry(
	ctx context.Context,
	policy model.RetryPolicy,
	fn func() error,
) error {

	var err error

	for attempt := 0;
		attempt <= policy.MaxRetries;
		attempt++ {

		err = fn()

		if err == nil {
			return nil
		}

		if !IsRetryable(err) {
			return err
		}

		if attempt ==
			policy.MaxRetries {

			return err
		}

		wait :=
			CalculateBackoff(
				policy,
				attempt,
			)

		select {

		case <-ctx.Done():
			return ctx.Err()

		case <-time.After(wait):
		}
	}

	return err
}