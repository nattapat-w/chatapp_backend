package entities

import (
	"github.com/nattapat-w/chatapp/core/message/entities"
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Chatname   string             `gorm:"not null"`
	Messages   []entities.Message `gorm:"foreignKey:ChatID"`
	ChatOwners []ChatOwners       `gorm:"foreignKey:ChatID"`
}
