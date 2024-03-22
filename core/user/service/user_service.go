package service

import (
	"github.com/nattapat-w/chatapp/core/user/entities"
	"github.com/nattapat-w/chatapp/core/user/model"
)

type UserService interface {
	Register(request *model.RegisterRequest) *model.Response
	GetAllUser() ([]entities.User, error)
}
