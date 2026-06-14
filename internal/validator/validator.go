package validator

import (
	"fmt"

	"taskforge/internal/model"
	"taskforge/internal/registry"
)

type Validator struct {
	registry *registry.Registry
}

func New(r *registry.Registry) *Validator {
	return &Validator{
		registry: r,
	}
}

func (v *Validator) Validate(
	workflow model.Workflow,
) []ValidationError {

	var errs []ValidationError

	taskMap := make(map[string]model.TaskDefinition)

	for _, task := range workflow.Tasks {

		if _, exists := taskMap[task.ID]; exists {
			errs = append(errs,
				ValidationError{
					Field: task.ID,
					Message: "duplicate task id",
				})
		}

		taskMap[task.ID] = task

		if !v.registry.Exists(task.Type) {
			errs = append(errs,
				ValidationError{
					Field: task.ID,
					Message: fmt.Sprintf(
						"unknown task type %s",
						task.Type,
					),
				})
		}
	}

	for _, task := range workflow.Tasks {

		for _, dep := range task.DependsOn {

			if _, exists := taskMap[dep]; !exists {
				errs = append(errs,
					ValidationError{
						Field: task.ID,
						Message: fmt.Sprintf(
							"missing dependency %s",
							dep,
						),
					})
			}
		}
	}

	if cycleExists(workflow.Tasks) {
		errs = append(errs,
			ValidationError{
				Field: "workflow",
				Message: "cycle detected",
			})
	}

	return errs
}