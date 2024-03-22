package service

import repository "github.com/nattapat-w/chatapp/core/message/repostiory"

type MessageServiceImpl struct {
	repo repository.MessageDBRepository
}

func NewMessageSerivceImpl(repo repository.MessageDBRepository) MessageService {
	return &MessageServiceImpl{repo: repo}
}

func (c MessageServiceImpl) CreateMessage(text string, userID uint, chatID uint) error {
	err := c.repo.CreateMessage(text, userID, chatID)
	if err != nil {
		return err
	}
	return nil
}
