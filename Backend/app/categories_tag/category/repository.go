package category

import (
	"context"

	"gorm.io/gorm"
)

type RepositoryDB_Category interface {
	ListCategories(ctx context.Context, tx *gorm.DB) ([]Category, error)
}
