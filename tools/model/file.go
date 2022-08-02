package model

import (
	"gorm.io/datatypes"
	"time"
)

type ToolFile struct {
	ID        int            `gorm:"primarykey" json:"id"`
	DirId     int            `json:"dir_id"`
	Dir       ToolFileDir    `json:"dir"`
	HasType   string         `gorm:"size:20" json:"has_type"`
	Driver    string         `gorm:"size:50" json:"driver"`
	Url       string         `gorm:"size:255" json:"url"`
	Path      string         `gorm:"size:255" json:"path"`
	Title     string         `gorm:"size:255" json:"title"`
	Ext       string         `gorm:"size:20" json:"ext"`
	Size      int            `json:"size"`
	Extend    datatypes.JSON `json:"extend"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
