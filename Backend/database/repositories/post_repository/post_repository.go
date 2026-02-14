package postrepository

import (
	"context"

	"github.com/laureano/devzone/app/post/post"
	"github.com/laureano/devzone/database/models"
	"gorm.io/gorm"
)

type postsRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) post.RepositoryDB_Post {
	return &postsRepository{db: db}
}

func (r *postsRepository) CreatePost(ctx context.Context, tx *gorm.DB, post *post.Post) error {
	postDAO := models.Post{
		Id_user: post.Id_user,
		Title:   post.Title,
		Content: post.Content,
	}

	if err := tx.WithContext(ctx).Create(&postDAO).Error; err != nil {
		return err
	}

	post.ID = postDAO.ID
	return nil
}

func (r *postsRepository) AddCategorieInPost(ctx context.Context, tx *gorm.DB, post *post.Post) error {
	for i := range post.Categories {
		postCategoriesDAO := models.Relation_categories{
			Id_post:          post.ID,
			Id_categorie_tag: post.Categories[i],
		}

		if err := tx.WithContext(ctx).Create(&postCategoriesDAO).Error; err != nil {
			return err
		}
	}
	return nil
}
