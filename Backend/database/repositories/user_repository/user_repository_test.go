package postrepository

import (
	"context"
	"testing"

	"github.com/go-jose/go-jose/v4/testutils/require"
	"github.com/google/uuid"
	"github.com/laureano/devzone/app/user/user"
	"github.com/laureano/devzone/config"
	"github.com/laureano/devzone/database/connect"
)

func TestRepositoryUser(t *testing.T) {
	cfg := config.Load()
	db, err := connect.ConnectToDB(cfg)
	require.NoError(t, err)

	userRepo := NewUserRepository(db)

	tests := []struct {
		nameCase string
		test     func(t *testing.T)
	}{
		{
			nameCase: "SuccessCreate",
			test: func(t *testing.T) {
				newUser := user.User{
					Id:        uuid.New(),
					Nickname:  "testuser",
					AvatarUrl: "https://example.com/avatar.png",
					Email:     "test@example.com",
				}

				err = userRepo.RegisterUser(context.Background(), newUser, db)
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