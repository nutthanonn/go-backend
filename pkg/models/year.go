package models

import (
	"time"

	"github.com/google/uuid"
)

type Year struct {
	Year_id     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Year_name   string    `json:"year_name"`
	Create_date time.Time `json:"create_date"`
	Update_date time.Time `json:"update_date"`
}
