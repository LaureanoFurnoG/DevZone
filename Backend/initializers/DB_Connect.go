package initializers

import (
	"github.com/laureano/devzone/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDB(cfg *config.Config) error{
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) //connect with the db

	if err != nil {
		log.Fatal("Failed to connect to DB")
	}
	return nil
}
