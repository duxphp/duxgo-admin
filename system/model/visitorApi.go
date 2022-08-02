package model

import (
	"gorm.io/datatypes"
)

type VisitorApi struct {
	ID      uint           `gorm:"primarykey" json:"id"`
	Date    datatypes.Date `gorm:"size:10" json:"date"`
	Url     string         `gorm:"size:255" json:"url"`
	Name    string         `gorm:"size:255" json:"name"`
	Method  string         `gorm:"size:10" json:"method"`
	Uv      int            `gorm:"size:10" json:"uv"`
	Pv      int            `gorm:"size:10" json:"pv"`
	MaxTime float64        `json:"max_time"`
	MinTime float64        `json:"min_time"`
}
