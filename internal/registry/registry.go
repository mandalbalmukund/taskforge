package registry

import (
	"fmt"
	"sync"
)

type Registry struct {
	mu       sync.RWMutex
	handlers map[string]TaskHandler
}

func New() *Registry {
	return &Registry{
		handlers: make(map[string]TaskHandler),
	}
}

func (r *Registry) Register(handler TaskHandler) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.handlers[handler.Type()] = handler
}

func (r *Registry) Get(taskType string) (TaskHandler, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	handler, ok := r.handlers[taskType]
	if !ok {
		return nil, fmt.Errorf("unknown task type: %s", taskType)
	}

	return handler, nil
}

func (r *Registry) Exists(taskType string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.handlers[taskType]
	return ok
}