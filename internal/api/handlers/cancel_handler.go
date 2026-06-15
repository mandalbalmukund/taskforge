package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"taskforge/internal/cancellation"
)

type CancelHandler struct {
	manager *cancellation.Manager
}

func NewCancelHandler(
	m *cancellation.Manager,
) *CancelHandler {

	return &CancelHandler{
		manager: m,
	}
}

func (h *CancelHandler) Cancel(
	c *gin.Context,
) {

	executionID :=
		c.Param("id")

	if !h.manager.Cancel(
		executionID,
	) {

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "execution not found",
			},
		)

		return
	}

	c.Status(http.StatusOK)
}