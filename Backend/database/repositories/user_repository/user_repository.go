package postrepository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/laureano/devzone/app/user/user"
	"github.com/laureano/devzone/database/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.RepositoryDB_User {
	return &userRepository{db: db}
}

func (r *userRepository) RegisterUser(ctx context.Context, user user.User, tx *gorm.DB) error {
	userDAO := models.User{
		Id:        user.Id,
		Nickname:  user.Nickname,
		AvatarUrl: user.AvatarUrl,
		Email:     user.Email,
	}

	if err := tx.WithContext(ctx).Create(&userDAO).Error; err != nil {
		return err
	}

	user.Id = userDAO.Id
	return nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id uuid.UUID, tx *gorm.DB) (*user.User, error) {
	userDAO := models.User{}
	err := tx.WithContext(ctx).Where("id = ?", id).First(&userDAO).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil 
		}
		return nil, err 
	}

	return &user.User{
		Id:        userDAO.Id,
		Nickname:  userDAO.Nickname,
		Email:     userDAO.Email,
		AvatarUrl: userDAO.AvatarUrl,
	}, nil
}