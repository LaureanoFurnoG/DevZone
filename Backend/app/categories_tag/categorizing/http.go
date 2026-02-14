package categorizing

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/laureano/devzone/app/categories_tag/category"
)

func NewHTTPHandler(g *echo.Group, svc Service) {
	v1 := g.Group("/v1/categories")

	v1.GET("", listCategoriesHandler(svc))
}

type listCategoriesResponse struct {
	Categories []category.Category `json:"categories"`
}

func listCategoriesHandler(svc Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		listCategories, err := svc.ListCategories(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "internal_server_error",
			})
		}

		return c.JSON(http.StatusOK, listCategoriesResponse{
			Categories: listCategories,
		})
	}
}
