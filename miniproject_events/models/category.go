package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Event    string `json:"event" form:"event"`
	Category string `json:"category" form:"category"`
}
