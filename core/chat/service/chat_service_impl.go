package service

import "github.com/nattapat-w/chatapp/core/chat/repository"

// import (
// 	"github.com/nattapat-w/chatapp/core/user/entities"
// 	"github.com/nattapat-w/chatapp/core/user/model"
// )

type ChatServiceImpl struct {
	repo repository.ChatDBRepository
}

func NewChatSerivceImpl(repo repository.ChatDBRepository) ChatService {
	return &ChatServiceImpl{repo: repo}
}

func (c ChatServiceImpl) CreateChat(userID1 uint, userID2 uint, chatName string) error {
	err := c.repo.CreateChat(userID1, userID2, chatName)
	if err != nil {
		return err
	}
	return nil
}
