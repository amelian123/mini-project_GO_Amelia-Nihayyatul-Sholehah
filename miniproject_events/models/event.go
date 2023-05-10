package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	EventID uint   `json:"event_id" json:"event_id"`
	UserID  int    `json:"user_id" form:"user_id"`
	Name    string `json:"name" json:"name"`
	Kuota   int    `json:"kuota" form:"kuota"`
	Harga   int    `json:"harga" form:"harga"`
}
