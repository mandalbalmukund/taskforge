package scheduler

import (
	"sync"

	"taskforge/internal/model"
)

type ExecutionState struct {
	mu    sync.RWMutex
	Tasks map[string]*NodeState
}

func NewExecutionState() *ExecutionState {

	return &ExecutionState{
		Tasks: make(map[string]*NodeState),
	}
}

func (s *ExecutionState) SetStatus(
	taskID string,
	status model.TaskStatus,
) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.Tasks[taskID].Status = status
}

func (s *ExecutionState) GetStatus(
	taskID string,
) model.TaskStatus {

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.Tasks[taskID].Status
}