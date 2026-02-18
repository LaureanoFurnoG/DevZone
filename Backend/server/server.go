package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/laureano/devzone/app/categories_tag/categorizing"
	"github.com/laureano/devzone/app/post/posting"
	"github.com/laureano/devzone/config"
	"github.com/laureano/devzone/database/connect"
	categoryrepository "github.com/laureano/devzone/database/repositories/category_repository"
	postrepository "github.com/laureano/devzone/database/repositories/post_repository"
)

func NewServer(cfg *config.Config) (*echo.Echo, error) {
	db, err := connect.ConnectToDB(cfg)
	if err != nil {
		return nil, err
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		MaxAge:       300,
	}))
	r := e.Group("/devzone-api")

	postRepositoryDB := postrepository.NewPostRepository(db)
	postService := posting.NewService(db, postRepositoryDB)

	categoryRepositoryDB := categoryrepository.NewCategoryRepository(db)
	categoryService := categorizing.NewService(db, categoryRepositoryDB)

	posting.NewHTTPHandler(r, postService, cfg)
	categorizing.NewHTTPHandler(r, categoryService)
	log.Printf("Server listening on port %v", cfg.ServerPort)
	return e, nil
}
