package model

import (
	"time"
)

type SystemUser struct {
	ID        uint         `gorm:"primarykey" json:"id"`
	Nickname  string       `gorm:"size:100" json:"nickname"`
	Username  string       `gorm:"type:varchar(255);uniqueIndex" json:"username"`
	Password  string       `gorm:"type:varchar(255)" json:"password"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Roles     []SystemRole `json:"roles" gorm:"many2many:system_user_role;"`
}
