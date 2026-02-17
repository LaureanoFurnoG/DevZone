package postrepository

import (
	"context"
	"fmt"

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
			PostID:     post.ID,
			CategoryID: post.Categories[i],
		}

		if err := tx.WithContext(ctx).Create(&postCategoriesDAO).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *postsRepository) ListPosts(ctx context.Context, tx *gorm.DB) ([]post.Post, error) {
	var postsDAO []models.Post

	if err := tx.WithContext(ctx).Order("created_at desc").
		Preload("Categories").
		Find(&postsDAO).Error; err != nil {
		return nil, err
	}

	posts := make([]post.Post, 0, len(postsDAO))
	for _, p := range postsDAO {
		postCurrent := make([]post.CategoriesPost, 0, len(p.Categories))
		for _, c := range p.Categories {
			postCurrent = append(postCurrent, post.CategoriesPost{ID: c.ID, Name: c.Name})
		}

		posts = append(posts, post.Post{
			ID:             p.ID,
			Id_user:        p.Id_user,
			Title:          p.Title,
			Content:        p.Content,
			CreatedAt:      p.CreatedAt,
			CategoriesData: postCurrent,
		})
	}

	return posts, nil
}

func (r *postsRepository) ListPostsByID(ctx context.Context, tx *gorm.DB, categoryID uint) ([]post.Post, error) {
	var postsDAO []models.Post
	fmt.Println(categoryID)
	if err := tx.WithContext(ctx).
		Joins("JOIN relation_categories ON relation_categories.post_id = posts.id").
		Where("relation_categories.categories_id = ?", categoryID).
		Preload("Categories").Order("posts.created_at desc").
		Find(&postsDAO).Error; err != nil {
		return nil, err
	}
	posts := make([]post.Post, 0, len(postsDAO))
	for _, p := range postsDAO {
		postCurrent := make([]post.CategoriesPost, 0, len(p.Categories))
		for _, c := range p.Categories {
			postCurrent = append(postCurrent, post.CategoriesPost{ID: c.ID, Name: c.Name})
		}

		posts = append(posts, post.Post{
			ID:             p.ID,
			Id_user:        p.Id_user,
			Title:          p.Title,
			Content:        p.Content,
			CreatedAt:      p.CreatedAt,
			CategoriesData: postCurrent,
		})
	}

	return posts, nil
}
