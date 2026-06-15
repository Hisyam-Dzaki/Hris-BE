package models

import "time"

type Company struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null;size:255"`
	Slug      string    `json:"slug" gorm:"uniqueIndex;size:100"`
	Address   string    `json:"address" gorm:"type:text"`
	Phone     string    `json:"phone" gorm:"size:50"`
	Email     string    `json:"email" gorm:"size:255"`
	Website   string    `json:"website" gorm:"size:255"`
	Logo      string    `json:"logo" gorm:"size:500"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
