package main

import (
	"log"

	"muraragi/street-racer-arena-backend/internal/database"
	"muraragi/street-racer-arena-backend/internal/seeds"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Warning: Could not load .env file:", err)
	}

	log.Println("Starting database seeder...")

	database.ConnectDB()
	database.MigrateDB()

	db := database.GetDB()
	seeds.SeedCars(db)

	log.Println("Database seeding complete.")
}
