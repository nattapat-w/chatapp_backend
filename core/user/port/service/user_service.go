package service

import (
	"github.com/nattapat-w/chatapp/core/user/model"
)

type UserService interface {
	Register(request *model.RegisterRequest) *model.Response
}
