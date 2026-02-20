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

	seen := make(map[string]struct{})
	var uniqueIDs []string
	for _, p := range posts {
		id := p.Id_user.String()
		if _, ok := seen[id]; !ok {
			seen[id] = struct{}{}
			uniqueIDs = append(uniqueIDs, id)
		}
	}

	const workers = 20

	type result struct {
		id   string
		user *identity.User
		err  error
	}

	jobs := make(chan string, len(uniqueIDs))
	results := make(chan result, len(uniqueIDs))

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for w := 0; w < workers; w++ {
		go func() {
			for id := range jobs {
				user, err := s.IdentitiesRepository.GetUserByID(ctx, id)
				results <- result{id: id, user: user, err: err}
			}
		}()
	}

	for _, id := range uniqueIDs {
		jobs <- id
	}
	close(jobs)

	userMap := make(map[string]*identity.User, len(uniqueIDs))
	for range uniqueIDs {
		r := <-results
		if r.err != nil {
			cancel()
			return nil, r.err
		}
		userMap[r.id] = r.user
	}

	for i := range posts {
		if user, ok := userMap[posts[i].Id_user.String()]; ok {
			posts[i].ProfileImage = user.ProfileImage
			posts[i].Username = user.Username
		}
	}

	return posts, nil
}

func (s *service) ListPostsByCategoryID(ctx context.Context, categoryID uint) ([]post.Post, error) {
	posts, err := s.repository.ListPostsByID(ctx, s.db.WithContext(ctx), categoryID)
	if err != nil {
		return nil, err
	}
	seen := make(map[string]struct{})
	var uniqueIDs []string
	for _, p := range posts {
		id := p.Id_user.String()
		if _, ok := seen[id]; !ok {
			seen[id] = struct{}{}
			uniqueIDs = append(uniqueIDs, id)
		}
	}

	const workers = 20

	type result struct {
		id   string
		user *identity.User
		err  error
	}

	jobs := make(chan string, len(uniqueIDs))
	results := make(chan result, len(uniqueIDs))

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for w := 0; w < workers; w++ {
		go func() {
			for id := range jobs {
				user, err := s.IdentitiesRepository.GetUserByID(ctx, id)
				results <- result{id: id, user: user, err: err}
			}
		}()
	}

	for _, id := range uniqueIDs {
		jobs <- id
	}
	close(jobs)

	userMap := make(map[string]*identity.User, len(uniqueIDs))
	for range uniqueIDs {
		r := <-results
		if r.err != nil {
			cancel()
			return nil, r.err
		}
		userMap[r.id] = r.user
	}

	for i := range posts {
		if user, ok := userMap[posts[i].Id_user.String()]; ok {
			posts[i].ProfileImage = user.ProfileImage
			posts[i].Username = user.Username
		}
	}

	return posts, nil
}
