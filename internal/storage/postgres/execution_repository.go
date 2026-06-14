package postgres

import (
	"database/sql"
	"encoding/json"

	"taskforge/internal/model"
)

type executionRepo struct {
	db *sql.DB
}

func NewExecutionRepository(
	db *sql.DB,
) ExecutionRepository {

	return &executionRepo{
		db: db,
	}
}

func (r *executionRepo) Create(
	execution model.Execution,
) error {

	state, err :=
		json.Marshal(
			execution.State,
		)

	if err != nil {
		return err
	}

	_, err = r.db.Exec(
		`
		INSERT INTO executions
		(
		    id,
		    workflow_id,
		    status,
		    state,
		    started_at,
		    finished_at
		)
		VALUES
		($1,$2,$3,$4,$5,$6)
		`,
		execution.ID,
		execution.WorkflowID,
		execution.Status,
		state,
		execution.StartedAt,
		execution.FinishedAt,
	)

	return err
}

func (r *executionRepo) Get(
	id string,
) (*model.Execution, error) {

	row :=
		r.db.QueryRow(
			`
			SELECT
			 workflow_id,
			 status,
			 state,
			 started_at,
			 finished_at
			FROM executions
			WHERE id=$1
			`,
			id,
		)

	var exec model.Execution
	var state []byte

	err :=
		row.Scan(
			&exec.WorkflowID,
			&exec.Status,
			&state,
			&exec.StartedAt,
			&exec.FinishedAt,
		)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(
		state,
		&exec.State,
	)

	return &exec, nil
}

func (r *executionRepo) UpdateStatus(
	id string,
	status model.ExecutionStatus,
) error {

	_, err :=
		r.db.Exec(
			`
			UPDATE executions
			SET status=$1
			WHERE id=$2
			`,
			status,
			id,
		)

	return err
}