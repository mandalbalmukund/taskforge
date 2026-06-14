package model

import "time"

type TaskDefinition struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Type        string                 `json:"type"`
	DependsOn   []string               `json:"dependsOn"`
	Config      map[string]interface{} `json:"config"`
	Condition   *Condition             `json:"condition,omitempty"`
	RetryPolicy RetryPolicy            `json:"retryPolicy"`
	Timeout     time.Duration          `json:"timeout"`
}

type Condition struct {
	Expression string `json:"expression"`
}