package main

import (
	"github.com/laureano/devzone/database/models"
	"github.com/laureano/devzone/initializers"
)

func main() {
	initializers.DB.AutoMigrate(&models.Categories{})
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.Relation_categories{})
}
