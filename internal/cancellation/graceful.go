package cancellation

import (
	"context"
	"time"
)

func WaitForShutdown(
	ctx context.Context,
	gracePeriod time.Duration,
) error {

	select {

	case <-ctx.Done():
		return nil

	case <-time.After(
		gracePeriod,
	):
		return context.DeadlineExceeded
	}
}