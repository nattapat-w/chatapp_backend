package service

type MessageService interface {
	CreateMessage(string, uint, uint) error
}
