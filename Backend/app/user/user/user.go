package user

import (
	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	AvatarUrl string    `json:"avatar_url"`
}
