package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID            string `gorm:"primaryKey;type:text;default:gen_random_uuid()" json:"id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Content       string         `json:"content"`
	ShareableLink string         `json:"shareable_link"`
	Username      string         `json:"username"`
	UserID        uint           `json:"user_id"`
}
