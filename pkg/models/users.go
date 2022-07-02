package models

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	User_id            uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	User_first_name    string         `json:"user_first_name"`
	User_last_name     string         `json:"user_last_name"`
	User_year_of_birth string         `json:"user_year_of_birth"`
	User_tel           string         `json:"user_tel"`
	User_email         string         `json:"user_email"`
	User_address       string         `json:"user_address"`
	Transections       []Transections `gorm:"foreignKey:User_id"`
	Create_at          time.Time
	Update_at          time.Time
	Delete_at          *time.Time
}
