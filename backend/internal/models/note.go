package models

import "gorm.io/gorm"
import "time"

type Note struct {
	ID            string `gorm:"primaryKey;type:text" json:"id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Content       string         `json:"content"`
	ShareableLink string         `json:"shareable_link"`
	Username      string         `json:"username"`
	UserID        uint           `json:"user_id"`
}
