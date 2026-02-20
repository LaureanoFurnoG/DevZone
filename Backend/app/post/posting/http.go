package posting

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/laureano/devzone/app/post/post"
	"github.com/laureano/devzone/config"
	"github.com/laureano/devzone/middlewares"
	"gorm.io/datatypes"
)

func NewHTTPHandler(g *echo.Group, svc Service, cfg *config.Config) error {
	v1 := g.Group("/v1/posts")
	kcVerifier, err := middlewares.NewKeycloakVerifier(cfg)
	if err != nil {
		return err
	}

	v1.POST("", createPostHandler(svc), kcVerifier.Middleware)
	v1.GET("", listPostsHandler(svc))
	v1.GET("/:categoryId", listPostsByCategoryIDHandler(svc))
	v1.GET("/publishedpost/:postId", postInformationHandler(svc))

	return nil
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

		roles := c.Get("roles").([]string)
		for i := 0; i < len(req.Categories); i++ {
			if req.Categories[i] == 4 && !contains(roles, "Admin") {
				return c.JSON(http.StatusForbidden, "You need admin role to use this category")
			}
		}

		err := svc.CreatePost(c.Request().Context(), req.Categories, req.Id_user, req.Title, req.Content)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, nil)
	}
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
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

type listPostsByCategoryIDResponse struct {
	Posts []post.Post `json:"posts"`
}

type listPostsByCategoryIDRequest struct {
	CategoryID uint `json:"category_id"`
}

func listPostsByCategoryIDHandler(svc Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req listPostsByCategoryIDRequest
		CategoryIDuint, err := strconv.ParseUint(c.Param("categoryId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Issue with parse categoryID")
		}
		req.CategoryID = uint(CategoryIDuint)

		posts, err := svc.ListPostsByCategoryID(c.Request().Context(), req.CategoryID)
		if err != nil {
			if err.Error() == "Post not exist" {
				return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, listPostsByCategoryIDResponse{
			Posts: posts,
		})
	}
}

type postInformationRequest struct {
	PostID uint
}

type postInformationResponse struct {
	Post *post.Post `json:"post"`
}

func postInformationHandler(svc Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req postInformationRequest
		postID, err := strconv.ParseUint(c.Param("postId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Issue with parse postID")
		}
		req.PostID = uint(postID)

		post, err := svc.PostInformationByID(c.Request().Context(), req.PostID)

		if err != nil {
			if err.Error() == "Post not exist" {
				return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, postInformationResponse{
			Post: post,
		})
	}
}
