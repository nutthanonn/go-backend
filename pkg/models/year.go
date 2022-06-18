package models

import "time"

type Year struct {
	Year_id     string    `json:"year_id" gorm:"primary_key"`
	Year_name   string    `json:"year_name"`
	Create_date time.Time `json:"create_date"`
	Update_date time.Time `json:"update_date"`
}
