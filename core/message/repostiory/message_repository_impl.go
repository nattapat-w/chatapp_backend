package repository

import (
	"fmt"

	"github.com/nattapat-w/chatapp/core/message/entities"
	"gorm.io/gorm"
)

type GormMessageRepository struct {
	db *gorm.DB
}

func NewGormMessageRepository(db *gorm.DB) MessageDBRepository {
	return &GormMessageRepository{db: db}
}

func (r *GormMessageRepository) CreateMessage(text string, userID uint, chatID uint) error {
	message := &entities.Message{
		Text:   text,
		UserID: userID,
		ChatID: chatID,
	}
	if err := r.db.Create(&entities.Message{Text: string(text), UserID: 1, ChatID: 1}).Error; err != nil {
		return err
	}
	fmt.Println(message)
	return nil
}
