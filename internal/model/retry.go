package model

import "time"

type RetryPolicy struct {
	MaxRetries  int           `json:"maxRetries"`
	InitialWait time.Duration `json:"initialWait"`
	Multiplier  float64       `json:"multiplier"`
}