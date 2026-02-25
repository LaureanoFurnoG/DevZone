package user

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RepositoryDB_User interface {
	RegisterUser(ctx context.Context, user User, tx *gorm.DB) error
	GetUserByID(ctx context.Context, id uuid.UUID, tx *gorm.DB) (*User, error)
}
