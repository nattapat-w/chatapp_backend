package repository

type MessageDBRepository interface {
	CreateMessage(string, uint, uint) error
}
