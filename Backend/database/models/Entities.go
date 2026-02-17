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
	PostID uint `gorm:"column:post_id"`
	Post   Post `gorm:"foreignKey:post_id;references:ID"`

	CategoryID uint       `gorm:"column:categories_id"`
	Category   Categories `gorm:"foreignKey:categories_id;references:ID"`
}
