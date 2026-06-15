package retry

import "errors"

var (
	ErrRetryable = errors.New("retryable")
	ErrPermanent = errors.New("permanent")
)

func IsRetryable(err error) bool {

	return errors.Is(
		err,
		ErrRetryable,
	)
}