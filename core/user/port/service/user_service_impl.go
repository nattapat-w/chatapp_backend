package service

import (
	"github.com/nattapat-w/chatapp/core/user/dto"
	"github.com/nattapat-w/chatapp/core/user/model/request"
	"github.com/nattapat-w/chatapp/core/user/model/response"
	"github.com/nattapat-w/chatapp/core/user/port/repository"
)

const (
	invalidUserNameErrMsg = "invalid username"
	invalidPasswordErrMsg = "invalid password"
)

type UserServiceImpl struct {
	repo repository.UserDBRepository
}

func NewUserServiceImpl(repo repository.UserDBRepository) UserService {
	return &UserServiceImpl{repo: repo}
}

func (u UserServiceImpl) Register(request *request.RegisterRequest) *response.Response {
	// validate request
	if len(request.Username) == 0 {
		return u.createFailedResponse(invalidUserNameErrMsg)
	}
	if len(request.Password) == 0 {
		return u.createFailedResponse(invalidPasswordErrMsg)
	}
	userDTO := dto.UserDTO{
		UserName:    request.Username,
		Password:    request.Password,
		DisplayName: request.DisplayName,
	}
	err := u.repo.CreateUser(userDTO)
	if err != nil {
		if err == repository.ErrDuplicateUser {
			return u.createFailedResponse(err.Error())
		}
		return u.createFailedResponse(err.Error())
	}
	registerData := response.RegisterDataResponse{
		UserName:    userDTO.UserName,
		DisplayName: userDTO.DisplayName,
	}
	return u.createSuccessResponse(registerData)
}

func (u UserServiceImpl) createFailedResponse(message string) *response.Response {
	var errorMessage *string
	if message == "" {
		errorMessage = nil
	} else {
		errorMessage = &message
	}
	return &response.Response{
		Status:       false,
		ErrorMessage: errorMessage,
	}
}
func (u UserServiceImpl) createSuccessResponse(data response.RegisterDataResponse) *response.Response {
	return &response.Response{
		Data:   data,
		Status: true,
	}
}
