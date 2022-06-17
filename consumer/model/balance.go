package model

import (
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	BalanceAPI
	IsThreshold *bool      `json:"is_threshold,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type BalanceAPI struct {
	WalletID *uuid.UUID `json:"wallet_id,omitempty" gorm:"primaryKey"`
	Amount   *float64   `json:"amount,omitempty"`
}
