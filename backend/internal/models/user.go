package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique; not null" json:"username"`
	Password string `json:"password"`
	Notes    []Note `json:"notes" gorm:"foreignKey:UserID"`
}
