package models

import "time"

type Department struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CompanyID uint      `json:"company_id" gorm:"index"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	Code      string    `json:"code" gorm:"size:20;not null"`
	ManagerID *uint     `json:"manager_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Company Company `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
}
