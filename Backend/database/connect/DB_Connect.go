package connect

import (
	"fmt"
	"time"

	"github.com/laureano/devzone/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectToDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=UTC", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	//open connection with postgres preparestmt == cache queries, improve the performance
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	dbSQL, err := db.DB() //connection pool
	if err != nil {
		return nil, err
	}

	dbSQL.SetMaxOpenConns(25)
	dbSQL.SetMaxIdleConns(10)
	dbSQL.SetConnMaxLifetime(time.Hour)

	return db, nil
}
