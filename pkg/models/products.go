package models

import (
	"time"

	"github.com/google/uuid"
)

type Products struct {
	Product_id            uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Product_code          string    `json:"product_code"`
	Product_name          string    `json:"product_name"`
	Product_remain        int       `json:"product_remain"`
	Product_price         float32   `json:"product_price"`
	Product_type_id       uuid.UUID
	Transections_Products []Transections_Products `gorm:"foreignKey:Product_id"`
	Create_at             time.Time
	Update_at             time.Time
	Delete_at             *time.Time
}
