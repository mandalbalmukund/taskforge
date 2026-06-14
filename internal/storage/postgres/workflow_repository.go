package postgres

import (
	"database/sql"
	"encoding/json"

	"taskforge/internal/model"
)

type workflowRepo struct {
	db *sql.DB
}

func NewWorkflowRepository(
	db *sql.DB,
) WorkflowRepository {

	return &workflowRepo{
		db: db,
	}
}

func (r *workflowRepo) Create(
	workflow model.Workflow,
) error {

	definition, err :=
		json.Marshal(workflow)

	if err != nil {
		return err
	}

	_, err = r.db.Exec(
		`
		INSERT INTO workflows
		(
		    id,
		    name,
		    version,
		    description,
		    definition,
		    created_at
		)
		VALUES
		($1,$2,$3,$4,$5,$6)
		`,
		workflow.ID,
		workflow.Name,
		workflow.Version,
		workflow.Description,
		definition,
		workflow.CreatedAt,
	)

	return err
}

func (r *workflowRepo) Get(
	id string,
) (*model.Workflow, error) {

	row :=
		r.db.QueryRow(
			`
			SELECT definition
			FROM workflows
			WHERE id=$1
			`,
			id,
		)

	var definition []byte

	if err := row.Scan(
		&definition,
	); err != nil {

		return nil, err
	}

	var workflow model.Workflow

	if err :=
		json.Unmarshal(
			definition,
			&workflow,
		); err != nil {

		return nil, err
	}

	return &workflow, nil
}