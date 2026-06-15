package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"taskforge/internal/model"
	"taskforge/internal/scheduler"
	"taskforge/internal/state"
	"taskforge/internal/storage/postgres"
)

type ExecutionService struct {
	repo      postgres.ExecutionRepository
	scheduler *scheduler.Scheduler
}

func NewExecutionService(
	repo postgres.ExecutionRepository,
	scheduler *scheduler.Scheduler,
) *ExecutionService {

	return &ExecutionService{
		repo: repo,
		scheduler: scheduler,
	}
}

func (s *ExecutionService) Start(
	ctx context.Context,
	workflowID string,
) (*model.Execution, error) {

	exec := model.Execution{
		ID:         uuid.New().String(),
		WorkflowID: workflowID,
		Status:     model.ExecutionRunning,
		State:      map[string]interface{}{},
		StartedAt:  time.Now(),
	}

	if err :=
		s.repo.Create(exec); err != nil {

		return nil, err
	}

	_ = state.New()

	return &exec, nil
}

func (s *ExecutionService) Get(
	id string,
) (*model.Execution, error) {

	return s.repo.Get(id)
}