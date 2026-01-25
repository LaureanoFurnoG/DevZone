package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Categories struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"name"`
}

type Post struct {
	ID      int            `gorm:"primaryKey"`
	Id_user uuid.UUID      `gorm:"id_user"`
	Title   string         `gorm:"title"`
	Content datatypes.JSON `gorm:"content"`
}

type Relation_categories struct {
	ID      int  `gorm:"primaryKey"`
	Id_post int  `gorm:"id_post"`
	Post    Post `gorm:"foreignKey:Id_post;references:ID"`

	Id_categorie_tag int        `gorm:"id_categorie_tag"`
	Categorie        Categories `gorm:"foreignKey:Id_categorie_tag;references:ID"`
}
