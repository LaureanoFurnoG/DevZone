package post

import (
	"context"

	"gorm.io/gorm"
)

type RepositoryDB interface {
	CreatePost(ctx context.Context, tx *gorm.DB, post *Post) error
}
