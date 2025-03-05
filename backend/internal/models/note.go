package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Content       string `json:"content"`
	ShareableLink string `json:"shareable_link"`
	Username      string `json:"username"`
	UserID        uint   `json:"user_id"`
}
