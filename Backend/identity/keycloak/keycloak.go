package keycloak

import (
	"context"
	"sync"
	"time"

	"github.com/Nerzal/gocloak/v13"
	"github.com/laureano/devzone/config"
	"github.com/laureano/devzone/identity"
)

type cachedToken struct {
	jwt       *gocloak.JWT
	expiresAt time.Time
}

type cachedUser struct {
	user      *identity.User
	expiresAt time.Time
}

type IdentitiesRepository struct {
	client   *gocloak.GoCloak
	realm    string
	ClientID string
	secret   string

	tokenMu    sync.Mutex
	tokenCache *cachedToken

	userMu    sync.RWMutex
	userCache map[string]cachedUser

	userCacheTTL time.Duration
}

func NewKeycloakRepository(cfg *config.Config) *IdentitiesRepository {
	client := gocloak.NewClient(cfg.KeycloakRealmURL)

	return &IdentitiesRepository{
		client:       client,
		realm:        cfg.KeycloakRealm,
		ClientID:     cfg.ClientID,
		secret:       cfg.KcRealmSecret,
		userCache:    make(map[string]cachedUser),
		userCacheTTL: 5 * time.Minute, 
	}
}

func (r *IdentitiesRepository) getToken(ctx context.Context) (*gocloak.JWT, error) {
	r.tokenMu.Lock()
	defer r.tokenMu.Unlock()

	if r.tokenCache != nil && time.Now().Before(r.tokenCache.expiresAt.Add(-30*time.Second)) {
		return r.tokenCache.jwt, nil
	}

	token, err := r.client.LoginClient(ctx, r.ClientID, r.secret, r.realm)
	if err != nil {
		return nil, err
	}

	r.tokenCache = &cachedToken{
		jwt:       token,
		expiresAt: time.Now().Add(time.Duration(token.ExpiresIn) * time.Second),
	}

	return token, nil
}

var _ identity.RepositoryIdentities = (*IdentitiesRepository)(nil)

func (r *IdentitiesRepository) GetUserByID(ctx context.Context, id string) (*identity.User, error) {
	r.userMu.RLock()
	if cached, ok := r.userCache[id]; ok && time.Now().Before(cached.expiresAt) {
		r.userMu.RUnlock()
		return cached.user, nil
	}
	r.userMu.RUnlock()

	token, err := r.getToken(ctx)
	if err != nil {
		return nil, err
	}

	kcUser, err := r.client.GetUserByID(ctx, token.AccessToken, r.realm, id)
	if err != nil {
		return nil, err
	}

	var profileImage *string
	if kcUser.Attributes != nil {
		profileImage = getAttr(*kcUser.Attributes, "profileImage")
	}

	user := &identity.User{
		ID:           gocloak.PString(kcUser.ID),
		Email:        gocloak.PString(kcUser.Email),
		Name:         gocloak.PString(kcUser.FirstName) + " " + gocloak.PString(kcUser.LastName),
		Username:     gocloak.PString(kcUser.Username),
		ProfileImage: profileImage,
	}

	r.userMu.Lock()
	r.userCache[id] = cachedUser{
		user:      user,
		expiresAt: time.Now().Add(r.userCacheTTL),
	}
	r.userMu.Unlock()

	return user, nil
}

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