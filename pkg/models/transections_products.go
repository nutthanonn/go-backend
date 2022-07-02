package models

import (
	"time"

	"github.com/google/uuid"
)

type Transections_Products struct {
	Transection_id uuid.UUID
	Product_id     uuid.UUID
	Amount         int `json:"amount"`
	Create_at      time.Time
	Update_at      time.Time
	Delete_at      *time.Time
}
