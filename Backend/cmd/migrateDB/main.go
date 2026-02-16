package main

import (
	"log"

	"github.com/laureano/devzone/config"
	"github.com/laureano/devzone/database/connect"
	"github.com/laureano/devzone/database/migrateDB"
)

func main() {
	cfg := config.Load()
	dbGorm, err := connect.ConnectToDB(cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := migrateDB.Migrate(dbGorm); err != nil {
		log.Fatal(err)
	}

	sqlDB, _ := dbGorm.DB()
	defer sqlDB.Close()

}
