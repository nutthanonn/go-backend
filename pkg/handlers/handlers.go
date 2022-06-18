package handlers

import "gorm.io/gorm"

type handlers struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handlers {
	return handlers{db}
}
