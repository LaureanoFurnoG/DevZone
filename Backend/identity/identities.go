package identity
import "context"

type RepositoryIdentities interface {
	GetUserByID(ctx context.Context, id string) (*User, error)
}
