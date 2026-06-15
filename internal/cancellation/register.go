package cancellation

import "context"

func (m *Manager) Register(
	executionID string,
	cancel context.CancelFunc,
) {

	m.mu.Lock()
	defer m.mu.Unlock()

	m.cancels[executionID] =
		cancel
}