package initializers

import (
	"fmt"
	"github.com/laureano/devzone/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(cfg *config.Config) error {
	adminDSN := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBPort,
	)

	adminDB, err := gorm.Open(postgres.Open(adminDSN), &gorm.Config{})
	if err != nil {
		return err
	}

	var exists bool
	checkQuery := `
		SELECT EXISTS (
			SELECT 1
			FROM pg_database
			WHERE datname = ?
		);
	`

	if err := adminDB.Raw(checkQuery, cfg.DBName).Scan(&exists).Error; err != nil {
		return err
	}

	if !exists {
		stmt := fmt.Sprintf("CREATE DATABASE %s", cfg.DBName)
		if err := adminDB.Exec(stmt).Error; err != nil {
			return err
		}
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}
