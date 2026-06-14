package api

import (
	"github.com/gin-gonic/gin"

	"taskforge/internal/api/handlers"
)

func RegisterRoutes(
	r *gin.Engine,

	workflow *handlers.WorkflowHandler,
	execution *handlers.ExecutionHandler,
	approval *handlers.ApprovalHandler,
	cancel *handlers.CancelHandler,
) {

	v1 := r.Group("/api/v1")

	{
		v1.POST(
			"/workflows",
			workflow.Create,
		)

		v1.GET(
			"/workflows/:id",
			workflow.Get,
		)

		v1.POST(
			"/workflows/:id/executions",
			execution.Start,
		)

		v1.GET(
			"/executions/:id",
			execution.Get,
		)

		v1.POST(
			"/executions/:id/cancel",
			cancel.Cancel,
		)

		v1.POST(
			"/executions/:executionId/tasks/:taskId/approve",
			approval.Resolve,
		)
	}
}