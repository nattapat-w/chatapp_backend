package service

import (
	"github.com/nattapat-w/chatapp/core/user/model/request"
	"github.com/nattapat-w/chatapp/core/user/model/response"
)

type UserService interface {
	Register(request *request.RegisterRequest) *response.Response
}
