package posting

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/laureano/devzone/app/post/post"
	"github.com/laureano/devzone/config"
	"github.com/laureano/devzone/database/connect"
	mockKeycloak "github.com/laureano/devzone/mocks/keycloak"
	mockPost "github.com/laureano/devzone/mocks/repositories/db/post"
	"github.com/stretchr/testify/require"
	"gorm.io/datatypes"
)

func TestCreatePosting(t *testing.T) {
	t.Parallel()
	cfg := config.Load()
	ctx := context.Background()
	db, err := connect.ConnectToDB(cfg)
	require.NoError(t, err)

	mockCtrl := gomock.NewController(t)

	mockPost := mockPost.NewMockRepositoryDB_Post(mockCtrl)
	mockKeycloak := mockKeycloak.NewMockRepositoryIdentities(mockCtrl)
	svc := NewService(db, mockKeycloak, mockPost)

	type ContentJson struct {
		Example string `json:"example"`
	}

	tests := []struct {
		nameCase string
		test     func(t *testing.T)
	}{
		{
			nameCase: "SuccessCreate",
			test: func(t *testing.T) {
				content := ContentJson{
					Example: "test2",
				}

				//convert
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

				err = svc.CreatePost(ctx, postDAO.Categories, postDAO.Id_user, postDAO.Title, postDAO.Content)
				require.NoError(t, err)
			},
		},
	}
	for _, tcase := range tests {
		t.Run("case", func(caseT *testing.T) {
			caseT.Run(tcase.nameCase, func(test *testing.T) {
				test.Parallel()
				tcase.test(test)
			})
		})
	}
}
