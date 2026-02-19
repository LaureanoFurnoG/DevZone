package posting

import (
	"context"

	"github.com/google/uuid"
	"github.com/laureano/devzone/app/post/post"
	"github.com/laureano/devzone/identity"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Service interface {
	CreatePost(ctx context.Context, categories []uint, Id_user uuid.UUID, title string, content datatypes.JSON) error
	ListPosts(ctx context.Context) ([]post.Post, error)
	ListPostsByCategoryID(ctx context.Context, categoryID uint) ([]post.Post, error)
}

type service struct {
	repository           post.RepositoryDB_Post
	IdentitiesRepository identity.RepositoryIdentities
	db                   *gorm.DB
}

func NewService(db *gorm.DB, identitiesRepo identity.RepositoryIdentities, repo post.RepositoryDB_Post) Service {
	return &service{
		repository:           repo,
		IdentitiesRepository: identitiesRepo,
		db:                   db,
	}
}

func (s *service) CreatePost(ctx context.Context, categories []uint, Id_user uuid.UUID, title string, content datatypes.JSON) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		postDAO := &post.Post{
			Id_user:    Id_user,
			Title:      title,
			Content:    content,
			Categories: categories,
		}

		err := s.repository.CreatePost(ctx, tx, postDAO)
		if err != nil {
			return err
		}
		err = s.repository.AddCategorieInPost(ctx, tx, postDAO)
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *service) ListPosts(ctx context.Context) ([]post.Post, error) {
	posts, err := s.repository.ListPosts(ctx, s.db.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	//this is a shit, i need change this in the future, for the moment is ok.
	for i := 0; i < len(posts); i++ {
		user, err := s.IdentitiesRepository.GetUserByID(ctx, posts[i].Id_user.String())
		if err != nil {
			return nil, err
		}

		posts[i].ProfileImage = user.ProfileImage
		posts[i].Username = user.Username
	}

	return posts, nil
}

func (s *service) ListPostsByCategoryID(ctx context.Context, categoryID uint) ([]post.Post, error) {
	return s.repository.ListPostsByID(ctx, s.db.WithContext(ctx), categoryID)
}
