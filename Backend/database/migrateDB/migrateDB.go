package migrateDB

import (
	"github.com/laureano/devzone/database/models"
	"gorm.io/gorm"
)
//migrate tables, receive the db
func Migrate(db *gorm.DB) error{
	return db.AutoMigrate(
		&models.Categories{},
		&models.Post{},
		&models.Relation_categories{},
	)
}
