package models

import (
	"github.com/google/uuid"
)

type Menu struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title   string    `json:"title" form:"title"`
	Content string    `json:"content" form:"image"`
	Image   string    `json:"image" form:"image"`

	Restaurant_ID string `gorm:"column:restaurant_id"  json:"restaurant_id"`
}
