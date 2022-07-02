package models

import (
	"time"

	"github.com/google/uuid"
)

type Employees struct {
	Employee_id         uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Employee_first_name string    `json:"employee_first_name"`
	Employee_last_name  string    `json:"employee_last_name"`
	Employee_tel        string    `json:"employee_tel_name"`
	Department_id       uuid.UUID
	Transections        []Transections `gorm:"foreignKey:Employee_id"`
	Create_at           time.Time
	Update_at           time.Time
	Delete_at           *time.Time
}
