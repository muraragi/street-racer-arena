package main

import (
	"log"

	"muraragi/street-racing-arena-backend/internal/api"
	"muraragi/street-racing-arena-backend/internal/auth"
	"muraragi/street-racing-arena-backend/internal/database"
	"muraragi/street-racing-arena-backend/internal/repositories"
	"muraragi/street-racing-arena-backend/internal/services"

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

	auth.InitializeAuth()

	repositories := repositories.InitializeRepositories(db)
	services := services.InitializeServices(repositories)

	router := api.InitializeRouter(api.RouterDependencies{
		BaseCarService: services.BaseCarService,
		UserService:    services.UserService,
		CarService:     services.CarService,
	})

	router.Run(":8080")
}
