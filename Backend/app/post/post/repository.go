package post

import (
	"context"

	"gorm.io/gorm"
)

type RepositoryDB_Post interface {
	CreatePost(ctx context.Context, tx *gorm.DB, post *Post) error
	AddCategorieInPost(ctx context.Context, tx *gorm.DB, post *Post) error
}
