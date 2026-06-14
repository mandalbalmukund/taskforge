package state

import "sync"

type SharedState struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

func New() *SharedState {

	return &SharedState{
		data: map[string]interface{}{
			"tasks": map[string]interface{}{},
		},
	}
}

func (s *SharedState) Snapshot()
	map[string]interface{} {

	s.mu.RLock()
	defer s.mu.RUnlock()

	copyMap := make(
		map[string]interface{},
		len(s.data),
	)

	for k, v := range s.data {
		copyMap[k] = v
	}

	return copyMap
}