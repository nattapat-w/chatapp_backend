package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName    string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	DisplayName string
}
