package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/laureano/devzone/config"
	"github.com/laureano/devzone/initializers"
)

func NewServer(cfg *config.Config) (*echo.Echo, error) {
	if err := initializers.ConnectToDB(cfg); err != nil {
		return nil, err
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	log.Printf("Server listening on port %s", cfg.ServerPort)
	return e, nil
}
