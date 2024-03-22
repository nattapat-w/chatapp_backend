package service

import (
	"github.com/nattapat-w/chatapp/core/auth/model"
)

type AuthService interface {
	// Login(request *model.LoginRequest) (*model.UserLoginResponse, error)
	Login(request *model.LoginRequest) (string, error)
}
