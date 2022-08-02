package model

type ToolFileDir struct {
	ID      int    `gorm:"primarykey" json:"id"`
	Name    string `gorm:"size:100" json:"name"`
	HasType string `gorm:"size:20" json:"has_type"`
}
