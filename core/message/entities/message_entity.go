package entities

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Text   string `json:"text"`
	UserID uint
	ChatID uint
}
