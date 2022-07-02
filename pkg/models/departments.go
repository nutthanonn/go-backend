package models

import (
	"time"

	"github.com/google/uuid"
)

type Departments struct {
	Department_id   uuid.UUID   `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Department_name string      `json:"department_name"`
	Employees       []Employees `gorm:"foreignKey:Department_id"`
	Create_at       time.Time
	Update_at       time.Time
	Delete_at       time.Time
}
