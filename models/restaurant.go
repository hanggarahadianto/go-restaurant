package models

import (
	"time"

	"github.com/google/uuid"
)

type Restaurant struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `json:"name" form:"name"`
	Content   string    `json:"content" form:"content"`
	Image     string    `json:"image" form:"image"`
	Phone     string    `json:"phone" form:"phone"`
	Order     []Order   `gorm:"Foreignkey:Restaurant_ID;association_foreignkey:ID;" json:"order"`
	Tags      []Tags    `gorm:"Foreignkey:Restaurant_ID;association_foreignkey:ID;" json:"tags"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreateRestaurantInput struct {
	Name    string `json:"name" form:"name"`
	Content string `json:"content" form:"content"`
	Image   string `json:"image" form:"image"`
	Phone   string `json:"phone" form:"phone"`
	// User      string    `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
