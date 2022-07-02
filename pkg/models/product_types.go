package models

import (
	"time"

	"github.com/google/uuid"
)

type Product_types struct {
	Product_type_id   uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Product_type_name string     `json:"product_type_name"`
	Products          []Products `gorm:"foreignKey:Product_type_id"`
	Create_at         time.Time
	Update_at         time.Time
	Delete_at         *time.Time
}
