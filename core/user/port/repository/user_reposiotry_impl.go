package repository

import (
	"errors"
	"strings"

	"github.com/nattapat-w/chatapp/core/user/dto"
	"github.com/nattapat-w/chatapp/core/user/entities"
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
func convertToUserEntity(userDTO dto.UserDTO) entities.User {
	// Perform conversion logic here
	userEntity := entities.User{
		UserName:    userDTO.UserName,
		Password:    userDTO.Password,
		DisplayName: userDTO.DisplayName,
	}
	return userEntity
}
