package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/labstack/echo/v4"
	"github.com/laureano/devzone/config"
)

type KeycloakVerifier struct {
	Verifier *oidc.IDTokenVerifier
}

type KeycloakRoles struct {
	RealmAccess struct {
		Roles []string `json:"roles"`
	} `json:"realm_access"`
}

func NewKeycloakVerifier(cfg *config.Config) (*KeycloakVerifier, error) {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, cfg.KeycloakRealmURL)
	if err != nil {
		return nil, err
	}
	return &KeycloakVerifier{
		Verifier: provider.Verifier(&oidc.Config{
			SkipClientIDCheck: true,
		}),
	}, nil
}

func (k *KeycloakVerifier) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, "Missing authorization header")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.JSON(http.StatusUnauthorized, "Malformed authorization header")
		}

		token := parts[1]
		ctx := context.Background()
		idToken, err := k.Verifier.Verify(ctx, token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}

		var claims map[string]interface{}
		if err := idToken.Claims(&claims); err != nil {
			return c.JSON(http.StatusUnauthorized, "Failed to parse claims")
		}

		var roles KeycloakRoles
		if err := idToken.Claims(&roles); err != nil {
			return c.JSON(http.StatusUnauthorized, "Failed to parse roles")
		}

		c.Set("claims", claims)
		c.Set("roles", roles.RealmAccess.Roles)

		return next(c)
	}
}
