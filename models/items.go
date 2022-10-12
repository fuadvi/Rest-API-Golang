package models

import (
	"time"
)

type Items struct {
	ID          uint   `gorm:"primaryKey"`
	Item_code   uint   `json:"Item_code" gorm:"not null"`
	Description string `json:"description" gorm:"not null;type:varchar(191)"`
	Quantity    uint
	CretaedAt   time.Time
	UpdatedAt   time.Time
}
