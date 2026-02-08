package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Name string `gorm:"name"`
}

type Post struct {
	gorm.Model
	Id_user uuid.UUID      `gorm:"id_user"`
	Title   string         `gorm:"title"`
	Content datatypes.JSON `gorm:"content"`
}

type Relation_categories struct {
	gorm.Model
	Id_post uint  `gorm:"id_post"`
	Post    Post `gorm:"foreignKey:Id_post;references:ID"`

	Id_categorie_tag uint        `gorm:"id_categorie_tag"`
	Categorie        Categories `gorm:"foreignKey:Id_categorie_tag;references:ID"`
}
