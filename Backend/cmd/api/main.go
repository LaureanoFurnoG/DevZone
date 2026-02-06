package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/laureano/devzone/config"
	"github.com/laureano/devzone/database/connect"
	"github.com/laureano/devzone/database/migrateDB"
	"github.com/laureano/devzone/server"
)

func main() {
	cfg := config.Load()
	dbGorm, err := connect.ConnectToDB(cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}
	
	if err := migrateDB.Migrate(dbGorm); err != nil{
		log.Fatal(err)
	}

	e, err := server.NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))); err != nil {
			log.Printf("shutting down the server")
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Printf("server forced to shutdo")
	}

	sqlDB, _ := dbGorm.DB()
	defer sqlDB.Close()

	log.Printf("exit sv")

}
