package models

import "time"

type Role struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CompanyID   uint      `json:"company_id" gorm:"index"`
	Name        string    `json:"name" gorm:"size:100;not null"`
	Description string    `json:"description" gorm:"size:500"`
	Permissions string    `json:"permissions" gorm:"type:text"` // JSON array of permission strings
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Company Company `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
}
