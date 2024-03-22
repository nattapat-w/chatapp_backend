package service

type ChatService interface {
	CreateChat(uint, uint, string) error
}
