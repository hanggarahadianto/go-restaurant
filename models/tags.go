package models

import (
	"github.com/google/uuid"
)

type Tags struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title         string    `json:"title" form:"title"`
	Restaurant_ID string    `gorm:"column:restaurant_id"  json:"restaurant_id"`
}
