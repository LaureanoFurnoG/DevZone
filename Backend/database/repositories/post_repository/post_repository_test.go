package postrepository

import (
	"encoding/json"
	"testing"

	"github.com/go-jose/go-jose/v4/testutils/require"
	"github.com/google/uuid"
	"github.com/laureano/devzone/app/post/post"
	"github.com/laureano/devzone/config"
	"github.com/laureano/devzone/database/connect"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func TestRepositoryPost(t *testing.T) {
	t.Parallel()

	postRepo := NewPostRepository()
	type ContentJson struct {
		Example string `json:"example"`
	}

	tests := []struct {
		nameCase string
		test     func(t *testing.T, db *gorm.DB)
	}{
		{
			nameCase: "SuccessCreate",
			test: func(t *testing.T, db *gorm.DB) {
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

				err = postRepo.CreatePost(db, &newPost)

			},
		},
	}
	for _, tcase := range tests {
		t.Run("case", func(caseT *testing.T) {
			caseT.Run(tcase.nameCase, func(test *testing.T) {
				test.Parallel()
				cfg := config.Load()

				db, err := connect.ConnectToDB(cfg)
				if err != nil {
					test.Error(err)
				}

				tcase.test(test, db)
			})
		})
	}
}
