package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"taskforge/internal/api/service"
	"taskforge/internal/model"
)

type WorkflowHandler struct {
	service *service.WorkflowService
}

func NewWorkflowHandler(
	s *service.WorkflowService,
) *WorkflowHandler {

	return &WorkflowHandler{
		service: s,
	}
}

func (h *WorkflowHandler) Create(
	c *gin.Context,
) {

	var workflow model.Workflow

	if err :=
		c.ShouldBindJSON(
			&workflow,
		); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	result, err :=
		h.service.Create(
			workflow,
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
		result,
	)
}

func (h *WorkflowHandler) Get(
	c *gin.Context,
) {

	id := c.Param("id")

	workflow, err :=
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
		workflow,
	)
}