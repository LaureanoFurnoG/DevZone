package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/laureano/devzone/config"
)

type UserSyncer interface {
	RegisterUser(ctx context.Context, id_user uuid.UUID, nickname string, email string, avatar_url string) error
}

type KeycloakVerifier struct {
	Verifier    *oidc.IDTokenVerifier
	userService UserSyncer
}

func NewKeycloakVerifier(cfg *config.Config, userService UserSyncer) (*KeycloakVerifier, error) {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, fmt.Sprintf("%s/realms/%s", cfg.KeycloakRealmURL, cfg.KeycloakRealm))
	if err != nil {
		return nil, err
	}
	return &KeycloakVerifier{
		Verifier:    provider.Verifier(&oidc.Config{SkipClientIDCheck: true}),
		userService: userService,
	}, nil
}
type KeycloakRoles struct {
	RealmAccess struct {
		Roles []string `json:"roles"`
	} `json:"realm_access"`
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

		ctx := context.Background()
		idToken, err := k.Verifier.Verify(ctx, parts[1])
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}

		userID, err := uuid.Parse(idToken.Subject)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid subject token")
		}

		var claims map[string]interface{}
		if err := idToken.Claims(&claims); err != nil {
			return c.JSON(http.StatusUnauthorized, "Failed to parse claims")
		}

		var roles KeycloakRoles
		if err := idToken.Claims(&roles); err != nil {
			return c.JSON(http.StatusUnauthorized, "Failed to parse roles")
		}

		err = k.userService.RegisterUser(
			ctx,
			userID,
			getString(claims, "preferred_username"),
			getString(claims, "email"),
			getString(claims, "profileImage"),
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to sync user")
		}

		c.Set("claims", claims)
		c.Set("userID", userID)
		c.Set("roles", roles.RealmAccess.Roles)

		return next(c)
	}
}

func getString(claims map[string]interface{}, key string) string {
	if val, ok := claims[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}
