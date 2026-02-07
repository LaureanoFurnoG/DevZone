package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	ServerPort int

	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing required env var: %s", key)
	}
	return value
}

func Load() *Config {
	serverPort, err := strconv.Atoi("5000")
	if err != nil {
		log.Fatal(err)
	}

	dbPort, err := strconv.Atoi("5432")
	if err != nil {
		log.Fatal(err)
	}
	
	return &Config{
		ServerPort: serverPort,

		DBHost:     "localhost",
		DBPort:     dbPort,
		DBUser:     "postgres",
		DBPassword: "Laureano",
		DBName:     "devzone",
	}
}
