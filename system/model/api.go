package model

type SystemApi struct {
	ID        int    `gorm:"primarykey" json:"id"`
	Name      string `gorm:"size:100" json:"name"`
	SecretId  string `gorm:"size:20" json:"secret_id"`
	SecretKey string `gorm:"size:32" json:"secret_key"`
	Status    bool   `json:"status"`
}
