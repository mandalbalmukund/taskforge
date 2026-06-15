package retry

import (
	"time"

	"taskforge/internal/model"
)

func CalculateBackoff(
	policy model.RetryPolicy,
	attempt int,
) time.Duration {

	delay :=
		policy.InitialWait

	for i := 0; i < attempt; i++ {

		delay =
			time.Duration(
				float64(delay) *
					policy.Multiplier,
			)
	}

	return delay
}