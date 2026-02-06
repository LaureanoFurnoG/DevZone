package post

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Post struct {
	ID      int            `json:"id"`
	Id_user uuid.UUID      `json:"id_user"`
	Title   string         `json:"title"`
	Content datatypes.JSON `json:"content"`
}
