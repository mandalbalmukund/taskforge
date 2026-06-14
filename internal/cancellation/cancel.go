package cancellation

func (m *Manager) Cancel(
	executionID string,
) bool {

	m.mu.Lock()
	defer m.mu.Unlock()

	cancel, ok :=
		m.cancels[executionID]

	if !ok {
		return false
	}

	cancel()

	delete(
		m.cancels,
		executionID,
	)

	return true
}