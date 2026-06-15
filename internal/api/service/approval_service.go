package service

import (
	"taskforge/internal/storage/postgres"
)

type ApprovalService struct {
	repo postgres.ApprovalRepository
}

func NewApprovalService(
	repo postgres.ApprovalRepository,
) *ApprovalService {

	return &ApprovalService{
		repo: repo,
	}
}

func (s *ApprovalService) Approve(
	executionID,
	taskID,
	user string,
) error {

	return s.repo.Approve(
		executionID,
		taskID,
		user,
	)
}

func (s *ApprovalService) Reject(
	executionID,
	taskID,
	user string,
) error {

	return s.repo.Reject(
		executionID,
		taskID,
		user,
	)
}