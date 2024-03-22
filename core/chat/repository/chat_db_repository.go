package repository

type ChatDBRepository interface {
	CreateChat(uint, uint, string) error
}
