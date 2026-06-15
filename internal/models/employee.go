package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CompanyID    uint           `json:"company_id" gorm:"index"`
	UserID       *uint          `json:"user_id" gorm:"index"`
	NIK          string         `json:"nik" gorm:"size:50;not null"`
	FirstName    string         `json:"first_name" gorm:"size:100;not null"`
	LastName     string         `json:"last_name" gorm:"size:100"`
	Email        string         `json:"email" gorm:"size:255"`
	Phone        string         `json:"phone" gorm:"size:50"`
	DateOfBirth  *time.Time     `json:"date_of_birth"`
	Gender       string         `json:"gender" gorm:"size:20"`
	Address      string         `json:"address" gorm:"type:text"`
	DepartmentID uint           `json:"department_id" gorm:"index"`
	LocationID   uint           `json:"location_id" gorm:"index"`
	Position     string         `json:"position" gorm:"size:255"`
	Employment   string         `json:"employment" gorm:"size:50"` // PKWT, PKWTT, Magang
	Status       string         `json:"status" gorm:"size:20;default:active"` // active, inactive, terminated
	JoinDate     *time.Time     `json:"join_date"`
	LeftDate     *time.Time     `json:"left_date"`
	Photo        string         `json:"photo" gorm:"size:500"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`

	Company    Company    `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Department Department `json:"department,omitempty" gorm:"foreignKey:DepartmentID"`
	Location   Location   `json:"location,omitempty" gorm:"foreignKey:LocationID"`
	User       *User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
