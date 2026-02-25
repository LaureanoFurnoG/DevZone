package posting

import (
	"context"

	"github.com/google/uuid"
	"github.com/laureano/devzone/app/post/post"
	"github.com/laureano/devzone/app/user/user"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Service interface {
	CreatePost(ctx context.Context, categories []uint, Id_user uuid.UUID, title string, content datatypes.JSON) error
	ListPosts(ctx context.Context) ([]post.Post, error)
	ListPostsByCategoryID(ctx context.Context, categoryID uint) ([]post.Post, error)
	PostInformationByID(ctx context.Context, postID uint) (*post.Post, error)
	DeletePost(ctx context.Context, postId uint, authorId uuid.UUID, userID uuid.UUID) error
}

type service struct {
	repository     post.RepositoryDB_Post
	repositoryUser user.RepositoryDB_User
	db             *gorm.DB
}

func NewService(db *gorm.DB, repo post.RepositoryDB_Post, userRepo user.RepositoryDB_User) Service {
	return &service{
		repository:     repo,
		repositoryUser: userRepo,
		db:             db,
	}
}

func (s *service) enrichPostsWithAuthors(ctx context.Context, posts []post.Post) error {
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
		user *user.User
		err  error
	}

	jobs := make(chan string, len(uniqueIDs))
	results := make(chan result, len(uniqueIDs))

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for w := 0; w < workers; w++ {
		go func() {
			for id := range jobs {
				idUser, err := uuid.Parse(id)
				if err != nil {
					return
				}
				u, err := s.repositoryUser.GetUserByID(ctx, idUser, s.db.WithContext(ctx))
				results <- result{id: id, user: u, err: err}
			}
		}()
	}

	for _, id := range uniqueIDs {
		jobs <- id
	}
	close(jobs)

	userMap := make(map[string]*user.User, len(uniqueIDs))
	for range uniqueIDs {
		r := <-results
		if r.err != nil {
			cancel()
			return r.err
		}
		if r.user == nil {
			continue 
		}
		userMap[r.id] = r.user
	}

	for i := range posts {
		if u, ok := userMap[posts[i].Id_user.String()]; ok {
			posts[i].ProfileImage = &u.AvatarUrl
			posts[i].Username = u.Nickname
		}
	}

	return nil
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
	if err := s.enrichPostsWithAuthors(ctx, posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *service) ListPostsByCategoryID(ctx context.Context, categoryID uint) ([]post.Post, error) {
	posts, err := s.repository.ListPostsByID(ctx, s.db.WithContext(ctx), categoryID)
	if err != nil {
		return nil, err
	}
	if err := s.enrichPostsWithAuthors(ctx, posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *service) PostInformationByID(ctx context.Context, postID uint) (*post.Post, error) {
	p, err := s.repository.PostInformation(ctx, s.db.WithContext(ctx), postID)
	if err != nil {
		return nil, err
	}
	if err := s.enrichPostsWithAuthors(ctx, []post.Post{*p}); err != nil {
		return nil, err
	}
	return p, nil
}
func (s *service) DeletePost(ctx context.Context, postId uint, authorId uuid.UUID, userID uuid.UUID) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := s.repository.DeletePost(ctx, postId, tx)
		if err != nil {
			return err
		}
		return nil
	})
}
