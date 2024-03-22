package entities

import "gorm.io/gorm"

type ChatOwners struct {
	gorm.Model
	ChatID uint
	UserID uint
}
