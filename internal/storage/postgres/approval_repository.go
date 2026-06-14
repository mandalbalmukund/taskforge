package postgres

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type approvalRepo struct {
	db *sql.DB
}

func NewApprovalRepository(
	db *sql.DB,
) ApprovalRepository {

	return &approvalRepo{
		db: db,
	}
}

func (r *approvalRepo) Create(
	executionID string,
	taskID string,
) error {

	_, err :=
		r.db.Exec(
			`
			INSERT INTO approvals
			(
			 id,
			 execution_id,
			 task_id
			)
			VALUES
			($1,$2,$3)
			`,
			uuid.New().String(),
			executionID,
			taskID,
		)

	return err
}

func (r *approvalRepo) Approve(
	executionID string,
	taskID string,
	approver string,
) error {

	_, err :=
		r.db.Exec(
			`
			UPDATE approvals
			SET
			  decision='APPROVED',
			  approver=$1,
			  approved_at=$2
			WHERE execution_id=$3
			  AND task_id=$4
			`,
			approver,
			time.Now(),
			executionID,
			taskID,
		)

	return err
}

func (r *approvalRepo) Reject(
	executionID string,
	taskID string,
	approver string,
) error {

	_, err :=
		r.db.Exec(
			`
			UPDATE approvals
			SET
			  decision='REJECTED',
			  approver=$1,
			  approved_at=$2
			WHERE execution_id=$3
			  AND task_id=$4
			`,
			approver,
			time.Now(),
			executionID,
			taskID,
		)

	return err
}