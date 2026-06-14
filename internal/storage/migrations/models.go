package postgres

import (
	"taskforge/internal/model"
)

type WorkflowRepository interface {
	Create(workflow model.Workflow) error
	Get(id string) (*model.Workflow, error)
}

type ExecutionRepository interface {
	Create(execution model.Execution) error
	Get(id string) (*model.Execution, error)
	UpdateStatus(id string, status model.ExecutionStatus) error
}

type ApprovalRepository interface {
	Create(
		executionID string,
		taskID string,
	) error

	Approve(
		executionID string,
		taskID string,
		approver string,
	) error

	Reject(
		executionID string,
		taskID string,
		approver string,
	) error
}