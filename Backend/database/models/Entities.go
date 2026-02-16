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
	Id_user    uuid.UUID      `gorm:"id_user"`
	Title      string         `gorm:"title"`
	Content    datatypes.JSON `gorm:"content"`
	Categories []Categories   `gorm:"many2many:relation_categories;"`
}

type Relation_categories struct {
	gorm.Model
	Post_id uint `gorm:"Post_id"`
	Post    Post `gorm:"foreignKey:Post_id;references:ID"`

	Categories_id uint       `gorm:"Categories_id"`
	Categorie        Categories `gorm:"foreignKey:Categories_id;references:ID"`
}
