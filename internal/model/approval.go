package model

import "time"

type ApprovalConfig struct {
	Timeout   time.Duration `json:"timeout"`
	OnTimeout string        `json:"onTimeout"`
}