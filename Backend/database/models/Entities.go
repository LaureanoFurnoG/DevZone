package models

import (
	"time"

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
	Id_user    uuid.UUID      `gorm:"type:uuid;index"`
	User       User           `gorm:"foreignKey:Id_user;references:ID"`
	Title      string         `gorm:"title"`
	Content    datatypes.JSON `gorm:"content"`
	Categories []Categories   `gorm:"many2many:relation_categories;"`
}

type Comment struct {
	gorm.Model
	Id_user    uuid.UUID      `gorm:"type:uuid;index"`
	User       User           `gorm:"foreignKey:Id_user;references:ID"`
	Id_Post    uint           `gorm:"type:uint;index"`
	Post       Post           `gorm:"foreignKey:Id_Post;references:ID"`
	Content    datatypes.JSON `gorm:"content"`
}

type Relation_categories struct {
	gorm.Model
	PostID uint `gorm:"column:post_id"`
	Post   Post `gorm:"foreignKey:post_id;references:ID"`

	CategoryID uint       `gorm:"column:categories_id"`
	Category   Categories `gorm:"foreignKey:categories_id;references:ID"`
}

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Nickname  string    `gorm:"nickname"`
	Email     string    `gorm:"email"`
	AvatarUrl string    `gorm:"avatar_url"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
