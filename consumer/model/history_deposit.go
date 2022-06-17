package model

import "time"

type HistoryDeposit struct {
	Amount        *float64       `json:"amount,omitempty"`
	CreatedAt     *time.Time     `json:"created_at,omitempty"`
	DifferentTime *time.Duration `json:"time,omitempty"`
}
