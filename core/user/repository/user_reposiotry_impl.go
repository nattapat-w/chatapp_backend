package repository

import (
	"errors"
	"strings"

	"github.com/nattapat-w/chatapp/core/user/dto"
	"github.com/nattapat-w/chatapp/core/user/entities"
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

func (r *GormUserRepository) CreateUser(user dto.UserDTO) (*entities.User, error) {
	userEntity := convertToUserEntity(user)
	result := r.db.Create(&userEntity)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), duplicateEntryMsg) {
			return nil, ErrDuplicateUser
		}
		return nil, result.Error
	}
	numRows := result.RowsAffected
	if numRows != numberRowInserted {
		return nil, ErrinsertUser
	}
	return &userEntity, nil
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
func (r *GormUserRepository) FindAllUser() ([]userEntity.User, error) {
	var users []userEntity.User
	// err := r.db.Find(&users)
	if err := r.db.Find(&users).Error; err != nil {
		// log.Fatalln(err)
		return nil, err
	}
	return users, nil
}
func convertToUserEntity(userDTO dto.UserDTO) userEntity.User {
	// Perform conversion logic here
	userEntity := userEntity.User{
		Username:    userDTO.Username,
		Password:    userDTO.Password,
		DisplayName: userDTO.DisplayName,
	}
	return userEntity
}
