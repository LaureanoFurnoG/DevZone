package post

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type CategoriesPost struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Post struct {
	ID             uint             `json:"id"`
	Id_user        uuid.UUID        `json:"id_user"`
	Title          string           `json:"title"`
	Content        datatypes.JSON   `json:"content"`
	Categories     []uint           `json:"categories"`
	CategoriesData []CategoriesPost `json:"categoriesdata"`
	CreatedAt      time.Time        `json:"created_at"`
}
