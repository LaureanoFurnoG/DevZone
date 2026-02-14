package categorizing

import (
	"context"

	"github.com/laureano/devzone/app/categories_tag/category"
	"gorm.io/gorm"
)

type Service interface {
	ListCategories(ctx context.Context) ([]category.Category, error)
}

type service struct {
	repository category.RepositoryDB_Category
	db         *gorm.DB
}

func NewService(db *gorm.DB, repo category.RepositoryDB_Category) Service {
	return &service{
		repository: repo,
		db:         db,
	}
}

func (s *service) ListCategories(ctx context.Context) ([]category.Category, error) {
	return s.repository.ListCategories(ctx, s.db.WithContext(ctx))
}

