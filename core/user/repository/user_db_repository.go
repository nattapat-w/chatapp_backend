package repository

import (
	"github.com/nattapat-w/chatapp/core/user/dto"
	"github.com/nattapat-w/chatapp/core/user/entities"
)

type UserDBRepository interface {
	CreateUser(user dto.UserDTO) (*entities.User, error)
	FindOneUser(user dto.UserDTO) (*entities.User, error)
	FindAllUser() ([]entities.User, error)
}
