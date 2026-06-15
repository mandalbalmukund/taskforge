package cancellation

import (
	"context"
	"sync"
)

type Manager struct {
	mu      sync.RWMutex
	cancels map[string]context.CancelFunc
}

func New() *Manager {

	return &Manager{
		cancels: make(
			map[string]context.CancelFunc,
		),
	}
}