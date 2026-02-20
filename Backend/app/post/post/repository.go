package post

import (
	"context"

	"gorm.io/gorm"
)

type RepositoryDB_Post interface {
	CreatePost(ctx context.Context, tx *gorm.DB, post *Post) error
	AddCategorieInPost(ctx context.Context, tx *gorm.DB, post *Post) error
	ListPosts(ctx context.Context, tx *gorm.DB) ([]Post, error)
	ListPostsByID(ctx context.Context, tx *gorm.DB, categoryID uint) ([]Post, error)
	PostInformation(ctx context.Context, tx *gorm.DB, postId uint) (*Post, error)
}
