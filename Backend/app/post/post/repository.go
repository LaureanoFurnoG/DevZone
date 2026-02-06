package post

import (
	"gorm.io/gorm"
)

type RepositoryDB interface {
	CreatePost(db *gorm.DB, post *Post) error
}