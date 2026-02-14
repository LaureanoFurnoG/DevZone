package categoryrepository

import (
	"context"
	"testing"

	"github.com/go-jose/go-jose/v4/testutils/require"
	"github.com/laureano/devzone/config"
	"github.com/laureano/devzone/database/connect"
)

func TestRepositoryCategiry(t *testing.T) {
	cfg := config.Load()
	db, err := connect.ConnectToDB(cfg)
	require.NoError(t, err)

	categoryRepo := NewCategoryRepository(db)

	tests := []struct {
		nameCase string
		test     func(t *testing.T)
	}{
		{
			nameCase: "ListCategories_Success",
			test: func(t *testing.T) {
				//add create category in the future
				/* categories := []models.Categories{
					{
						ID: 1,
						Name: "Keycloak",
					},
					{
						ID: 2,
						Name:"React",
					},
				} */

				categoryRepo.ListCategories(context.Background(), db)
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
