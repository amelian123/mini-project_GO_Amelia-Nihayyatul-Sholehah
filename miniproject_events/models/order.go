package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CategoryID int      `json:"category_id" form:"category_id"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID     int      `json:"user_id" form:"user_id"`
	User       User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EventName  string   `json:"event_name" form:"event_name"`
}
