package repository

import (
	"errors"
	"strings"

	"github.com/nattapat-w/chatapp/core/user/dto"
	userEntity "github.com/nattapat-w/chatapp/core/user/entities"
	"gorm.io/gorm"
)

var (
	ErrDuplicateUser = errors.New("duplicate user (username is used)")
)

const (
	duplicateEntryMsg = "duplicate key value violates unique constraint"
	numberRowInserted = 1
)

var (
	ErrinsertUser = errors.New("failed to insert user")
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserDBRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) CreateUser(user dto.UserDTO) error {
	userEntity := convertToUserEntity(user)
	result := r.db.Create(&userEntity)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), duplicateEntryMsg) {
			return ErrDuplicateUser
		}
		return result.Error
	}
	numRows := result.RowsAffected
	if numRows != numberRowInserted {
		return ErrinsertUser
	}
	return nil
}
func (r *GormUserRepository) FindOneUser(user dto.UserDTO) (*userEntity.User, error) {
	selectedUser := new(userEntity.User)
	if err := r.db.Where("username = ?", user.Username).First(selectedUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return selectedUser, nil
}
func convertToUserEntity(userDTO dto.UserDTO) userEntity.User {
	// Perform conversion logic here
	userEntity := userEntity.User{
		Username: userDTO.Username,
		Password: userDTO.Password,
	}
	return userEntity
}
