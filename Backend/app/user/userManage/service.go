package userManage

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/laureano/devzone/app/user/user"
	"gorm.io/gorm"
)

type Service interface {
	RegisterUser(ctx context.Context, id_user uuid.UUID, nickname string, email string, avatar_url string) error
}

type service struct {
	repository user.RepositoryDB_User
	db         *gorm.DB
}

func NewService(db *gorm.DB, repo user.RepositoryDB_User) Service {
	return &service{
		repository: repo,
		db:         db,
	}
}

func (s *service) RegisterUser(ctx context.Context, id_user uuid.UUID, nickname string, email string, avatar_url string) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		userFound, err := s.repository.GetUserByID(ctx, id_user, tx)
		if err != nil {
			return err
		}

		if userFound != nil {
			return nil
		}

		newUser := user.User{
			Id:        id_user,
			Nickname:  nickname,
			Email:     email,
			AvatarUrl: avatar_url,
		}
		fmt.Print("aca")
		return s.repository.RegisterUser(ctx, newUser, tx)
	})
}
