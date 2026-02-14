package categoryrepository

import (
	"context"

	"github.com/laureano/devzone/app/categories_tag/category"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) category.RepositoryDB_Category {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) ListCategories(ctx context.Context, tx *gorm.DB) ([]category.Category, error) {
	var categories = []category.Category{}
	if err := tx.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
