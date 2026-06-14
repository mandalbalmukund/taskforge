package model

type ExecutionStatus string

const (
	ExecutionPending   ExecutionStatus = "PENDING"
	ExecutionRunning   ExecutionStatus = "RUNNING"
	ExecutionCompleted ExecutionStatus = "COMPLETED"
	ExecutionFailed    ExecutionStatus = "FAILED"
	ExecutionCancelled ExecutionStatus = "CANCELLED"
)

type TaskStatus string

const (
	TaskPending          TaskStatus = "PENDING"
	TaskReady            TaskStatus = "READY"
	TaskRunning          TaskStatus = "RUNNING"
	TaskSuccess          TaskStatus = "SUCCESS"
	TaskFailed           TaskStatus = "FAILED"
	TaskSkipped          TaskStatus = "SKIPPED"
	TaskConditionFalse   TaskStatus = "CONDITION_FALSE"
	TaskWaitingApproval  TaskStatus = "WAITING_APPROVAL"
	TaskCancelled        TaskStatus = "CANCELLED"
	TaskTimedOut         TaskStatus = "TIMED_OUT"
)