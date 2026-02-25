package userManage

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/laureano/devzone/config"
	"github.com/laureano/devzone/database/connect"
	mockUser "github.com/laureano/devzone/mocks/repositories/db/user"
	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {
	t.Parallel()
	cfg := config.Load()
	ctx := context.Background()
	db, err := connect.ConnectToDB(cfg)
	require.NoError(t, err)

	mockCtrl := gomock.NewController(t)
	mockUser := mockUser.NewMockRepositoryDB_User(mockCtrl)

	svc := NewService(db, mockUser)

	tests := []struct {
		nameCase string
		test     func(t *testing.T)
	}{
		{
			nameCase: "SuccessCreate",
			test: func(t *testing.T) {
				idUser := uuid.New()

				mockUser.EXPECT().
					RegisterUser(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil)

				err = svc.RegisterUser(ctx, idUser, "testuser", "test@example.com", "https://example.com/avatar.png")
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