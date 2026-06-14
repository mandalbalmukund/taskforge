package timeout

import "context"

func IsTimeout(
	err error,
) bool {

	return err == context.DeadlineExceeded
}