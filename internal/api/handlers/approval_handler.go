package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"taskforge/internal/api/dto"
	"taskforge/internal/api/service"
)

type ApprovalHandler struct {
	service *service.ApprovalService
}

func NewApprovalHandler(
	s *service.ApprovalService,
) *ApprovalHandler {

	return &ApprovalHandler{
		service: s,
	}
}

func (h *ApprovalHandler) Resolve(
	c *gin.Context,
) {

	var req dto.ApprovalRequest

	if err :=
		c.ShouldBindJSON(
			&req,
		); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	executionID :=
		c.Param("executionId")

	taskID :=
		c.Param("taskId")

	var err error

	if req.Approved {

		err =
			h.service.Approve(
				executionID,
				taskID,
				req.Approver,
			)

	} else {

		err =
			h.service.Reject(
				executionID,
				taskID,
				req.Approver,
			)
	}

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.Status(http.StatusOK)
}