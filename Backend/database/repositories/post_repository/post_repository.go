package postrepository

import (
	"github.com/laureano/devzone/app/post/post"
	"gorm.io/gorm"
)

type postsRepository struct{}

func NewPostRepository() post.RepositoryDB {
	return &postsRepository{}
}

func (r *postsRepository) CreatePost(db *gorm.DB, post *post.Post) error {	
	if err := db.Create(post).Error; err != nil{
		return err
	}
	return nil
}