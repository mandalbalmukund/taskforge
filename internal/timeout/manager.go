package timeout

import (
	"context"
	"time"
)

func WithTimeout(
	parent context.Context,
	duration time.Duration,
) (
	context.Context,
	context.CancelFunc,
) {

	return context.WithTimeout(
		parent,
		duration,
	)
}