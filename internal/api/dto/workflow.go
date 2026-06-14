package dto

type CreateWorkflowRequest struct {
	Name        string      `json:"name" binding:"required"`
	Description string      `json:"description"`
	Version     int         `json:"version"`
	Tasks        interface{} `json:"tasks"`
}