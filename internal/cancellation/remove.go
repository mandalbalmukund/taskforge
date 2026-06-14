package cancellation

func (m *Manager) Remove(
	executionID string,
) {

	m.mu.Lock()
	defer m.mu.Unlock()

	delete(
		m.cancels,
		executionID,
	)
}