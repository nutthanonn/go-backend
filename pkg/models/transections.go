package models

import (
	"time"

	"github.com/google/uuid"
)

type Transections struct {
	Transection_id        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	User_id               uuid.UUID
	Employee_id           uuid.UUID
	Transections_Products []Transections_Products `gorm:"foreignKey:Transection_id"`
	Create_at             time.Time
	Update_at             time.Time
	Delete_at             *time.Time
}
