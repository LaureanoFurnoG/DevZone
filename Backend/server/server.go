package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/laureano/devzone/config"
	postrepository "github.com/laureano/devzone/database/repositories/post_repository"
)

func NewServer(cfg *config.Config) (*echo.Echo, error) {

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	
	r := e.Group("/devzone-api")

	postRepositoryDB := postrepository.NewPostRepository()

	log.Printf("Server listening on port %v", cfg.ServerPort)
	return e, nil
}
