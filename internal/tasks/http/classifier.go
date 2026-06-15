package http

func RetryableStatus(code int) bool {

	if code == 429 {
		return true
	}

	if code >= 500 {
		return true
	}

	return false
}