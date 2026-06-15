package models

import "time"

type Location struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CompanyID uint      `json:"company_id" gorm:"index"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	Address   string    `json:"address" gorm:"type:text"`
	City      string    `json:"city" gorm:"size:100"`
	Timezone  string    `json:"timezone" gorm:"size:50;default:Asia/Jakarta"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Company Company `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
}
