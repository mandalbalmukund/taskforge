package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"taskforge/internal/api/service"
)

type ExecutionHandler struct {
	service *service.ExecutionService
}

func NewExecutionHandler(
	s *service.ExecutionService,
) *ExecutionHandler {

	return &ExecutionHandler{
		service: s,
	}
}

func (h *ExecutionHandler) Start(
	c *gin.Context,
) {

	workflowID :=
		c.Param("id")

	exec, err :=
		h.service.Start(
			context.Background(),
			workflowID,
		)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		exec,
	)
}

func (h *ExecutionHandler) Get(
	c *gin.Context,
) {

	id := c.Param("id")

	exec, err :=
		h.service.Get(id)

	if err != nil {

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		exec,
	)
}