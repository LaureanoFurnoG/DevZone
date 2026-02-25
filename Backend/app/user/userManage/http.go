package userManage

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/laureano/devzone/config"
	"github.com/laureano/devzone/middlewares"
)

func NewHTTPHandler(g *echo.Group, svc Service, cfg *config.Config) error {
	v1 := g.Group("/v1/user")
	kcVerifier, err := middlewares.NewKeycloakVerifier(cfg, svc)
	if err != nil {
		return err
	}

	v1.POST("", createUserHandler(svc), kcVerifier.Middleware)
	return nil
}

type registerUser struct {
	Id        uuid.UUID `json:"id"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	AvatarUrl string    `json:"avatar_url"`
}

func createUserHandler(svc Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req registerUser

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid request body")
		}

		err := svc.RegisterUser(c.Request().Context(), req.Id, req.Nickname, req.Email, req.AvatarUrl)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, nil)
	}
}
