package approval

import (
	"sync"
)

type Status string

const (
	Pending  Status = "PENDING"
	Approved Status = "APPROVED"
	Rejected Status = "REJECTED"
)

type ApprovalStore struct {
	mu       sync.RWMutex
	decisions map[string]Status
}

func NewStore() *ApprovalStore {

	return &ApprovalStore{
		decisions: make(map[string]Status),
	}
}

func (s *ApprovalStore) Approve(id string) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.decisions[id] = Approved
}

func (s *ApprovalStore) Reject(id string) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.decisions[id] = Rejected
}

func (s *ApprovalStore) Get(id string) Status {

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.decisions[id]
}