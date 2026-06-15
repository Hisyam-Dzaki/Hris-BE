package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CompanyID uint      `json:"company_id" gorm:"index"`
	Email     string    `json:"email" gorm:"uniqueIndex;size:255;not null"`
	Password  string    `json:"-" gorm:"not null"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	RoleID    uint      `json:"role_id" gorm:"index"`
	RoleName  string    `json:"role_name" gorm:"size:100;default:employee"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Company Company `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Role    Role    `json:"role,omitempty" gorm:"foreignKey:RoleID"`
}
