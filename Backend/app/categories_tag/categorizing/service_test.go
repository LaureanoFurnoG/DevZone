package categorizing

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/laureano/devzone/app/categories_tag/category"
	"github.com/laureano/devzone/app/post/post"
	"github.com/laureano/devzone/app/post/posting"
	"github.com/laureano/devzone/config"
	"github.com/laureano/devzone/database/connect"
	mockKeycloak "github.com/laureano/devzone/mocks/keycloak"
	mockCategory "github.com/laureano/devzone/mocks/repositories/db/category"
	mockPost "github.com/laureano/devzone/mocks/repositories/db/post"
	"gorm.io/datatypes"
)

func TestListCategories(t *testing.T) {
	t.Parallel()
	cfg := config.Load()
	ctx := context.Background()
	db, err := connect.ConnectToDB(cfg)
	require.NoError(t, err)

	mockCtrl := gomock.NewController(t)

	mockCategory := mockCategory.NewMockRepositoryDB_Category(mockCtrl)
	svc := NewService(db, mockCategory)

	mockPost := mockPost.NewMockRepositoryDB_Post(mockCtrl)
	mockKeycloak := mockKeycloak.NewMockRepositoryIdentities(mockCtrl)
	svcPost := posting.NewService(db, mockKeycloak, mockPost)

	type ContentJson struct {
		Example string `json:"example"`
	}

	tests := []struct {
		nameCase string
		test     func(t *testing.T)
	}{
		{
			nameCase: "SuccessListCategories",
			test: func(t *testing.T) {
				content := ContentJson{
					Example: "test2",
				}

				bytes, err := json.Marshal(content)
				require.NoError(t, err)

				postDAO := &post.Post{
					Id_user:    uuid.New(),
					Title:      "title example",
					Content:    datatypes.JSON(bytes),
					Categories: []uint{1, 2, 3},
				}

				mockPost.EXPECT().
					CreatePost(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil)

				mockPost.EXPECT().
					AddCategorieInPost(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil)

				err = svcPost.CreatePost(ctx, postDAO.Categories, postDAO.Id_user, postDAO.Title, postDAO.Content)
				require.NoError(t, err)

				mockCategory.EXPECT().
					ListCategories(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]category.Category{}, nil)

				categories, err := svc.ListCategories(ctx)
				require.NoError(t, err)
				require.NotNil(t, categories)
			},
		},
	}
	for _, tcase := range tests {
		t.Run("case", func(t *testing.T) {
			t.Run(tcase.nameCase, func(test *testing.T) {
				test.Parallel()
				tcase.test(test)
			})
		})
	}
}
