package posting

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/laureano/devzone/app/post/post"
	"github.com/laureano/devzone/config"
	"github.com/laureano/devzone/middlewares"
	"gorm.io/datatypes"
)

func NewHTTPHandler(g *echo.Group, svc Service, userSvc middlewares.UserSyncer, cfg *config.Config) error {
	v1 := g.Group("/v1/posts")
	kcVerifier, err := middlewares.NewKeycloakVerifier(cfg, userSvc)
	if err != nil {
		return err
	}

	v1.POST("", createPostHandler(svc), kcVerifier.Middleware)
	v1.POST("/comment/:postId", createCommentHandler(svc), kcVerifier.Middleware)
	v1.GET("", listPostsHandler(svc))
	v1.GET("/searchpost/:title", listSearchPostsHandler(svc))
	v1.GET("/:categoryId", listPostsByCategoryIDHandler(svc))
	v1.GET("/publishedpost/:postId", postInformationHandler(svc))
	v1.DELETE("/:postId/:authorId", deletePostHandler(svc), kcVerifier.Middleware)
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
			return c.JSON(http.StatusBadRequest, "Issue with parse categoryID")
		}
		req.CategoryID = uint(CategoryIDuint)

		posts, err := svc.ListPostsByCategoryID(c.Request().Context(), req.CategoryID)
		if err != nil {
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
			return c.JSON(http.StatusBadRequest, "Issue with parse postID")
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

type deletePostRequest struct {
	UserID   uuid.UUID
	PostID   uint
	AuthorID uuid.UUID
}

func deletePostHandler(svc Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req deletePostRequest
		postID, err := strconv.ParseUint(c.Param("postId"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "issue with parse postID")
		}

		authorID, err := uuid.Parse(c.Param("authorId"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "issue with parse authorID")
		}

		userID, err := uuid.Parse(fmt.Sprint(c.Get("userID")))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Issue with the subject in the token")
		}

		req.PostID = uint(postID)
		req.AuthorID = authorID
		req.UserID = userID
		fmt.Println(req.UserID, req.AuthorID)
		if req.AuthorID != req.UserID {
			return c.JSON(http.StatusUnauthorized, "Wtf bro, i think that this post isn't your post")
		}

		err = svc.DeletePost(c.Request().Context(), req.PostID, req.AuthorID, req.UserID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, "Post deleted successfully")
	}
}

type listSearchPostsTitleResponse struct {
	Posts []post.Post `json:"posts"`
}

type listSearchPostsTitleRequest struct {
	Title string `json:"title"`
}

func listSearchPostsHandler(svc Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req listSearchPostsTitleRequest

		req.Title = c.Param("title")

		posts, err := svc.SearchPost(c.Request().Context(), req.Title)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, listSearchPostsTitleResponse{
			Posts: posts,
		})
	}
}

type createCommentRequest struct {
	Id_user uuid.UUID      `json:"id_user"`
	Id_post uint           `json:"postId"`
	Content datatypes.JSON `json:"content"`
}

func createCommentHandler(svc Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req createCommentRequest

		IdPostParam := c.Param("postId")

		idPostParse, err := strconv.ParseUint(IdPostParam, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid request param")
		}

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid request body")
		}

		req.Id_post = uint(idPostParse)
		
		err = svc.CreateComment(c.Request().Context(), req.Id_user, req.Id_post, req.Content)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, nil)
	}
}
