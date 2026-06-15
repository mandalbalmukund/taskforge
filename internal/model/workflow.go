package model

import "time"

type Workflow struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Version     int              `json:"version"`
	Description string           `json:"description"`
	Tasks       []TaskDefinition `json:"tasks"`
	CreatedAt   time.Time        `json:"createdAt"`
}