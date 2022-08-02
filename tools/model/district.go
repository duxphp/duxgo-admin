package model

import (
	"time"
)

type ToolDistrict struct {
	ID        int       `gorm:"primarykey" json:"id"`
	ParentId  int       `json:"parent_id"`
	Code      string    `gorm:"size:20" json:"code"`
	Name      string    `gorm:"size:100" json:"name"`
	Level     int       `gorm:"size:1" json:"level"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
