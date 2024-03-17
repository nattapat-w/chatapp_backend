package service

import (
	"github.com/nattapat-w/chatapp/core/user/dto"
	"github.com/nattapat-w/chatapp/core/user/model"
	"github.com/nattapat-w/chatapp/core/user/port/repository"
	"golang.org/x/crypto/bcrypt"
)

const (
	invalidUserNameErrMsg = "invalid username"
	invalidPasswordErrMsg = "invalid password"
	errHashPassword       = "hash password error"
)

type UserServiceImpl struct {
	repo repository.UserDBRepository
}

func NewUserServiceImpl(repo repository.UserDBRepository) UserService {
	return &UserServiceImpl{repo: repo}
}

func (u UserServiceImpl) Register(request *model.RegisterRequest) *model.Response {
	if len(request.Username) == 0 {
		return u.createFailedResponse(invalidUserNameErrMsg)
	}
	if len(request.Password) == 0 {
		return u.createFailedResponse(invalidPasswordErrMsg)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return u.createFailedResponse(errHashPassword)
	}
	request.Password = string(hashedPassword)
	userDTO := dto.UserDTO{
		Username:    request.Username,
		Password:    request.Password,
		DisplayName: request.DisplayName,
	}
	err = u.repo.CreateUser(userDTO)
	if err != nil {
		if err == repository.ErrDuplicateUser {
			return u.createFailedResponse(err.Error())
		}
		return u.createFailedResponse(err.Error())
	}
	registerData := model.RegisterDataResponse{
		UserName:    userDTO.Username,
		DisplayName: userDTO.DisplayName,
	}
	return u.createSuccessResponse(registerData)
}
func (u UserServiceImpl) createFailedResponse(message string) *model.Response {
	var errorMessage *string
	if message == "" {
		errorMessage = nil
	} else {
		errorMessage = &message
	}
	return &model.Response{
		Status:       false,
		ErrorMessage: errorMessage,
	}
}
func (u UserServiceImpl) createSuccessResponse(data model.RegisterDataResponse) *model.Response {
	return &model.Response{
		Data:   data,
		Status: true,
	}
}
