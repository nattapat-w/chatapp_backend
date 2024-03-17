package repository

import (
	"github.com/nattapat-w/chatapp/core/user/entities"
)

type AuthDBRepository interface {
	SignUserAccessToken(req *entities.User) (string, error)
}
