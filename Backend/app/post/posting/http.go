package posting

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/laureano/devzone/app/post/post"
	"gorm.io/datatypes"
)

func NewHTTPHandler(g *echo.Group, svc Service) {
	v1 := g.Group("/v1/posts")

	v1.POST("", createPostHandler(svc))
	v1.GET("", listPostsHandler(svc))
}

type createPostRequest struct {
	Categories []uint         `json:"categories"`
	Id_user    uuid.UUID      `json:"id_user"`
	Title      string         `json:"title"`
	Content    datatypes.JSON `json:"content"`
}

func createPostHandler(svc Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req createPostRequest

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid request body")
		}

		err := svc.CreatePost(c.Request().Context(), req.Categories, req.Id_user, req.Title, req.Content)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, nil)
	}
}

type listPostsResponse struct {
	Posts []post.Post `json:"posts"`
}

func listPostsHandler(svc Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := svc.ListPosts(c.Request().Context())
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, listPostsResponse{
			Posts: posts,
		})
	}
}
