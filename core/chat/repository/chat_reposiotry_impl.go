package repository

import (
	"github.com/nattapat-w/chatapp/core/chat/entities"
	"gorm.io/gorm"
)

type GormChatRepository struct {
	db *gorm.DB
}

func NewGormChatRepository(db *gorm.DB) ChatDBRepository {
	return &GormChatRepository{db: db}
}

func (r *GormChatRepository) CreateChat(userID1 uint, userID2 uint, chatName string) error {
	chat := &entities.Chat{
		// UserID:   userID,
		Chatname: chatName,
	}
	if err := r.db.Create(chat).Error; err != nil {
		return err
	}

	chatOwner1 := &entities.ChatOwners{
		ChatID: chat.ID,
		UserID: userID1,
	}
	if err := r.db.Create(chatOwner1).Error; err != nil {
		return err
	}

	chatOwner2 := &entities.ChatOwners{
		ChatID: chat.ID,
		UserID: userID2,
	}
	if err := r.db.Create(chatOwner2).Error; err != nil {
		return err
	}
	return nil
}
