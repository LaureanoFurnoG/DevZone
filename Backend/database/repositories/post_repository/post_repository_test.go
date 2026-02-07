package postrepository

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/go-jose/go-jose/v4/testutils/require"
	"github.com/google/uuid"
	"github.com/laureano/devzone/app/post/post"
	"github.com/laureano/devzone/config"
	"github.com/laureano/devzone/database/connect"
	"gorm.io/datatypes"
)

func TestRepositoryPost(t *testing.T) {
	cfg := config.Load()
	db, err := connect.ConnectToDB(cfg)
	require.NoError(t, err)

	postRepo := NewPostRepository(db)
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
				userID := uuid.New()

				content := ContentJson{
					Example: "test2",
				}

				//convert
				bytes, err := json.Marshal(content)
				require.NoError(t, err)

				newPost := post.Post{
					Id_user: userID,
					Title:   "test2",
					Content: datatypes.JSON(bytes),
				}

				err = postRepo.CreatePost(context.Background(), db, &newPost)
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
