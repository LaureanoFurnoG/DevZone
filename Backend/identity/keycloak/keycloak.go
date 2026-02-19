package keycloak

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/laureano/devzone/config"
	"github.com/laureano/devzone/identity"
)

type IdentitiesRepository struct {
	client   *gocloak.GoCloak
	realm    string
	ClientID string
	secret   string
}

func NewKeycloakRepository(cfg *config.Config) *IdentitiesRepository {
	client := gocloak.NewClient(cfg.KeycloakRealmURL)

	return &IdentitiesRepository{
		client:   client,
		realm:    cfg.KeycloakRealm,
		ClientID: cfg.ClientID,
		secret:   cfg.KcRealmSecret,
	}
}

func (r *IdentitiesRepository) loginClient(ctx context.Context) (*gocloak.JWT, error) {
	token, err := r.client.LoginClient(
		ctx,
		r.ClientID,
		r.secret,
		r.realm,
	)
	if err != nil {
		return nil, err
	}

	return token, nil
}
var _ identity.RepositoryIdentities = (*IdentitiesRepository)(nil)

func (r *IdentitiesRepository) GetUserByID(ctx context.Context, id string) (*identity.User, error) {
	token, err := r.loginClient(ctx)
	if err != nil {
		return nil, err
	}

	kcUser, err := r.client.GetUserByID(
		ctx, token.AccessToken,
		r.realm,
		id,
	)
	if err != nil {
		return nil, err
	}

	var profileImage *string
	if kcUser.Attributes != nil {
		profileImage = getAttr(*kcUser.Attributes, "profileImage")
	}

	return &identity.User{
		ID:           gocloak.PString(kcUser.ID),
		Email:        gocloak.PString(kcUser.Email),
		Name:         gocloak.PString(kcUser.FirstName) + " " + gocloak.PString(kcUser.LastName),
		Username:     gocloak.PString(kcUser.Username),
		ProfileImage: profileImage,
	}, nil
}

// function to obtain the attributes (i have the profileImage here)
func getAttr(attrs map[string][]string, key string) *string {
	if attrs == nil {
		return nil
	}
	values, ok := attrs[key]
	if !ok || len(values) == 0 {
		return nil
	}

	return &values[0]
}
