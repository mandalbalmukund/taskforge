package model

import "time"

type Execution struct {
	ID          string
	WorkflowID  string
	Status      ExecutionStatus
	State       map[string]interface{}
	StartedAt   time.Time
	FinishedAt  *time.Time
}