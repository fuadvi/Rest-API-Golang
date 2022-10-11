package models

import (
	"time"
)

type Orders struct {
	ID            uint   `gorm:"primaryKey"`
	Customer_name string `gorm:"not null; type:varchar(191)"`
	Ordered_at    time.Time
	UpdatedAt     time.Time
}
