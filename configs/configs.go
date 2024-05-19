package configs

import (
	"log"
	"os"

	"ecurrency/pkg/repository"

	"github.com/joho/godotenv"
)

type Server struct {
	Port string
}

type Config struct {
	Server   Server
	Database repository.DBConfig
}

func ReadConfigs() *Config {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Failed to read .env: %v", err)
	}

	server_port := os.Getenv("SERVER_PORT")

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	sslmode := os.Getenv("SSLMODE")

	config := Config{
		Server: Server{Port: server_port},
		Database: repository.DBConfig{
			Host:     host,
			Port:     port,
			Username: username,
			Password: password,
			DBName:   dbname,
			SSLMode:  sslmode,
		},
	}

	return &config
}
