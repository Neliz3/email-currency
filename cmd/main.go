package main

import (
	"ecurrency/configs"
	"ecurrency/pkg/handler"
	"ecurrency/pkg/repository"
	"ecurrency/pkg/service"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	configs := configs.ReadConfigs()

	db, err := repository.NewPostgres(repository.DBConfig{
		Host:     configs.Database.Host,
		Port:     configs.Database.Port,
		Username: configs.Database.Username,
		Password: configs.Database.Password,
		DBName:   configs.Database.DBName,
		SSLMode:  configs.Database.SSLMode,
	})

	if err != nil {
		log.Fatalf("Failed to connect to db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoute().Run(configs.Server.Port)

}
