package model

import "gorm.io/datatypes"

type SystemRole struct {
	ID          int            `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:100" json:"name"`
	Permissions datatypes.JSON `json:"permissions"`
}
