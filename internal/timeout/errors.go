package timeout

import "errors"

var (
	ErrTimeout =
		errors.New("task timeout")
)