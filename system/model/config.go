package model

import "gorm.io/datatypes"

type SystemConfig struct {
	HasType string            `gorm:"size:255" json:"has_type"`
	HasId   uint              `gorm:"size:11" json:"has_id"`
	Data    datatypes.JSONMap `json:"status"`
}
