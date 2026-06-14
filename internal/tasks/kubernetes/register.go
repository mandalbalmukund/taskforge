package tasks

import (
	"taskforge/internal/registry"

	approvaltask "taskforge/internal/tasks/approval"
	httptask "taskforge/internal/tasks/http"
	k8stask "taskforge/internal/tasks/kubernetes"
	scripttask "taskforge/internal/tasks/script"
)

func RegisterAll(
	reg *registry.Registry,
) {

	reg.Register(
		httptask.New(),
	)

	reg.Register(
		scripttask.New(),
	)

	reg.Register(
		approvaltask.New(
			approvaltask.NewStore(),
		),
	)

	reg.Register(
		k8stask.New(),
	)
}