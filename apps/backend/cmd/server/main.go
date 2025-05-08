package main

import (
	"log"

	"muraragi/street-racer-arena-backend/internal/api"
	"muraragi/street-racer-arena-backend/internal/database"
	"muraragi/street-racer-arena-backend/internal/repositories"
	"muraragi/street-racer-arena-backend/internal/services"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Warning: Could not load .env file:", err)
	}

	database.ConnectDB()
	database.MigrateDB()

	db := database.GetDB()

	repositories := repositories.InitializeRepositories(db)
	services := services.InitializeServices(repositories)

	router := api.InitializeRouter(api.RouterDependencies{
		BaseCarService: services.BaseCarService,
	})

	router.Run(":8080")
}
