package model

import (
	"time"
)

type VisitorOperate struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Type      string    `gorm:"size:10" json:"type"`
	UserId    uint      `gorm:"size:11" json:"user_id"`
	Url       string    `gorm:"size:255" json:"url"`
	Method    string    `gorm:"size:10" json:"method"`
	Params    string    `json:"params"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
