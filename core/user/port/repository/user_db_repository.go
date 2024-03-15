package repository

import (
	"github.com/nattapat-w/chatapp/core/user/dto"
)

type UserDBRepository interface {
	CreateUser(user dto.UserDTO) error
}
