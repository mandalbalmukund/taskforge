package service

import (
	"time"

	"github.com/google/uuid"

	"taskforge/internal/model"
	"taskforge/internal/storage/postgres"
)

type WorkflowService struct {
	repo postgres.WorkflowRepository
}

func NewWorkflowService(
	repo postgres.WorkflowRepository,
) *WorkflowService {

	return &WorkflowService{
		repo: repo,
	}
}

func (s *WorkflowService) Create(
	workflow model.Workflow,
) (*model.Workflow, error) {

	workflow.ID =
		uuid.New().String()

	workflow.CreatedAt =
		time.Now()

	err :=
		s.repo.Create(workflow)

	if err != nil {
		return nil, err
	}

	return &workflow, nil
}

func (s *WorkflowService) Get(
	id string,
) (*model.Workflow, error) {

	return s.repo.Get(id)
}