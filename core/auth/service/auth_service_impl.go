package service

import (
	"github.com/nattapat-w/chatapp/core/auth/model"
	authRepo "github.com/nattapat-w/chatapp/core/auth/repository"
	"github.com/nattapat-w/chatapp/core/user/dto"
	userRepo "github.com/nattapat-w/chatapp/core/user/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	userRepo userRepo.UserDBRepository
	authRepo authRepo.AuthDBRepository
}

func NewAuthServiceImpl(authRepo authRepo.AuthDBRepository, userRepo userRepo.UserDBRepository) AuthService {
	return &AuthServiceImpl{authRepo: authRepo, userRepo: userRepo}
}

// func (a AuthServiceImpl) Login(request *model.LoginRequest) (*model.UserLoginResponse, error) {
func (a AuthServiceImpl) Login(request *model.LoginRequest) (string, error) {

	userDTO := dto.UserDTO{
		Username: request.Username,
		Password: request.Password,
	}
	selectedUser, err := a.userRepo.FindOneUser(userDTO)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(userDTO.Password))
	if err != nil {
		return "", err
	}
	token, err := a.authRepo.SignUserAccessToken(selectedUser)
	if err != nil {
		return "", err
	}
	// res := &model.UserLoginResponse{
	// 	AccessToken: token,
	// }
	return token, nil
}
